package controllers

import "github.com/gofiber/fiber/v2"

func Health(c *fiber.Ctx) error {
	// TODO: support to check if endpoints are working fine
	return c.JSON(fiber.Map{
		"status": "OK",
	})
}
