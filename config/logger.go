package config

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	LoggerConfig logger.Config
	LogFile      *os.File
)

func InitLogger() {
	file, err := os.OpenFile("./app_logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	LoggerConfig = logger.Config{
		Output: file,
		Format: time.Now().Format("2006-01-02") + " ${time}|[${ip}]|${method}|${path}|${status}|${latency}|${error} \n",
	}

	LogFile = file
}
