package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type ReqHeader struct {
	Authorization string
}

func Authrequired() fiber.Handler {
	return func(c *fiber.Ctx) error {

		if c.Cookies("access_id") == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "You are not authorized"})
		}

		token, err := jwt.Parse(c.Cookies("access_id"), func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWTKEY")), nil
		})

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not verify token"})
		}

		payload := token.Claims.(jwt.MapClaims)

		c.Locals("user_email", payload["email"])

		return c.Next()
	}
}
