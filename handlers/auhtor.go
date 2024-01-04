package handlers

import (
	"go-lms/config"
	"go-lms/entities"

	"github.com/gofiber/fiber/v2"
)

func GetAuthor(ctx *fiber.Ctx) error {
	var authors []entities.Author
	config.Database.Find(&authors)
	return ctx.Status(200).JSON(authors)
}
