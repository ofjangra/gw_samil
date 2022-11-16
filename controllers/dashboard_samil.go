package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ofjangra/gwonline/db_helpers"
	"github.com/ofjangra/gwonline/models"
)

func Home_Samil(c *fiber.Ctx) error {
	user := new(models.SamilEmployee)
	userEmail := c.Locals("user_email").(string)

	fmt.Println(userEmail)

	decodeErr := db_helpers.GetSamilEmployeeByEmail(userEmail).Decode(&user)

	if decodeErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	return c.Status(200).JSON(fiber.Map{"userlink": user.UserLink})
}
