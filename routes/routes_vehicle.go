package routes

import (
	"github.com/gofiber/fiber/v2"
	admin_controllers "github.com/ofjangra/gwonline/controllers/admin"
	"github.com/ofjangra/gwonline/middleware"
)

func Router_Vehicle(app *fiber.App) {
	Router := app

	Router.Post("/api/gwvehicles/create", middleware.AdminAuth(), admin_controllers.CreateVehicleEntry)

	Router.Get("/api/gwvehicles/vehicle/:id", middleware.AdminAuth(), admin_controllers.GetVehicleById)

	Router.Get("/api/gwvehicles/vehicles", middleware.AdminAuth(), admin_controllers.GetAllVehicles)
}
