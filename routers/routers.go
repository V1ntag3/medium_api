package routers

import (
	"medium_api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// create routers
	// static achives
	app.Static("/uploads", "./uploads")
	// auth routers
	app.Post("/api/auth/register", controllers.Register)
	app.Post("/api/auth/login", controllers.Login)
	app.Get("/api/auth/profile", controllers.Profile)
	app.Post("/api/auth/logout", controllers.Logout)
	app.Delete("/api/auth/profile", controllers.Delete)
	// upload images
	app.Post("/api/imageProfile", controllers.ImageProfileUpload)
	app.Post("/api/imageWallpaper", controllers.ImageWallpaperUpload)
	// articule

	app.Post("/api/articule/create", controllers.CreateArticule)
	app.Get("/api/articules/my", controllers.GetAllArticulesMyUser)
	app.Get("/api/articules/all", controllers.GetAllArticules)
	app.Get("/api/articules/:id", controllers.GetAllArticulespScificUser)

}
