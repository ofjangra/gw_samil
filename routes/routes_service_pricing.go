package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	pricing_controllers "github.com/ofjangra/gwonline/controllers/pricings"
	middleware "github.com/ofjangra/gwonline/middleware"
)

func Router_ServicePricing(app *fiber.App) {

	router := app

	router.Post("/api/sp/automobile_dealer", middleware.AdminAuth(), pricing_controllers.CreateEntry_SP_AutomobileDealer)

	router.Get("/api/sp/automobile_dealer", pricing_controllers.GetEntry_SP_AutomobileDealer)

	router.Put("/api/sp/automobile_dealer/:id", middleware.AdminAuth(), pricing_controllers.UpdateEntry_SP_AutomobileDealer)

}
