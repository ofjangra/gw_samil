package pricing_controllers

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ofjangra/gwonline/db"
	db_helpers "github.com/ofjangra/gwonline/db_helpers/pricing"
	"github.com/ofjangra/gwonline/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePricingEntry(c *fiber.Ctx) error {

	type entryProps struct {
		ServiceFor string `query:"service_for" json:"service_for"`
		Subject    string `query:"subject" json:"subject"`
	}

	newEntryProps := new(entryProps)

	if queryParseErr := c.QueryParser(newEntryProps); queryParseErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please provide valid data"})

	}

	if newEntryProps.Subject == "" || newEntryProps.ServiceFor == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please provide valid data"})
	}

	newEntryProps.Subject = strings.ToLower(newEntryProps.Subject)

	newEntryProps.ServiceFor = strings.ToLower(newEntryProps.ServiceFor)

	if newEntryProps.Subject == "passing" {

		newLoadingEntry := new(models.SP_Loading)

		newConsEqpEntry := new(models.SP_CE)

		newFarmEqpEntry := new(models.SP_FE)

		switch newEntryProps.ServiceFor {

		case "loading":
			if bodyParseErr := c.BodyParser(&newLoadingEntry); bodyParseErr != nil {
				fmt.Println(bodyParseErr)
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to create entry"})
			}
			if newLoadingEntry.Passing < 1 {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Passing"})
			}

			if newLoadingEntry.RTO == "" || newLoadingEntry.Location == "" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please fill the required fields"})
			}

			newLoadingEntry.ID = primitive.NewObjectID()

			newLoadingEntry.RTO = strings.ToUpper(newLoadingEntry.RTO)

			newLoadingEntry.CreatedBy = c.Locals("employee_id").(string)

			newLoadingEntry.UpdatedBy = newLoadingEntry.CreatedBy

			newLoadingEntry.CreatedOn = primitive.NewDateTimeFromTime(time.Now())

			newLoadingEntry.UpdatedOn = newLoadingEntry.CreatedOn

			insertionErr := db_helpers.CreateEntry_SP_Loading(newLoadingEntry.Passing, newLoadingEntry.RTO, newLoadingEntry, db.SP_LoadingCollection)

			if insertionErr != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": insertionErr.Error()})
			}

			return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Entery Created Successfully"})

		case "construction_eqp":
			if bodyParseErr := c.BodyParser(&newConsEqpEntry); bodyParseErr != nil {
				fmt.Println(bodyParseErr)
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to create entry"})
			}
			if newConsEqpEntry.Passing < 1 {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Passing"})
			}

			if newConsEqpEntry.RTO == "" || newConsEqpEntry.Location == "" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please fill the required fields"})
			}

			newConsEqpEntry.ID = primitive.NewObjectID()

			newConsEqpEntry.RTO = strings.ToUpper(newConsEqpEntry.RTO)

			newConsEqpEntry.CreatedBy = c.Locals("employee_id").(string)

			newConsEqpEntry.UpdatedBy = newConsEqpEntry.CreatedBy

			newConsEqpEntry.CreatedOn = primitive.NewDateTimeFromTime(time.Now())

			newConsEqpEntry.UpdatedOn = newConsEqpEntry.CreatedOn

			insertionErr := db_helpers.CreateEntry_SP_Loading(newConsEqpEntry.Passing, newConsEqpEntry.RTO, newConsEqpEntry, db.SP_ConstructionEqp)

			if insertionErr != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": insertionErr.Error()})
			}

			return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Entery Created Successfully"})

		case "farm_eqp":
			if bodyParseErr := c.BodyParser(&newFarmEqpEntry); bodyParseErr != nil {
				fmt.Println(bodyParseErr)
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to create entry"})
			}
			if newFarmEqpEntry.Passing < 1 {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Passing"})
			}

			if newFarmEqpEntry.RTO == "" || newFarmEqpEntry.Location == "" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please fill the required fields"})
			}

			newFarmEqpEntry.ID = primitive.NewObjectID()

			newFarmEqpEntry.RTO = strings.ToUpper(newFarmEqpEntry.RTO)

			newFarmEqpEntry.CreatedBy = c.Locals("employee_id").(string)

			newFarmEqpEntry.UpdatedBy = newFarmEqpEntry.CreatedBy

			newFarmEqpEntry.CreatedOn = primitive.NewDateTimeFromTime(time.Now())

			newFarmEqpEntry.UpdatedOn = newFarmEqpEntry.CreatedOn

			insertionErr := db_helpers.CreateEntry_SP_Loading(newFarmEqpEntry.Passing, newFarmEqpEntry.RTO, newFarmEqpEntry, db.SP_FarmEqp)

			if insertionErr != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": insertionErr.Error()})
			}

			return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Entery Created Successfully"})
		}
	}

	if newEntryProps.Subject == "seating" {

		newMotorCabEntry := new(models.SP_Auto_Cab)

		newBusEntry := new(models.SP_Bus)

		switch newEntryProps.ServiceFor {
		case "bus":
			if bodyParseErr := c.BodyParser(&newBusEntry); bodyParseErr != nil {
				log.Fatal(bodyParseErr)
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Data Bus"})
			}

			if newBusEntry.Seating < 1 {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid number of seats"})
			} else if newBusEntry.RTO == "" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please select an RTO"})
			}

			newBusEntry.ID = primitive.NewObjectID()

			newBusEntry.RTO = strings.ToUpper(newBusEntry.RTO)

			newBusEntry.CreatedBy = c.Locals("employee_id").(string)

			newBusEntry.UpdatedBy = newBusEntry.CreatedBy

			newBusEntry.CreatedOn = primitive.NewDateTimeFromTime(time.Now())

			newBusEntry.UpdatedOn = newBusEntry.CreatedOn

			insertionErr := db_helpers.CreateEntry_SP_Seating(newBusEntry.Seating, newBusEntry.RTO, newBusEntry, db.SP_BusCollection)

			if insertionErr != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": insertionErr.Error()})
			}

			return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Entery Created Successfully"})

		case "motor_cab":
			if bodyParseErr := c.BodyParser(&newMotorCabEntry); bodyParseErr != nil {
				log.Fatal(bodyParseErr)
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Data Bus"})
			}

			if newMotorCabEntry.Seating < 1 {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid number of seats"})
			} else if newMotorCabEntry.RTO == "" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please select an RTO"})
			}

			newMotorCabEntry.ID = primitive.NewObjectID()

			newMotorCabEntry.RTO = strings.ToUpper(newMotorCabEntry.RTO)

			newMotorCabEntry.CreatedBy = c.Locals("employee_id").(string)

			newMotorCabEntry.UpdatedBy = newMotorCabEntry.CreatedBy

			newMotorCabEntry.CreatedOn = primitive.NewDateTimeFromTime(time.Now())

			newMotorCabEntry.UpdatedOn = newMotorCabEntry.CreatedOn

			insertionErr := db_helpers.CreateEntry_SP_Seating(newMotorCabEntry.Seating, newMotorCabEntry.RTO, newMotorCabEntry, db.SP_MotorCab)

			if insertionErr != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": insertionErr.Error()})
			}

			return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Entery Created Successfully"})

		}
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to create entry, Invalid data"})
}

