package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaoluizhilario/go-cockroach-poc/internal/database"
)

func InitAPI() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
