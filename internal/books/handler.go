package books

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, service Service) {
	api := app.Group("/api/books")

	api.Get("/", func(c *fiber.Ctx) error {
		books, err := service.ListBooks()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(books)
	})

	api.Get("/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
		}
		book, err := service.GetBook(uint(id))
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
		}
		return c.JSON(book)
	})

	api.Post("/", func(c *fiber.Ctx) error {
		var book Book
		if err := c.BodyParser(&book); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		if err := service.CreateBook(&book); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(book)
	})

	api.Put("/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
		}
		var book Book
		if err := c.BodyParser(&book); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		book.ID = uint(id)
		if err := service.UpdateBook(&book); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(book)
	})

	api.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
		}
		if err := service.DeleteBook(uint(id)); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.SendStatus(204)
	})
}
