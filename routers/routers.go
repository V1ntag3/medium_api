package routers

import (
	"medium_api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// create routers

	// auth routers
	app.Post("/api/auth/register", controllers.Register)
	app.Post("/api/auth/login", controllers.Login)
	app.Get("/api/auth/profile", controllers.Profile)
	app.Post("/api/auth/logout", controllers.Logout)
	app.Delete("/api/auth/profile", controllers.Delete)

}
