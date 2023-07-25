package main

import (
	"github.com/stuttgart-things/machineShop-api/api"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/version", api.GetVersion)
	app.Post("/vm", api.CreateVM)
}

curl -X POST -H "Content-Type: application/json" --data "{\"name\": \"Angels and Demons\", \"name\": \"Angels and Demons\", \"os\": \"Dan Brown\"}" http://localhost:3000/vm
