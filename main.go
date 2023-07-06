package main

import (
	"log"
	"medium_api/database"
	"medium_api/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect("database.db")

	app := fiber.New()

	routers.Setup(app)

	log.Fatal(app.Listen(":8000"))

}
