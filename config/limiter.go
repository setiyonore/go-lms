package config

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/limiter"
)

var Limiter = limiter.Config{
	Max:        20,
	Expiration: 30 * time.Second,
}
