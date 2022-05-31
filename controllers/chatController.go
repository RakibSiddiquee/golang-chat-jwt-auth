package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func StoreMessage(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Check if authenticated
	if !IsLoggedIn {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	return c.JSON(IsLoggedIn)
}
