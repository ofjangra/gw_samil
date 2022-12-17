package admin_controllers

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/ofjangra/gwonline/db_helpers"
	"github.com/ofjangra/gwonline/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Signup_employee(c *fiber.Ctx) error {

	newEmployee := new(models.Employee)

	bodyParseErr := c.BodyParser(&newEmployee)

	if bodyParseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse body", "err": bodyParseErr.Error()})
	}

	fmt.Println(newEmployee)

	if newEmployee.Email == "" || newEmployee.Name == "" || newEmployee.Password == "" || newEmployee.Contact == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please provide required fields"})
	}

	if newEmployee.EmployeeType == "" || newEmployee.DOJ == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please provide required fields"})
	}

	if len(newEmployee.Password) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Password length must be grater than 6"})
	}

	passwordHash, hashErr := bcrypt.GenerateFromPassword([]byte(newEmployee.Password), 12)

	if hashErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	newEmployee.Email = strings.ToLower(newEmployee.Email)

	newEmployee.Password = string(passwordHash)

	newEmployee.ID = primitive.NewObjectID()

	newEmployee.CreatedOn = primitive.NewDateTimeFromTime(time.Now())

	newEmployee.UpdatedOn = primitive.NewDateTimeFromTime(time.Now())

	newEmployee.CreatedBy = c.Locals("employee_id").(string)

	newEmployee.UpdatedBy = c.Locals("employee_id").(string)

	employeeInsertionErr := db_helpers.InsertEmployee(newEmployee)

	if employeeInsertionErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": employeeInsertionErr.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Employee Created Successfully"})
}

func Signin_employee(c *fiber.Ctx) error {

	JWTKEY := os.Getenv("JWTKEY")

	type signinReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	employee := new(models.Employee)

	employeeCreds := new(signinReq)

	bodyParseErr := c.BodyParser(&employeeCreds)

	if bodyParseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse body"})
	}

	if employeeCreds.Email == "" || employeeCreds.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please provide required fields"})
	}

	thisEmployee := db_helpers.GetEmployeeByEmail(employeeCreds.Email)

	if thisEmployee.Err() != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	decodeErr := thisEmployee.Decode(&employee)

	if decodeErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	passMatchErr := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(employeeCreds.Password))

	if passMatchErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	tokenClaims := jwt.MapClaims{
		"id":  employee.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	tokenString, tokenErr := token.SignedString([]byte(JWTKEY))

	if tokenErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to login"})
	}

	cookie := new(fiber.Cookie)

	cookie.Name = "aid_gad"
	cookie.Value = tokenString
	cookie.HTTPOnly = true
	// cookie.SameSite = "None"
	// cookie.Secure = true
	cookie.Expires = time.Now().Add(24 * time.Hour)

	c.Cookie(cookie)

	return c.SendStatus(fiber.StatusOK)

}

func CreateSuperAdmin(c *fiber.Ctx) error {

	employee := new(models.Employee)

	updater := new(models.Employee)

	updaterId := c.Locals("employee_id").(string)

	thisUpdater, findingErr := db_helpers.GetEmployeeById(updaterId)

	if findingErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request Failed"})
	}

	updaterDecodeErr := thisUpdater.Decode(&updater)

	if updaterDecodeErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request Failed"})
	}

	if updater.EmployeeType != "Super Admin" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request Denied"})
	}

	empId := c.Params("id")

	thisEmployee, thisEmpErr := db_helpers.GetEmployeeById(empId)

	if thisEmpErr != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Request Failed"})
	}

	employeeDecodeErr := thisEmployee.Decode(&employee)

	if employeeDecodeErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request Failed"})
	}

	if employee.EmployeeType == "Super Admin" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Already a super admin"})
	}

	allowedOps := map[string]bool{"view": true, "add": true, "edit": true, "delete": true}

	privillages := map[string]map[string]bool{
		"dashboard":            allowedOps,
		"customer_master":      allowedOps,
		"employee_master":      allowedOps,
		"user_master":          allowedOps,
		"vehicle_master":       allowedOps,
		"admin_master":         allowedOps,
		"bulk_data_master":     allowedOps,
		"other_options_master": allowedOps,
	}

	updateBody := bson.M{"privillages": privillages, "employee_type": "Super Admin"}

	updateBody["updated_by"] = updaterId

	updateBody["updated_on"] = primitive.NewDateTimeFromTime(time.Now())

	updateErr := db_helpers.UpdateEmployeeProfile(empId, updateBody)

	if updateErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fulfill the request"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "New super admin added"})

}
