package route

import (
	"go-lms/config"
	"go-lms/handlers"
	"go-lms/repository"
	"go-lms/service"

	"github.com/gofiber/fiber/v2"
)

func NewLibraryMemberRoute(app fiber.Router) {
	db := *config.Database
	libraryMemberRepository := repository.NewLibraryMember(&db)
	libraryMemberService := service.NewLibraryMember(libraryMemberRepository)
	libraryMemberHandler := handlers.NewLibraryMemberHandler(libraryMemberService)
	app.Get("library_members", func(c *fiber.Ctx) error {
		return libraryMemberHandler.GetLibraryMembers(c)
	})
	app.Get("library_members/:id", func(c *fiber.Ctx) error {
		return libraryMemberHandler.GetLibraryMemberById(c)
	})
	app.Post("library_members/searchByName", func(c *fiber.Ctx) error {
		return libraryMemberHandler.GetLibraryMemberByName(c)
	})
	app.Post("library_members", func(c *fiber.Ctx) error {
		return libraryMemberHandler.AddLibrarryMember(c)
	})
	app.Put("library_members/:id", func(c *fiber.Ctx) error {
		return libraryMemberHandler.UpdateLibrarryMember(c)
	})
}
