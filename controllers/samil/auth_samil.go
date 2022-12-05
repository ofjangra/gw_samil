package samil_controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/ofjangra/gwonline/db_helpers"
	"github.com/ofjangra/gwonline/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Signup_samil(c *fiber.Ctx) error {

	newUser := new(models.SamilUser)

	bodyParseErr := c.BodyParser(&newUser)

	if bodyParseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	if newUser.Email == "" || newUser.Password == "" || newUser.UserLink == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please provide required fields"})
	}

	if len(newUser.Password) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Password must contain atleast 6 characters"})
	}

	passwordHash, hashErr := bcrypt.GenerateFromPassword([]byte(newUser.Password), 12)

	if hashErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	newUser.Password = string(passwordHash)

	newUser.ID = primitive.NewObjectID()

	newUser.CreatedOn = primitive.NewDateTimeFromTime(time.Now())

	newUser.UpdatedOn = primitive.NewDateTimeFromTime(time.Now())

	signupErr := db_helpers.InsertUser_Samil(newUser)

	if signupErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": signupErr.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User Created"})
}

func Signin_samil(c *fiber.Ctx) error {

	type signinReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	JWTKEY := os.Getenv("JWTKEY")

	userCreds := new(signinReq)

	user := new(models.SamilUser)

	bodyParseErr := c.BodyParser(&userCreds)

	if bodyParseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong 1"})
	}

	if userCreds.Email == "" || userCreds.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please fill the required fields"})
	}

	thisuser := db_helpers.GetSamilUserByEmail(userCreds.Email)

	if thisuser.Err() != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	decodeErr := thisuser.Decode(&user)

	if decodeErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong 2"})
	}

	passMatchErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userCreds.Password))

	if passMatchErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Credentials"})
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

	cookie.Name = "access_id_sml"
	cookie.Value = tokenString
	cookie.HTTPOnly = true
	cookie.Expires = time.Now().Add(24 * time.Hour)

	c.Cookie(cookie)

	return c.SendStatus(fiber.StatusOK)

}

func Logout_samil(c *fiber.Ctx) error {

	cookie := new(fiber.Cookie)

	cookie.Name = "access_id"
	cookie.Value = ""
	cookie.HTTPOnly = true
	cookie.Expires = time.Now().Add(3 * time.Second)

	c.Cookie(cookie)

	return c.SendStatus(200)
}
