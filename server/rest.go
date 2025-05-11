package server

import (
	"go-fiber-grpc/database"
	"go-fiber-grpc/internal/books"

	"github.com/gofiber/fiber/v2"
)

func StartREST() {
	db := database.Connect()
	repo := books.NewRepository(db)
	service := books.NewService(repo)

	app := fiber.New()
	// Define REST routes and use service to handle them
	books.RegisterRoutes(app, service)

	app.Listen(":3001")
}
