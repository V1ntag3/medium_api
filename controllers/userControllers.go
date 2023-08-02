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

	var data map[string]string
	var user models.User

	token, err := utilities.IsAuthenticadToken(c, SecretKey)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	if err = c.BodyParser(&data); err != nil {
		return err
	}

	if err := database.DB.Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	if !utilities.OnlyEmptySpaces(data["name"]) {
		user.Name = data["name"]
	}

	if !utilities.OnlyEmptySpaces(data["surname"]) {
		user.Surname = data["surname"]
	}

	if !utilities.OnlyEmptySpaces(data["about"]) {
		user.About = data["about"]
	}

	if err := database.DB.Save(&user).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(user)

}

// User data
func Profile(c *fiber.Ctx) error {

	var followingsCount int64
	var followersCount int64
	var user models.User

	token, err := utilities.IsAuthenticadToken(c, SecretKey)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	if err := database.DB.Select("id, name, surname, about, date_member, email, image_profile").Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	if err := database.DB.Model(&models.UserFollower{}).Where("follower_id = ?", user.Id).Count(&followersCount).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	if err := database.DB.Model(&models.UserFollower{}).Where("user_id = ?", user.Id).Count(&followingsCount).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	data := map[string]interface{}{
		"id":            user.Id,
		"name":          user.Name,
		"surname":       user.Surname,
		"email":         user.Email,
		"dateMember":    user.DateMember,
		"about":         user.About,
		"image_profile": user.ImageProfile,
		"followers":     followersCount,
		"followings":    followingsCount,
	}

	return c.JSON(data)
}

// Delete user
func Delete(c *fiber.Ctx) error {

	var user models.User

	token, err := utilities.IsAuthenticadToken(c, SecretKey)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	if err := database.DB.Where("id = ?", claims.Issuer).First(&user).Delete(&user).Error; err != nil {
		return fiber.ErrNotFound
	}

	utilities.UnauthorizedToken(token.Raw)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// Rescue followings
func Following(c *fiber.Ctx) error {
	var user models.User

	if err := database.DB.Preload("Following").First(&user, "id = ?", c.Params("id")).Error; err != nil {
		return fiber.ErrNotFound
	}

	return c.JSON(user.Followings)
}

// Rescue followers
func Followers(c *fiber.Ctx) error {
	var user models.User

	if err := database.DB.Preload("Followers").First(&user, "id = ?", c.Params("id")).Error; err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(user.Followers)
}

// Follow a user
func Follow(c *fiber.Ctx) error {

	var user models.User
	var userFollow models.User
	var existingUserFollower models.UserFollower

	token, err := utilities.IsAuthenticadToken(c, SecretKey)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	if err := database.DB.Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		return fiber.ErrNotFound
	}
	if err := database.DB.Where("id = ?", c.Params("id")).First(&userFollow).Error; err != nil {
		return fiber.ErrNotFound
	}

	// Verify if user follow a user
	if err := database.DB.Where("user_id = ? AND follower_id = ?", user.Id, userFollow.Id).First(&existingUserFollower).Error; err == nil {
		return c.JSON(fiber.Map{
			"message": "Você já está seguindo este usuário.",
		})
	}

	// Create a new
	userFollower := models.UserFollower{
		UserID:     user.Id,
		FollowerID: userFollow.Id,
	}

	if err := database.DB.Create(&userFollower).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"message": "User follow with success.",
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
