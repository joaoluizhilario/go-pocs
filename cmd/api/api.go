package api

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joaoluizhilario/go-cockroach-poc/docs"
	"github.com/joaoluizhilario/go-cockroach-poc/internal/handlers"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"go.uber.org/fx"
)

//	@title			Supermarket API
//	@version		2.0
//	@description	Supermarket API for studies.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:3000
// @BasePath	/
// @schemes	http
func NewFiberAPI(lc fx.Lifecycle, productAPIHandler *handlers.ProductAPIHandler, orderAPIHandler *handlers.OrderAPIHandler) *fiber.App {

	app := fiber.New()

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Get("/product", productAPIHandler.ListProducts)
	app.Post("/product", productAPIHandler.CreateProduct)
	app.Get("/order", orderAPIHandler.ListOrders)
	app.Post("/order", orderAPIHandler.CreateOrder)

	app.Listen(":3000")

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Runnig fiber server over fx on 8080")
			go app.Listen(":8080")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}
