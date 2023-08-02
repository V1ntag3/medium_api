package controllers

import (
	"log"
	"medium_api/database"
	"medium_api/models"
	"medium_api/utilities"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func CreateArticle(c *fiber.Ctx) error {

	token, err := utilities.IsAuthenticadToken(c, SecretKey)

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
	if utilities.OnlyEmptySpaces(data["abstract"]) {
		error_validation["abstract"] = "Abstract is invalid"
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
	// parse incomming image file
	uniqueId := uuid.New()

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	err = utilities.Base64ToImage(data["bannerImage"], "./uploads/articles/"+filename+".jpg")

	if err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "Error converting image to Base64 ", "data": nil})
	}

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	var idValue = claims.Issuer

	article := models.Article{
		Id:          uuid.NewString(),
		Abstract:    data["abstract"],
		Title:       data["title"],
		Subtile:     data["subtitle"],
		Text:        data["text"],
		BannerImage: "/uploads/articles/" + filename + ".jpg",
		CreateTime:  utilities.DateTimeNow(),

		UserId: idValue,
	}

	err_db := database.DB.Create(&article)

	if err_db.Error != nil {

		if err_db.Error.Error() == "UNIQUE constraint failed: users.email" {

			error_validation["email"] = "E-mail already registered"
			return c.Status(400).JSON(error_validation)

		}
	}
	return c.JSON(article)

}

// Return all articules with page and limit
func GetAllArticles(c *fiber.Ctx) error {
	// used in filter
	pageNumber, _ := strconv.Atoi(c.Query("page", "1"))
	limitNumber, _ := strconv.Atoi(c.Query("limit", "10"))
	offset := (pageNumber - 1) * limitNumber

	var articles []models.Article

	if err := database.DB.Offset(offset).Limit(limitNumber).Find(&articles).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(articles)

}

func GetAllArticlespScificUser(c *fiber.Ctx) error {

	var users models.User
	database.DB.Model(&models.User{}).Preload("Articles").Where("id = ?", c.Params("id")).First(&users)

	return c.JSON(users.Articles)

}

func GetAllArticlesMyUser(c *fiber.Ctx) error {

	// if user is authenticad this method rescue token
	token, err := utilities.IsAuthenticadToken(c, SecretKey)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	// filter and send articles
	var users models.User
	database.DB.Model(&models.User{}).Preload("Articles").Where("id = ?", claims.Issuer).First(&users)

	return c.JSON(users.Articles)

}
