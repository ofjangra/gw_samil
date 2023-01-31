package admin_controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	db_helpers "github.com/ofjangra/gwonline/db_helpers/vehicles"
	"github.com/ofjangra/gwonline/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateVehicleEntry(c *fiber.Ctx) error {
	newVehicle := new(models.Vehicle)

	bodyParseErr := c.BodyParser(&newVehicle)

	if bodyParseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	if newVehicle.User == "" || newVehicle.Company == "" || newVehicle.RegisterationNumber == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please provide required fields"})
	}

	newVehicle.ID = primitive.NewObjectID()

	newVehicle.CreatedOn = primitive.NewDateTimeFromTime(time.Now())

	newVehicle.UpdatedOn = primitive.NewDateTimeFromTime(time.Now())

	newVehicle.CreatedBy = c.Locals("employee_id").(string)

	insertionErr := db_helpers.InsertAVehicle(newVehicle)

	if insertionErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": insertionErr.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Vehicle Created"})

}

func GetVehicleById(c *fiber.Ctx) error {
	vehicle := new(models.Vehicle)

	vehicleId := c.Params("id")

	thisVehicle, err := db_helpers.GetAVehicleById(vehicleId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	if thisVehicle.Err() != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "vehicle not found"})
	}

	decodeErr := thisVehicle.Decode(&vehicle)

	if decodeErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	return c.Status(fiber.StatusOK).JSON(vehicle)

}

func GetAllVehicles(c *fiber.Ctx) error {
	vehicles, err := db_helpers.GetAllVehicles()

	if err != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Failed to fetch vehicles"})
	}

	return c.Status(fiber.StatusOK).JSON(vehicles)
}
