package controllers

import (
	"encoding/json"

	"github.com/adityaGaneshJajpure/rpc-proxy-server/pkg/dto"
	"github.com/adityaGaneshJajpure/rpc-proxy-server/pkg/proxy"
	"github.com/gofiber/fiber/v2"
)

func Proxy(ctx *fiber.Ctx) error {
	var request dto.RPCRequest
	if err := json.Unmarshal(ctx.Body(), &request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid Request",
		})
	}

	response, err := proxy.ProxyHandler(request)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(response)
}
