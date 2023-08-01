package controllers

import (
	"log"
	"medium_api/database"
	"medium_api/models"
	"medium_api/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "iaoiwjdaojwdoiciwoeinow"

func Register(c *fiber.Ctx) error {

	var data map[string]string
	// recue databody
	err := c.BodyParser((&data))

	if err != nil {
		return err
	}
	// validation data
	error_validation := make(map[string]string)

	if utilities.ContainsNumber(data["name"]) || utilities.OnlyEmptySpaces(data["name"]) {
		error_validation["name"] = "Name is invalid"
	}

	if utilities.ContainsNumber(data["surname"]) || utilities.OnlyEmptySpaces(data["surname"]) {
		error_validation["surname"] = "Surname is invalid"
	}

	if !utilities.ValidateEmail(data["email"]) {
		error_validation["email"] = "E-mail is invalid"
	}

	if !utilities.ValidatePassword(data["password"]) {
		error_validation["password"] = "Password is invalid"
	}
	if len(error_validation) != 0 {
		return c.Status(400).JSON(error_validation)
	}

	// hash of password
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Id:         uuid.New().String(),
		Name:       data["name"],
		Surname:    data["surname"],
		Email:      data["email"],
		Password:   password,
		DateMember: utilities.DateTimeNow(),
	}

	err_db := database.DB.Create(&user)

	if err_db.Error != nil {

		if err_db.Error.Error() == "UNIQUE constraint failed: users.email" {

			error_validation["email"] = "E-mail already registered"
			return c.Status(400).JSON(error_validation)

		}
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser((&data))

	if err != nil {
		return err
	}
	log.Print(data)
	var user models.User
	database.DB.Where("email= ?", data["email"]).First(&user)
	if user.Id == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Id,
		ExpiresAt: utilities.DateTimeNowAddHoursUnix(24),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.Status(404).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	// cookie := fiber.Cookie{
	// 	Name:     "jwt",
	// 	Value:    token,
	// 	Expires:  utilities.DateTimeNowAddHours(24),
	// 	HTTPOnly: true,
	// }

	// c.Cookie(&cookie)
	utilities.AuthorizedToken(token)
	return c.JSON(fiber.Map{
		"token":   token,
		"expires": utilities.DateTimeNowAddHours(24),
	})
}

func Logout(c *fiber.Ctx) error {

	token, err := utilities.IsAuthenticadToken(c, SecretKey)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	utilities.UnauthorizedToken(token.Raw)

	// cookieLogout := fiber.Cookie{
	// 	Name:     "jwt",
	// 	Value:    "",
	// 	Expires:  utilities.DateTimeNowAddHours(-24),
	// 	HTTPOnly: true,
	// }

	// c.Cookie(&cookieLogout)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}
