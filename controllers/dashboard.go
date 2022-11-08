package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ofjangra/gwonline/db_helpers"
	"github.com/ofjangra/gwonline/models"
)

func Home(c *fiber.Ctx) error {
	user := new(models.User)
	userID := c.Locals("user_id").(string)

	fmt.Println(userID)

	decodeErr := db_helpers.GetUserByUserID(userID).Decode(&user)

	if decodeErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	return c.Status(200).JSON(fiber.Map{"userlink": user.UserLink})
}
