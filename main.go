package main

import (
	"go-lms/config"
	"go-lms/route"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.Connect()
	route.Setup(app)
	log.Fatal(app.Listen(":8080"))
}
