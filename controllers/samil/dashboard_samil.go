package samil_controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	db_helpers "github.com/ofjangra/gwonline/db_helpers/users"
	"github.com/ofjangra/gwonline/models"
)

func Home_Samil(c *fiber.Ctx) error {
	user := new(models.SamilUser)
	userID := c.Locals("user_id").(string)

	fmt.Println(userID)

	decodeErr := db_helpers.GetSamilUserById(userID).Decode(&user)

	if decodeErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	return c.Status(200).JSON(fiber.Map{"userlink": user.UserLink, "challan_dash": user.ChallanDash, "challan_dash_url": user.ChallanDashURL})
}
