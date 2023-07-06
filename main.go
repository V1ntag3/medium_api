package main

import (
	"log"
	"medium_api/database"
	"medium_api/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect("database.db")

	app := fiber.New()

	app.Use(
		cors.New(
			cors.Config{
				AllowCredentials: true,
			}))

	routers.Setup(app)

	log.Fatal(app.Listen(":8000"))

}
