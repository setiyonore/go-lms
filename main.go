package main

import (
	"go-lms/config"
	"go-lms/route"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	config.InitLogger()
	loggerConfig := config.LoggerConfig
	app.Use(logger.New(loggerConfig))
	app.Use(limiter.New(config.Limiter))
	config.Connect()
	route.Init(app)
	log.Fatal(app.Listen(":8080"))
}
