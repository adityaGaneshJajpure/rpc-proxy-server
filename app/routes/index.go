package routes

import (
	"github.com/adityaGaneshJajpure/rpc-proxy-server/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Post("/", controllers.Proxy)
	app.Get("/_healthz", controllers.Health)
}
