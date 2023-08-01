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

	// Middleware CORS para permitir solicitações apenas do domínio http://localhost:3000
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin, Authorization",
		AllowOrigins: "http://localhost:3000",
		// AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	routers.Setup(app)

	log.Fatal(app.Listen(":8000"))

}
