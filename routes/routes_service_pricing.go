package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	pricing_controllers "github.com/ofjangra/gwonline/controllers/pricings"
	middleware "github.com/ofjangra/gwonline/middleware"
)

func Router_ServicePricing(app *fiber.App) {

	router := app

	router.Post("/api/sp/create_entry", middleware.AdminAuth(), pricing_controllers.CreatePricingEntry)

	router.Get("/api/sp/passing", pricing_controllers.GetPricingEntry_passing)

	router.Get("/api/sp/seating", pricing_controllers.GetPricingEntry_seating)

	router.Post("/api/sp/partner", middleware.AdminAuth(), pricing_controllers.CreateEntry_SP_Partner)

	router.Get("/api/sp/partner", pricing_controllers.GetEntry_SP_Partner)

	// router.Put("/api/sp/automobile_dealer/:id", middleware.AdminAuth(), pricing_controllers.UpdateEntry_SP_Loading)

}
