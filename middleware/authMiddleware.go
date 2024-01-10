package middleware

import (
	"go-lms/config"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.Secret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(401).JSON(err.Error())
		},
	})
}
