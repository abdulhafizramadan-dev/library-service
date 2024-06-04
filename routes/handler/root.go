package handler

import "github.com/gofiber/fiber/v2"

// RootHandler is responsible to become root endpoint of service
func RootHandler(c *fiber.Ctx) error {
	return c.SendString("Welcome to Library Service API!")
}
