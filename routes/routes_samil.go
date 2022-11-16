package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ofjangra/gwonline/controllers"
	"github.com/ofjangra/gwonline/middleware"
)

func Router_Samil(app *fiber.App) {

	auth := middleware.Authrequired()
	router := app

	router.Post("/api/samil/signup", controllers.Signup_samil)

	router.Post("/api/samil/signin", controllers.Signin_samil)

	router.Get("/api/samil/logout", controllers.Logout_samil)

	router.Get("/api/samil", auth, controllers.Home_Samil)

}
