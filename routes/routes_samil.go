package routes

import (
	"github.com/gofiber/fiber/v2"
	samil_controllers "github.com/ofjangra/gwonline/controllers/samil"
	"github.com/ofjangra/gwonline/middleware"
)

func Router_Samil(app *fiber.App) {

	auth := middleware.Authrequired()
	router := app

	router.Post("/api/samil/signup", samil_controllers.Signup_samil)

	router.Post("/api/samil/signin", samil_controllers.Signin_samil)

	router.Get("/api/samil/logout", samil_controllers.Logout_samil)

	router.Get("/api/samil/u/:id", middleware.AdminAuth(), samil_controllers.GetSamilUserProfile)

	router.Get("/api/samil/users", middleware.AdminAuth(), samil_controllers.GetAllSamilUsers)

	router.Get("/api/samil", auth, samil_controllers.Home_Samil)

}
