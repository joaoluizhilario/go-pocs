package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaoluizhilario/go-cockroach-poc/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/product", handlers.ListProducts)
	app.Post("/product", handlers.CreateProduct)
	app.Get("/order", handlers.ListOrders)
	app.Post("/order", handlers.CreateOrder)
}
