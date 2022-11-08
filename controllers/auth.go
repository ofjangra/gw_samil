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

type signinReq struct {
	UserID   string `json:"userid"`
	Password string `json:"password"`
}

func Signup(c *fiber.Ctx) error {

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

	fmt.Println(userCreds.UserID)

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

	sessionTokenClaims := jwt.MapClaims{
		"id":  user.UserID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	sessionToken := jwt.NewWithClaims(jwt.SigningMethodHS256, sessionTokenClaims)

	sessionTokenString, tokenErr := sessionToken.SignedString([]byte(JWTKEY))

	if tokenErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to login"})
	}

	cookie := new(fiber.Cookie)

	cookie.Name = "access_id"
	cookie.Value = sessionTokenString
	cookie.HTTPOnly = true
	cookie.Expires = time.Now().Add(24 * time.Hour)

	c.Cookie(cookie)

	return c.SendStatus(fiber.StatusOK)

}
