package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaoluizhilario/go-cockroach-poc/internal/handlers"
	"github.com/joaoluizhilario/go-cockroach-poc/internal/repositories"
)

func setupRoutes(app *fiber.App) {

	productApiHandler := handlers.ProductAPIHandler{
		Repository: new(repositories.ProductRepository),
	}

	orderApiHandler := handlers.OrderAPIHandler{
		Repository: new(repositories.OrderRepository),
	}

	app.Get("/product", productApiHandler.ListProducts)
	app.Post("/product", productApiHandler.CreateProduct)
	app.Get("/order", orderApiHandler.ListOrders)
	app.Post("/order", orderApiHandler.CreateOrder)
}