func GetPricingEntry_passing(c *fiber.Ctx) error {

	type reqProps struct {
		ServiceFor string  `query:"service_for" json:"service_for"`
		Passing    float32 `query:"passing" json:"passing"`
		RTO        string  `query:"rto" json:"rto"`
	}

	newReqProps := new(reqProps)

	if queryParseErr := c.QueryParser(newReqProps); queryParseErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query"})
	}

	if newReqProps.Passing < 1 || newReqProps.RTO == "" || newReqProps.ServiceFor == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid passing or rto"})
	}

	newReqProps.RTO = strings.ToUpper(newReqProps.RTO)

	newReqProps.ServiceFor = strings.ToLower(newReqProps.ServiceFor)

	switch newReqProps.ServiceFor {
	case "loading":
		entryLoading := new(models.SP_Loading)
		entryRes := db_helpers.GetAnEntry_SP_Loading(newReqProps.Passing, newReqProps.RTO, db.SP_LoadingCollection)

		decodeErr := entryRes.Decode(&entryLoading)

		if decodeErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Entry not found"})
		}

		return c.Status(fiber.StatusOK).JSON(entryLoading)

	case "construction_eqp":
		entryConsEqp := new(models.SP_CE)
		entryRes := db_helpers.GetAnEntry_SP_Loading(newReqProps.Passing, newReqProps.RTO, db.SP_ConstructionEqp)

		decodeErr := entryRes.Decode(&entryConsEqp)

		if decodeErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Entry not found"})
		}

		return c.Status(fiber.StatusOK).JSON(entryConsEqp)

	case "farm_eqp":
		entryFarmEqp := new(models.SP_FE)
		entryRes := db_helpers.GetAnEntry_SP_Loading(newReqProps.Passing, newReqProps.RTO, db.SP_FarmEqp)

		decodeErr := entryRes.Decode(&entryFarmEqp)

		if decodeErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Entry not found"})
		}

		return c.Status(fiber.StatusOK).JSON(entryFarmEqp)
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "entry Not Found"})
}
func GetPricingEntry_seating(c *fiber.Ctx) error {

	type reqProps struct {
		ServiceFor string `query:"service_for" json:"service_for"`
		Seating    int32  `query:"seating" json:"seating"`
		RTO        string `query:"rto" json:"rto"`
	}

	newReqProps := new(reqProps)

	if queryParseErr := c.QueryParser(newReqProps); queryParseErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query"})
	}

	if newReqProps.Seating < 1 || newReqProps.RTO == "" || newReqProps.ServiceFor == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid seating or rto"})
	}

	newReqProps.RTO = strings.ToUpper(newReqProps.RTO)

	newReqProps.ServiceFor = strings.ToLower(newReqProps.ServiceFor)

	switch newReqProps.ServiceFor {
	case "bus":
		entryBus := new(models.SP_Bus)
		entryRes := db_helpers.GetAnEntry_SP_Seating(newReqProps.Seating, newReqProps.RTO, db.SP_BusCollection)

		decodeErr := entryRes.Decode(&entryBus)

		if decodeErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Entry not found"})
		}

		return c.Status(fiber.StatusOK).JSON(entryBus)

	case "motor_cab":
		entryMotorCab := new(models.SP_Auto_Cab)
		entryRes := db_helpers.GetAnEntry_SP_Seating(newReqProps.Seating, newReqProps.RTO, db.SP_MotorCab)

		decodeErr := entryRes.Decode(&entryMotorCab)

		if decodeErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Entry not found"})
		}

		return c.Status(fiber.StatusOK).JSON(entryMotorCab)
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "entry Not Found"})
}

func UpdateEntry_SP_Loading(c *fiber.Ctx) error {

	entryID := c.Params("id")

	updateBody := bson.M{}

	bodyParseErr := c.BodyParser(&updateBody)

	if bodyParseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	updaterId := c.Locals("employee_id")

	updateBody["updated_by"] = updaterId

	updateBody["updated_on"] = primitive.NewDateTimeFromTime(time.Now())

	entryUpdateErr := db_helpers.UpdateEntry_SP_Loading(entryID, updateBody)

	if entryUpdateErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": entryUpdateErr.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Entry Updated successfully"})
}
