package controllers

import (
	"medium_api/database"
	"medium_api/models"
	"medium_api/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// User data
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

	if err := database.DB.Select("id, name, surname, about, date_member, email, image_profile").Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		return fiber.ErrNotFound
	}

	var followingsCount int64
	var followersCount int64

	if err := database.DB.Model(&models.UserFollower{}).Where("follower_id = ?", user.Id).Count(&followersCount).Error; err != nil {
		return fiber.ErrNotFound
	}

	if err := database.DB.Model(&models.UserFollower{}).Where("user_id = ?", user.Id).Count(&followingsCount).Error; err != nil {
		return fiber.ErrNotFound
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

// User data by id
func UserDataById(c *fiber.Ctx) error {

	_, err := utilities.IsAuthenticadToken(c, SecretKey)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	var user models.User

	if err := database.DB.Select("id, name, surname, about, date_member, email, image_profile").Where("id = ?", c.Params("id")).First(&user).Error; err != nil {
		return fiber.ErrNotFound
	}

	var followingsCount int64
	var followersCount int64

	if err := database.DB.Model(&models.UserFollower{}).Where("follower_id = ?", user.Id).Count(&followersCount).Error; err != nil {
		return fiber.ErrNotFound
	}

	if err := database.DB.Model(&models.UserFollower{}).Where("user_id = ?", user.Id).Count(&followingsCount).Error; err != nil {
		return fiber.ErrNotFound
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

	if err := c.BodyParser(&data); err != nil {
		return fiber.ErrInternalServerError
	}

	var user models.User

	if err := database.DB.Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		return fiber.ErrNotFound
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

// Delete user
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

	token, err := utilities.IsAuthenticadToken(c, SecretKey)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	if err := database.DB.Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		return fiber.ErrNotFound
	}

	var userFollow models.User

	if err := database.DB.Where("id = ?", c.Params("id")).First(&userFollow).Error; err != nil {
		return fiber.ErrNotFound
	}

	var existingUserFollower models.UserFollower

	// Verify if user follow a user
	if err := database.DB.Where("user_id = ? AND follower_id = ?", user.Id, userFollow.Id).First(&existingUserFollower).Error; err == nil {
		return c.JSON(fiber.Map{
			"message": "user is already being followed",
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
		"message": "user successfully followed",
	})

}

// Unfollow user
func UnFollow(c *fiber.Ctx) error {

	token, err := utilities.IsAuthenticadToken(c, SecretKey)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	if err := database.DB.Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		return fiber.ErrNotFound
	}

	var userFollow models.User

	if err := database.DB.Where("id = ?", c.Params("id")).First(&userFollow).Error; err != nil {
		return fiber.ErrNotFound
	}

	var existingUserFollower models.UserFollower

	if err := database.DB.Where("user_id = ? AND follower_id = ?", user.Id, userFollow.Id).First(&existingUserFollower).Error; err != nil {
		return c.JSON(fiber.Map{
			"message": "user is not being followed",
		})
	}

	if err := database.DB.Delete(&existingUserFollower).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"message": "Usuário deixado de seguir com sucesso.",
	})
}
