package samil_controllers

import (
	"github.com/gofiber/fiber/v2"
	db_helpers "github.com/ofjangra/gwonline/db_helpers/users"
	"github.com/ofjangra/gwonline/models"
)

func GetSamilUserProfile(c *fiber.Ctx) error {

	user := new(models.SamilUser)

	userid := c.Params("id")

	thisUser := db_helpers.GetSamilUserById(userid)

	if thisUser.Err() != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	decodeErr := thisUser.Decode(&user)

	if decodeErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	return c.Status(fiber.StatusOK).JSON(user)

}

func GetAllSamilUsers(c *fiber.Ctx) error {
	users, err := db_helpers.GetAllSamilUsers()

	if err != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
