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

type signinReq struct {
	UserID   string `json:"userid"`
	Password string `json:"password"`
}

func Signup(c *fiber.Ctx) error {

	JWTKEY := os.Getenv("JWTKEY")

	newUser := new(models.User)

	bodyParseErr := c.BodyParser(&newUser)

	if bodyParseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	if newUser.UserID == "" || newUser.UserLink == "" || newUser.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please fill all the required fields"})
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

	signupErr := db_helpers.Insertuser(newUser)

	if signupErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": signupErr.Error()})
	}

	tokenClaims := jwt.MapClaims{
		"id":  newUser.UserID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Account Created"})
}

func Signin(c *fiber.Ctx) error {

	JWTKEY := os.Getenv("JWTKEY")

	userCreds := new(signinReq)

	user := new(models.User)

	bodyParseErr := c.BodyParser(&userCreds)

	if bodyParseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong 1"})
	}

	if userCreds.UserID == "" || userCreds.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please fill the required fields"})
	}

	thisUser := db_helpers.GetUserByUserID(userCreds.UserID)

	if thisUser.Err() != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	decodeErr := thisUser.Decode(&user)

	if decodeErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong 2"})
	}

	passMatchErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userCreds.Password))

	if passMatchErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Credentials"})
	}

	tokenClaims := jwt.MapClaims{
		"id":  user.UserID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
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

func Logout(c *fiber.Ctx) error {

	cookie := new(fiber.Cookie)

	cookie.Name = "access_id"
	cookie.Value = ""
	cookie.HTTPOnly = true
	cookie.Expires = time.Now().Add(3 * time.Second)

	c.Cookie(cookie)

	return c.SendStatus(200)
}
