package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ofjangra/gwonline/controllers"
	"github.com/ofjangra/gwonline/middleware"
)

func Router(app *fiber.App) {

	auth := middleware.Authrequired()
	router := app

	router.Post("/api/signup", controllers.Signup)

	router.Post("/api/signin", controllers.Signin)

	router.Get("/api/logout", controllers.Logout)

	router.Get("/api", auth, controllers.Home)

}
