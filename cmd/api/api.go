package api

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/joaoluizhilario/go-cockroach-poc/docs"
	"github.com/joaoluizhilario/go-cockroach-poc/internal/database"
	fiberSwagger "github.com/swaggo/fiber-swagger"
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

func InitAPI() {
	database.ConnectDb()

	app := fiber.New()

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	setupRoutes(app)

	app.Listen(":3000")
}
