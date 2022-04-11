package main

import (
	"fmt"
	"prodapi/app/routes"
	"prodapi/app/services"
	"prodapi/app/utils"

	"github.com/gofiber/fiber/v2"

	jwt "github.com/gofiber/jwt/v3"
)

func main() {
	app := fiber.New()
	utils.ConnectToDb()
	app.Get("/get_token", services.Login)
	//authorization middleware
	app.Use(jwt.New(jwt.Config{
		SigningKey: []byte("secretkeyan"),
	}))
	routes.RoutesRegistry(app)
	app.Listen(fmt.Sprintf("%s:%s", utils.Config("HOST", "localhost"), utils.Config("PORT", "5001")))
}
