package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AdminAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Cookies("aid_gad") == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "You are not authorized"})
		}

		token, err := jwt.Parse(c.Cookies("aid_gad"), func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWTKEY")), nil
		})

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not verify token"})
		}

		payload := token.Claims.(jwt.MapClaims)

		c.Locals("employee_id", payload["id"])
		return c.Next()
	}
}
