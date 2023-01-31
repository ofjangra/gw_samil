package admin_controllers

import (
	"github.com/gofiber/fiber/v2"
	db_helpers "github.com/ofjangra/gwonline/db_helpers/users"
	"github.com/ofjangra/gwonline/models"
)

func GetCurrProfile(c *fiber.Ctx) error {
	user := new(models.FleetUser)

	userId := c.Locals("user_id").(string)

	thisuser, resErr := db_helpers.GetFleetUserById(userId)

	if resErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	decodeErr := thisuser.Decode(&user)

	if decodeErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": user})

}
