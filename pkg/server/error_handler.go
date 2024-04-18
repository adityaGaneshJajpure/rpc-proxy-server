package server

import (
	"github.com/gofiber/fiber/v2"
)

var ErrorHandler = func(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	var message interface{}
	if code == fiber.StatusOK {
		message = struct {
			Message string `json:"message"`
		}{err.Error()}
	} else {
		message = struct {
			Error string `json:"error"`
		}{err.Error()}
	}
	return c.Status(code).JSON(message)
}
