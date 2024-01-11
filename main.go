package main

import (
	"go-lms/config"
	"go-lms/route"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	app := fiber.New()
	app.Use(limiter.New(config.Limiter))
	config.Connect()
	route.Setup(app)
	log.Fatal(app.Listen(":8080"))
}
