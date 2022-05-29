package main

import (
	"github.com/RakibSiddiquee/golang-chat-jwt-auth/database"
	"github.com/RakibSiddiquee/golang-chat-jwt-auth/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
