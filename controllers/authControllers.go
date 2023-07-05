package controllers

import (
	"medium_api/database"
	"medium_api/models"
	"medium_api/utilities"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello World!!")
}

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
