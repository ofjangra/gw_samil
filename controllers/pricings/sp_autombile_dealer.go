package pricing_controllers

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ofjangra/gwonline/db_helpers"
	"github.com/ofjangra/gwonline/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateEntry_SP_AutomobileDealer(c *fiber.Ctx) error {
	newEntry := new(models.SP_AutomobileDealer)

	if bodyParseErr := c.BodyParser(&newEntry); bodyParseErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to create entry"})
	}

	if newEntry.Passing < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Passing"})
	}

	if newEntry.RTO == "" || newEntry.Location == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please fill the required fields"})
	}

	newEntry.ID = primitive.NewObjectID()

	newEntry.RTO = strings.ToUpper(newEntry.RTO)

	newEntry.CreatedBy = c.Locals("employee_id").(string)

	newEntry.UpdatedBy = newEntry.CreatedBy

	newEntry.CreatedOn = primitive.NewDateTimeFromTime(time.Now())

	newEntry.UpdatedOn = newEntry.CreatedOn

	insertionErr := db_helpers.CreateEntry_SP_AutomobileDealer(newEntry)

	if insertionErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": insertionErr.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Entery Created Successfully"})
}

func GetEntry_SP_AutomobileDealer(c *fiber.Ctx) error {

	entry := new(models.SP_AutomobileDealer)

	type reqProps struct {
		Passing float64 `query:"passing" json:"passing"`
		RTO     string  `query:"rto" json:"rto"`
	}

	getReqProps := new(reqProps)

	if queryParseErr := c.QueryParser(getReqProps); queryParseErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No passign and RTO were selected"})

	}

	entryRes := db_helpers.GetAnEntry_SP_AutomobileDealer(getReqProps.Passing, getReqProps.RTO)

	decodeErr := entryRes.Decode(&entry)

	if decodeErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch the entry"})
	}

	return c.Status(fiber.StatusOK).JSON(entry)
}

func UpdateEntry_SP_AutomobileDealer(c *fiber.Ctx) error {

	entryID := c.Params("id")

	updateBody := bson.M{}

	bodyParseErr := c.BodyParser(&updateBody)

	if bodyParseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	updaterId := c.Locals("employee_id")

	updateBody["updated_by"] = updaterId

	updateBody["updated_on"] = primitive.NewDateTimeFromTime(time.Now())

	entryUpdateErr := db_helpers.UpdateEntry_SP_AutomobileDealer(entryID, updateBody)

	if entryUpdateErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": entryUpdateErr.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Entry Updated successfully"})
}
