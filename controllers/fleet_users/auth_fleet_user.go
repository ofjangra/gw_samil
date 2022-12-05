package admin_controllers

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/ofjangra/gwonline/db_helpers"
	"github.com/ofjangra/gwonline/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Signup_fleet_user(c *fiber.Ctx) error {

	newUser := new(models.FleetUser)

	bodyParseErr := c.BodyParser(&newUser)

	if bodyParseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse body", "err": bodyParseErr.Error()})
	}

	fmt.Println(newUser)

	if newUser.Password == "" || newUser.Phone == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please provide required fields"})
	}

	if len(newUser.Password) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Password length must be grater than 6"})
	}

	passwordHash, hashErr := bcrypt.GenerateFromPassword([]byte(newUser.Password), 12)

	if hashErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	newUser.Password = string(passwordHash)

	newUser.ID = primitive.NewObjectID()

	newUser.CreatedOn = primitive.NewDateTimeFromTime(time.Now())

	newUser.UpdatedOn = primitive.NewDateTimeFromTime(time.Now())

	newUser.CreatedBy = c.Locals("employee_id").(string)

	newUser.UpdatedBy = c.Locals("employee_id").(string)

	userInsertionErr := db_helpers.CreateFleetUser(newUser)

	if userInsertionErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": userInsertionErr.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Account Created Successfully"})
}

func Signin_fleet_user(c *fiber.Ctx) error {

	JWTKEY := os.Getenv("JWTKEY")

	type signinReq struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	user := new(models.FleetUser)

	userCreds := new(signinReq)

	bodyParseErr := c.BodyParser(&userCreds)

	if bodyParseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse body"})
	}

	if userCreds.Phone == "" || userCreds.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please provide required fields"})
	}

	thisUser := db_helpers.GetEmployeeByEmail(userCreds.Phone)

	if thisUser.Err() != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	decodeErr := thisUser.Decode(&user)

	if decodeErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	passMatchErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userCreds.Password))

	if passMatchErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	tokenClaims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	tokenString, tokenErr := token.SignedString([]byte(JWTKEY))

	if tokenErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to login"})
	}

	cookie := new(fiber.Cookie)

	cookie.Name = "aid_fleet_user"
	cookie.Value = tokenString
	cookie.HTTPOnly = true
	cookie.Expires = time.Now().Add(24 * time.Hour)

	c.Cookie(cookie)

	return c.SendStatus(fiber.StatusOK)

}

func Logout_fleet_user(c *fiber.Ctx) error {

	cookie := new(fiber.Cookie)

	cookie.Name = "aid_fleet_user"
	cookie.Value = ""
	cookie.HTTPOnly = true
	cookie.Expires = time.Now().Add(3 * time.Second)

	c.Cookie(cookie)

	return c.SendStatus(200)
}
