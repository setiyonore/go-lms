package main

import (
	"go-lms/config"
	"go-lms/route"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	config.InitLogger()
	loggerConfig := config.LoggerConfig
	app.Use(logger.New(loggerConfig))
	app.Use(limiter.New(config.Limiter))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:3000,http://localhost:3000",
		AllowMethods:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin", // Adapt if required
		AllowCredentials: true,
	}))
	config.Connect()
	route.Init(app)
	log.Fatal(app.Listen(":8080"))
}
