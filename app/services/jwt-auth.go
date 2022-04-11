package services

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// type User struct {
// 	Username string `josn:"username"`
// 	Password string `josn:"password"`
// }

func Login(c *fiber.Ctx) error {
	// var user User
	// c.BodyParser(&user)

	// if user.Password != "eunimart" {
	// 	return c.SendStatus(fiber.StatusUnauthorized)
	// }

	claims := jwt.MapClaims{
		"username": "dilip",
		"admin":    true,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed_token, err := token.SignedString([]byte("secretkeyan"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": false, "message": "Error in generating authorization token", "error": err.Error()})
	}
	return c.JSON(fiber.Map{"token": signed_token})
}
