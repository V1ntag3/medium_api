package routers

import (
	"medium_api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// create routers

	app.Post("/api/auth/register", controllers.Register)

}
