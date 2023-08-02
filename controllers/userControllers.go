package controllers

import (
	"medium_api/database"
	"medium_api/models"
	"medium_api/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Update name, surname and about
func UpdateUser(c *fiber.Ctx) error {
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

	if err = c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	if !utilities.OnlyEmptySpaces(data["name"]) {
		user.Name = data["name"]
	}

	if !utilities.OnlyEmptySpaces(data["surname"]) {
		user.Surname = data["surname"]
	}

	if !utilities.OnlyEmptySpaces(data["about"]) {
		user.About = data["about"]
	}

	database.DB.Save(&user)

	return c.JSON(user)

}
func Profile(c *fiber.Ctx) error {

	token, err := utilities.IsAuthenticadToken(c, SecretKey)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Delete(c *fiber.Ctx) error {

	token, err := utilities.IsAuthenticadToken(c, SecretKey)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user).Delete(&user)

	cookieLogout := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  utilities.DateTimeNowAddHours(-24),
		HTTPOnly: true,
	}

	c.Cookie(&cookieLogout)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
func Following(c *fiber.Ctx) error {
	var user models.User
	if err := database.DB.Preload("Following").First(&user, "id = ?", c.Params("id")).Error; err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(user.Following)
}

func Followers(c *fiber.Ctx) error {
	var user models.User
	if err := database.DB.Preload("Followers").First(&user, "id = ?", c.Params("id")).Error; err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(user.Followers)
}

func Follow(c *fiber.Ctx) error {
	//  Pega usuario
	token, err := utilities.IsAuthenticadToken(c, SecretKey)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)
	//
	var userFollow models.User
	database.DB.Where("id = ?", c.Params("id")).First(&userFollow)

	if userFollow.Id == "" {
		return fiber.ErrNotFound
	}

	// Verifica se o usuário já está seguindo
	var existingUserFollower models.UserFollower
	database.DB.Where("user_id = ? AND follower_id = ?", user.Id, userFollow.Id).First(&existingUserFollower)

	if existingUserFollower.UserID != "" {
		// O usuário já está seguindo, não é necessário adicionar novamente
		return c.JSON(fiber.Map{
			"message": "Você já está seguindo este usuário.",
		})
	}

	// Cria um novo relacionamento na tabela de relacionamento
	userFollower := models.UserFollower{
		UserID:     user.Id,
		FollowerID: userFollow.Id,
	}
	if err := database.DB.Create(&userFollower).Error; err != nil {
		return c.JSON(fiber.Map{
			"error": "Erro ao seguir o usuário.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Usuário seguido com sucesso.",
	})

}

func UnFollow(c *fiber.Ctx) error {
	//  Pega usuario
	token, err := utilities.IsAuthenticadToken(c, SecretKey)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)
	//
	var userFollow models.User
	database.DB.Where("id = ?", c.Params("id")).First(&userFollow)

	if userFollow.Id == "" {
		return fiber.ErrNotFound
	}

	var existingUserFollower models.UserFollower

	err2 := database.DB.Where("user_id = ? AND follower_id = ?", user.Id, userFollow.Id).First(&existingUserFollower).Error
	if err2 != nil {
		// O usuário não está seguindo, não é necessário remover
		return c.JSON(fiber.Map{
			"message": "Você ainda não está seguindo este usuário.",
		})
	}

	// Remove o relacionamento da tabela de relacionamento
	if err := database.DB.Delete(&existingUserFollower).Error; err != nil {
		return c.JSON(fiber.Map{
			"error": "Erro ao deixar de seguir o usuário.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Usuário deixado de seguir com sucesso.",
	})
}
