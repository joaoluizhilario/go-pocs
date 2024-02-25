package main

import (
	"github.com/joaoluizhilario/go-cockroach-poc/cmd/api"
	"github.com/joaoluizhilario/go-cockroach-poc/internal/database"
	"github.com/joaoluizhilario/go-cockroach-poc/internal/handlers"
	"github.com/joaoluizhilario/go-cockroach-poc/internal/repositories"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(database.CreateCockroachConnection),
		fx.Provide(repositories.NewOrderRepository),
		fx.Provide(repositories.NewProductRepository),
		fx.Provide(handlers.NewOrderAPIHandler),
		fx.Provide(handlers.NewProductAPIHandler),
		fx.Invoke(api.NewFiberAPI),
	).Run()
}
