package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ofjangra/gwonline/controllers"
)

func Router_Employees(app *fiber.App) {
	router := app

	router.Post("/api/gwemployees/signup", controllers.Signup_employee)
	router.Post("/api/gwemployees/signin", controllers.Signin_employee)
}
