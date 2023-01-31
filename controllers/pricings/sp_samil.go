package pricing_controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	db_helpers "github.com/ofjangra/gwonline/db_helpers/pricing"
	"github.com/ofjangra/gwonline/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateEntry_SP_Partner(c *fiber.Ctx) error {

	type entryProps struct {
		VehicleCategory string `query:"vehicle_category" json:"vehicle_category"`
		RTO             string `query:"rto" json:"rto"`
	}

	newEntryProps := new(entryProps)

	if queryParseErr := c.QueryParser(newEntryProps); queryParseErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No category and RTO were selected"})

	}

	if newEntryProps.VehicleCategory == "" || newEntryProps.RTO == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No passing and RTO were selected"})
	}

	newEntryProps.VehicleCategory = strings.ToUpper(newEntryProps.VehicleCategory)
	newEntryProps.RTO = strings.ToUpper(newEntryProps.RTO)

	if newEntryProps.VehicleCategory == "4W" || newEntryProps.VehicleCategory == "CV" || newEntryProps.VehicleCategory == "CE" {

		newEntry_4W := new(models.SP_Partner_4W)

		if bodyParseErr := c.BodyParser(&newEntry_4W); bodyParseErr != nil {
			fmt.Println(bodyParseErr)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to create entry"})
		}

		if newEntry_4W.RTO == "" || newEntry_4W.VehicleCategory == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please fill the required fields"})
		}

		newEntry_4W.ID = primitive.NewObjectID()

		newEntry_4W.RTO = strings.ToUpper(newEntry_4W.RTO)

		newEntry_4W.VehicleCategory = strings.ToUpper(newEntry_4W.VehicleCategory)

		newEntry_4W.CreatedBy = c.Locals("employee_id").(string)

		newEntry_4W.UpdatedBy = newEntry_4W.CreatedBy

		newEntry_4W.CreatedOn = primitive.NewDateTimeFromTime(time.Now())

		newEntry_4W.UpdatedOn = newEntry_4W.CreatedOn

		insertionErr := db_helpers.CreateEntry_SP_Partner(newEntry_4W.VehicleCategory, newEntry_4W.RTO, newEntry_4W)

		if insertionErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": insertionErr.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Entery Created Successfully 4w"})

	} else if newEntryProps.VehicleCategory == "2W" {
		newEntry_2W := new(models.SP_Partner_2W)

		if bodyParseErr := c.BodyParser(&newEntry_2W); bodyParseErr != nil {
			fmt.Println(bodyParseErr)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to create entry"})
		}

		if newEntry_2W.RTO == "" || newEntry_2W.VehicleCategory == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please fill the required fields"})
		}

		newEntry_2W.ID = primitive.NewObjectID()

		newEntry_2W.RTO = strings.ToUpper(newEntry_2W.RTO)

		newEntry_2W.VehicleCategory = strings.ToUpper(newEntry_2W.VehicleCategory)

		newEntry_2W.CreatedBy = c.Locals("employee_id").(string)

		newEntry_2W.UpdatedBy = newEntry_2W.CreatedBy

		newEntry_2W.CreatedOn = primitive.NewDateTimeFromTime(time.Now())

		newEntry_2W.UpdatedOn = newEntry_2W.CreatedOn

		insertionErr := db_helpers.CreateEntry_SP_Partner(newEntry_2W.VehicleCategory, newEntry_2W.RTO, newEntry_2W)

		if insertionErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": insertionErr.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Entery Created Successfully 2w "})
	} else if newEntryProps.VehicleCategory == "3W" {
		newEntry_3W := new(models.SP_Partner_3W)

		if bodyParseErr := c.BodyParser(&newEntry_3W); bodyParseErr != nil {
			fmt.Println(bodyParseErr)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to create entry"})
		}

		if newEntry_3W.RTO == "" || newEntry_3W.VehicleCategory == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please fill the required fields"})
		}

		newEntry_3W.ID = primitive.NewObjectID()

		newEntry_3W.RTO = strings.ToUpper(newEntry_3W.RTO)

		newEntry_3W.VehicleCategory = strings.ToUpper(newEntry_3W.VehicleCategory)

		newEntry_3W.CreatedBy = c.Locals("employee_id").(string)

		newEntry_3W.UpdatedBy = newEntry_3W.CreatedBy

		newEntry_3W.CreatedOn = primitive.NewDateTimeFromTime(time.Now())

		newEntry_3W.UpdatedOn = newEntry_3W.CreatedOn

		insertionErr := db_helpers.CreateEntry_SP_Partner(newEntry_3W.VehicleCategory, newEntry_3W.RTO, newEntry_3W)

		if insertionErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": insertionErr.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Entery Created Successfully 3w"})
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Something went wrong"})

}

func GetEntry_SP_Partner(c *fiber.Ctx) error {

	entry4W := new(models.SP_Partner_4W)
	entry3W := new(models.SP_Partner_3W)
	entry2W := new(models.SP_Partner_2W)

	type reqProps struct {
		VehicleCategory string `query:"vehicle_category" json:"vehicle_category"`
		RTO             string `query:"rto" json:"rto"`
	}

	getReqProps := new(reqProps)

	if queryParseErr := c.QueryParser(getReqProps); queryParseErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No category and RTO were selected"})

	}

	fmt.Println(getReqProps)

	if getReqProps.VehicleCategory == "4W" || getReqProps.VehicleCategory == "CV" || getReqProps.VehicleCategory == "CE" {

		entryRes := db_helpers.GetAnEntry_SP_Partner(getReqProps.VehicleCategory, getReqProps.RTO)

		decodeErr := entryRes.Decode(&entry4W)

		if decodeErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Entry not found"})
		}

		return c.Status(fiber.StatusOK).JSON(entry4W)

	}
	if getReqProps.VehicleCategory == "3W" {

		entryRes := db_helpers.GetAnEntry_SP_Partner("3W", getReqProps.RTO)

		decodeErr := entryRes.Decode(&entry3W)

		if decodeErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Entry not found"})
		}

		return c.Status(fiber.StatusOK).JSON(entry3W)

	}
	if getReqProps.VehicleCategory == "2W" {

		entryRes := db_helpers.GetAnEntry_SP_Partner("2W", getReqProps.RTO)

		decodeErr := entryRes.Decode(&entry2W)

		if decodeErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Entry not found"})
		}

		return c.Status(fiber.StatusOK).JSON(entry2W)

	}

	return nil
}

func UpdateEntry_SP_Partner(c *fiber.Ctx) error {

	entryID := c.Params("id")

	updateBody := bson.M{}

	bodyParseErr := c.BodyParser(&updateBody)

	if bodyParseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	updaterId := c.Locals("employee_id")

	updateBody["updated_by"] = updaterId

	updateBody["updated_on"] = primitive.NewDateTimeFromTime(time.Now())

	entryUpdateErr := db_helpers.UpdateEntry_SP_Partner(entryID, updateBody)

	if entryUpdateErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": entryUpdateErr.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Entry Updated successfully"})
}
