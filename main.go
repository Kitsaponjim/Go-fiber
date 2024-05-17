package main

import (
	r "go-fiber-test/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	r.InetRoutes(app)
	app.Listen(":3000")
}
