package routes

import (
	"github.com/RakibSiddiquee/golang-chat-jwt-auth/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
}