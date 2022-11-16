package controllers

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

func Signup_employee(c *fiber.Ctx) error {

	newEmployee := new(models.Employee)

	bodyParseErr := c.BodyParser(&newEmployee)

	if bodyParseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create employee"})
	}

	fmt.Println(newEmployee)

	if newEmployee.Email == "" || newEmployee.Name == "" || newEmployee.Password == "" || newEmployee.Contact == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please provide required fields"})
	}

	if newEmployee.EmployeeType == "" || newEmployee.DOJ == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please provide required fields"})
	}

	if len(newEmployee.Password) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Password length must be grater than 6"})
	}

	passwordHash, hashErr := bcrypt.GenerateFromPassword([]byte(newEmployee.Password), 12)

	if hashErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	newEmployee.Password = string(passwordHash)

	newEmployee.ID = primitive.NewObjectID()

	employeeInsertionErr := db_helpers.InsertEmployee(newEmployee)

	if employeeInsertionErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": employeeInsertionErr.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Employee Created Successfully"})
}

func Signin_employee(c *fiber.Ctx) error {

	JWTKEY := os.Getenv("JWTKEY")

	type signinReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	employee := new(models.Employee)

	employeeCreds := new(signinReq)

	bodyParseErr := c.BodyParser(&employeeCreds)

	if bodyParseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse body"})
	}

	if employeeCreds.Email == "" || employeeCreds.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please provide required fields"})
	}

	thisEmployee := db_helpers.GetEmployeeByEmail(employeeCreds.Email)

	if thisEmployee.Err() != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	decodeErr := thisEmployee.Decode(&employee)

	if decodeErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	passMatchErr := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(employeeCreds.Password))

	if passMatchErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	tokenClaims := jwt.MapClaims{
		"email": employee.Email,
		"exp":   time.Now().Add(time.Minute * 45).Unix(),
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
	cookie.Expires = time.Now().Add(45 * time.Minute)

	c.Cookie(cookie)

	return c.SendStatus(fiber.StatusOK)

}
