package handler

import (
	"github.com/gofiber/fiber/v2"
)

// PingHandler is responsible for checking the service
func PingHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "pong",
	})
}
