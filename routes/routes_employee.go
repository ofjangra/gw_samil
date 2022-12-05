package routes

import (
	"github.com/gofiber/fiber/v2"
	admin_controllers "github.com/ofjangra/gwonline/controllers/admin"
	"github.com/ofjangra/gwonline/middleware"
)

func Router_Employees(app *fiber.App) {
	router := app

	router.Post("/api/main/emp/signup", middleware.AdminAuth(), admin_controllers.Signup_employee)

	router.Post("/api/main/emp/signin", admin_controllers.Signin_employee)

	router.Get("/api/gwadmin/profile", middleware.AdminAuth(), admin_controllers.GetAdminProfile)

	router.Get("/api/main/emp/all", middleware.AdminAuth(), admin_controllers.GetAllEmployees)

	router.Put("/api/main/empupdate/:id", middleware.AdminAuth(), admin_controllers.UpdateEmployeeProfile)

	router.Put("/api/main/setsuperadmin/:id", middleware.AdminAuth(), admin_controllers.CreateSuperAdmin)

	router.Delete("/api/main/empdelete/:id", middleware.AdminAuth(), admin_controllers.DeleteEmployee)

	router.Get("/api/main/emp/:id", middleware.AdminAuth(), admin_controllers.GetEmployeeProfile)

}
