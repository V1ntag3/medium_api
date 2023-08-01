package utilities

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func IsAuthenticadCookie(c *fiber.Ctx, SecretKey string) (*jwt.Token, error) {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, errors.New("unauthorized")
	}
	return token, nil
}
func IsAuthenticadToken(c *fiber.Ctx, SecretKey string) (*jwt.Token, error) {

	reqToken := c.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")

	if len(splitToken) <= 1 {
		return nil, errors.New("unauthorized")
	}
	reqToken = splitToken[1]

	if !IsAuthorizedToken(reqToken) {
		return nil, errors.New("unauthorized")
	}

	token, err := jwt.ParseWithClaims(reqToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, errors.New("unauthorized")
	}

	return token, nil
}
