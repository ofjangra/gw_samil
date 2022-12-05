package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ofjangra/gwonline/routes"
)

func main() {
	port := os.Getenv("PORT")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
	}))

	app.Static("/", "./dist")

	routes.Router_Samil(app)
	routes.Router_Employees(app)
	routes.Router_Vehicle(app)

	indexPath, indexPathErr := filepath.Abs("./dist/index.html")
	samilPath, samilPathErr := filepath.Abs("./dist/samil.html")
	adminPath, adminPathErr := filepath.Abs("./dist/admin.html")

	if samilPathErr != nil {
		fmt.Println(samilPathErr)
	}
	if indexPathErr != nil {
		fmt.Println(indexPathErr)
	}
	if adminPathErr != nil {
		fmt.Println(adminPathErr)
	}
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile(indexPath)
	})
	app.Get("/samil/*", func(c *fiber.Ctx) error {
		return c.SendFile(samilPath)
	})

	app.Get("/admin/*", func(c *fiber.Ctx) error {
		return c.SendFile(adminPath)
	})

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile(indexPath)
	})

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
