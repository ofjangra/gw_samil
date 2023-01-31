package main

import (
	// "fmt"
	"fmt"
	"log"
	"os"
	"path/filepath"

	// "path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ofjangra/gwonline/routes"
)

func main() {
	port := os.Getenv("PORT")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
	}))

	app.Static("/", "./dist")

	// app.Static("/", "./build_SP_ATMBD")

	routes.Router_Samil(app)
	routes.Router_Employees(app)
	routes.Router_Vehicle(app)
	routes.Router_ServicePricing(app)

	indexPath, indexPathErr := filepath.Abs("./dist/index.html")
	samilPath, samilPathErr := filepath.Abs("./dist/samil.html")
	adminPath, adminPathErr := filepath.Abs("./dist/admin.html")
	// sp_automobile_dealerPath, sp_automobile_dealerPathErr := filepath.Abs("./build_SP_ATMBD/index.html")

	if samilPathErr != nil {
		fmt.Println(samilPathErr)
	}

	if indexPathErr != nil {
		fmt.Println(indexPathErr)
	}

	if adminPathErr != nil {
		fmt.Println(adminPathErr)
	}

	// if sp_automobile_dealerPathErr != nil {
	// 	fmt.Println(adminPathErr)
	// }

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile(indexPath)
	})

	app.Get("/samil/*", func(c *fiber.Ctx) error {
		return c.SendFile(samilPath)
	})

	app.Get("/admin/*", func(c *fiber.Ctx) error {
		return c.SendFile(adminPath)
	})
	// app.Get("/automobile_dealer/*", func(c *fiber.Ctx) error {
	// 	return c.SendFile(sp_automobile_dealerPath)
	// })

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile(indexPath)
	})

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
