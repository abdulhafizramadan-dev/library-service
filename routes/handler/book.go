package handler

import (
	"github.com/gofiber/fiber/v2"
	"gojek.com/abdul/prebootcamp/database"
	"gojek.com/abdul/prebootcamp/model"
)

func AddBookHandler(ctx *fiber.Ctx) error {
	book := new(model.Book)
	if err := ctx.BodyParser(book); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	database.DB.Db.Create(&book)

	return ctx.Status(fiber.StatusCreated).JSON(
		fiber.Map{
			"id":            book.ID,
			"title":         book.Title,
			"price":         book.Price,
			"publishedDate": book.PublishedDate,
			"message":       "Book successfully added to the library.",
		})
}

func GetAllBookHandler(ctx *fiber.Ctx) error {
	var books []model.Book
	database.DB.Db.Find(&books)

	var bookResponse []model.BookResponse
	for _, book := range books {
		bookResponse = append(bookResponse, book.MapToResponse())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"books": bookResponse,
	})
}

func GetBookByIdHandler(ctx *fiber.Ctx) error {
	var book model.Book
	id := ctx.Params("id")
	database.DB.Db.Find(&book, "id = ?", id)

	if book.ID == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(book.MapToResponse())
}

func UpdateBookHandler(ctx *fiber.Ctx) error {
	var book model.Book
	id := ctx.Params("id")

	database.DB.Db.Find(&book, "id = ?", id)

	if book.ID == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	var updateBookRequest model.UpdateBookRequest

	if err := ctx.BodyParser(&updateBookRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	book.Title = updateBookRequest.Title

	database.DB.Db.Save(&book)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      book.ID,
		"title":   book.Title,
		"message": "Book title successfully updated.",
	})
}

func DeleteBookHandler(ctx *fiber.Ctx) error {
	var book model.Book
	id := ctx.Params("id")

	database.DB.Db.Find(&book, "id = ?", id)

	if book.ID == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	err := database.DB.Db.Delete(&book).Error

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":      book.ID,
		"message": "Book successfully deleted.",
	})
}
