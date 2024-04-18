package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/adityaGaneshJajpure/rpc-proxy-server/app/routes"
	"github.com/adityaGaneshJajpure/rpc-proxy-server/pkg/constants"
	"github.com/adityaGaneshJajpure/rpc-proxy-server/pkg/server"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	if len(os.Getenv(constants.RPC_PROVIDER)) == 0 {
		log.Panic("RPC_PROVIDER env variable not set")
	}

	app := fiber.New(fiber.Config{
		Prefork:      true,
		ErrorHandler: server.ErrorHandler,
	})

	routes.RegisterRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
