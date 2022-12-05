package admin_controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ofjangra/gwonline/db_helpers"
	"github.com/ofjangra/gwonline/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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

func GetEmployeeProfile(c *fiber.Ctx) error {
	employee := new(models.Employee)

	employeeId := c.Params("id")

	thisEmployee, resErr := db_helpers.GetEmployeeById(employeeId)

	if resErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	decodeErr := thisEmployee.Decode(&employee)

	if decodeErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	viewerId := c.Locals("employee_id").(string)

	viewer := new(models.Employee)

	thisViewer, viewerResErr := db_helpers.GetEmployeeById(viewerId)

	if viewerResErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	viewerDecodeErr := thisViewer.Decode(&viewer)

	if viewerDecodeErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"vas": viewer.EmployeeType == "Super Admin", "employee": employee})
}

func DeleteEmployee(c *fiber.Ctx) error {
	empId := c.Params("id")

	deleteErr := db_helpers.DeleteEmployee(empId)

	if deleteErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": deleteErr.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Profile Deleted"})
}

func GetAllEmployees(c *fiber.Ctx) error {
	employees, err := db_helpers.GetAllEmployees()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch employees"})
	}

	return c.Status(fiber.StatusOK).JSON(employees)

}

func UpdateEmployeeProfile(c *fiber.Ctx) error {
	updateBody := bson.M{}

	bodyParseErr := c.BodyParser(&updateBody)

	if bodyParseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	value, ok := updateBody["password"]

	if ok {

		passHash, passHashErr := bcrypt.GenerateFromPassword([]byte(value.(string)), 12)

		if passHashErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update password"})
		}
		updateBody["password"] = string(passHash)

		updaterId := c.Locals("employee_id").(string)

		updateBody["updated_by"] = updaterId

		updateBody["updated_on"] = primitive.NewDateTimeFromTime(time.Now())

		employeeId := c.Params("id")

		profileUpdateErr := db_helpers.UpdateEmployeeProfile(employeeId, updateBody)

		if profileUpdateErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": profileUpdateErr.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Password updated successfully"})

	}

	updaterId := c.Locals("employee_id").(string)

	updateBody["updated_by"] = updaterId

	updateBody["updated_on"] = primitive.NewDateTimeFromTime(time.Now())

	employeeId := c.Params("id")

	profileUpdateErr := db_helpers.UpdateEmployeeProfile(employeeId, updateBody)

	if profileUpdateErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": profileUpdateErr.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Profile updated successfully"})
}
