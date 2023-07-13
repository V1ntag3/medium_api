package controllers

import (
	"log"
	"medium_api/database"
	"medium_api/models"
	"medium_api/utilities"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CreateArticule(c *fiber.Ctx) error {
	// validate user
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var data map[string]string
	// recue databody
	err = c.BodyParser((&data))

	if err != nil {
		return err
	}
	// validation data
	error_validation := make(map[string]string)

	if utilities.OnlyEmptySpaces(data["title"]) {
		error_validation["title"] = "Title is invalid"
	}

	if utilities.OnlyEmptySpaces(data["subtitle"]) {
		error_validation["subtitle"] = "Subtitle is invalid"
	}

	if utilities.OnlyEmptySpaces(data["text"]) {
		error_validation["text"] = "Text is invalid"
	}

	if utilities.OnlyEmptySpaces(data["bannerImage"]) {
		error_validation["bannerImage"] = "Image is invalid"
	}
	if len(error_validation) != 0 {
		return c.Status(400).JSON(error_validation)
	}

	// change base 64 to image

	idValue, err := strconv.ParseUint(claims.Issuer, 10, 32)

	articule := models.Articule{
		Title:       data["title"],
		Subtile:     data["subtitle"],
		Text:        data["text"],
		BannerImage: data["bannerImage"],
		CreateTime:  utilities.DateTimeNow(),
		UserId:      uint(idValue),
	}

	err_db := database.DB.Create(&articule)

	if err_db.Error != nil {

		if err_db.Error.Error() == "UNIQUE constraint failed: users.email" {

			error_validation["email"] = "E-mail already registered"
			return c.Status(400).JSON(error_validation)

		}
	}

	return c.JSON(articule)

}

func GetAllArticules(c *fiber.Ctx) error {
	// validate user
	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	// filter and send articules
	var articules []models.Articule
	result := database.DB.Find(&articules)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return c.JSON(articules)

}

func GetAllArticulespScificUser(c *fiber.Ctx) error {
	// validate user
	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	id := c.Params("id")
	// filter and send articules
	var users models.User
	database.DB.Model(&models.User{}).Preload("Articules").Where("id = ?", id).First(&users)
	return c.JSON(users.Articules)

}

func GetAllArticulesMyUser(c *fiber.Ctx) error {
	// validate user
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	// filter and send articules
	var users models.User
	database.DB.Model(&models.User{}).Preload("Articules").Where("id = ?", claims.Issuer).First(&users)
	return c.JSON(users.Articules)

}
