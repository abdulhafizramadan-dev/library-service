package routes

import (
	"github.com/gofiber/fiber/v2"
	"gojek.com/abdul/prebootcamp/routes/handler"
)

// SetupRoutes is responsible to managing the routes of entire service
func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.RootHandler)
	app.Get("/ping", handler.PingHandler)
	app.Get("/healthz", handler.HealthzHandler)

	book := app.Group("/books")
	book.Get("/", handler.GetAllBookHandler)
	book.Post("/", handler.AddBookHandler)
	book.Get("/:id", handler.GetBookByIdHandler)
	book.Put("/:id", handler.UpdateBookHandler)
	book.Delete("/:id", handler.DeleteBookHandler)
}
