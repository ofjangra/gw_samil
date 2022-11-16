package controllers

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

	newEmployee := new(models.SamilEmployee)

	bodyParseErr := c.BodyParser(&newEmployee)

	if bodyParseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	if newEmployee.Email == "" || newEmployee.Name == "" || newEmployee.Password == "" || newEmployee.Contact == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please provide required fields"})
	}

	if newEmployee.EmployeeType == "" || newEmployee.DOJ == "" || newEmployee.UserLink == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please provide required fields"})
	}

	if len(newEmployee.Password) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Password must contain atleast 6 characters"})
	}

	passwordHash, hashErr := bcrypt.GenerateFromPassword([]byte(newEmployee.Password), 12)

	if hashErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	newEmployee.Password = string(passwordHash)

	newEmployee.ID = primitive.NewObjectID()

	signupErr := db_helpers.InsertEmployee_Samil(newEmployee)

	if signupErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": signupErr.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Employee Created"})
}

func Signin_samil(c *fiber.Ctx) error {

	type signinReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	JWTKEY := os.Getenv("JWTKEY")

	userCreds := new(signinReq)

	employee := new(models.SamilEmployee)

	bodyParseErr := c.BodyParser(&userCreds)

	if bodyParseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong 1"})
	}

	if userCreds.Email == "" || userCreds.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please fill the required fields"})
	}

	thisEmployee := db_helpers.GetSamilEmployeeByEmail(userCreds.Email)

	if thisEmployee.Err() != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	decodeErr := thisEmployee.Decode(&employee)

	if decodeErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong 2"})
	}

	passMatchErr := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(userCreds.Password))

	if passMatchErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Credentials"})
	}

	tokenClaims := jwt.MapClaims{
		"email": employee.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	tokenString, tokenErr := token.SignedString([]byte(JWTKEY))

	if tokenErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to login"})
	}

	cookie := new(fiber.Cookie)

	cookie.Name = "access_id"
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