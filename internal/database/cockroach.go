package database

import (
	"context"
	"log"
	"os"

	"github.com/joaoluizhilario/go-cockroach-poc/internal/models"
	"go.uber.org/fx"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type CockroachGormService struct {
	Db *gorm.DB
}

func (g *CockroachGormService) Connect() error {
	return nil
}

func (g *CockroachGormService) Disconnect() error {
	cockroachDB, err := g.Db.DB()
	if err != nil {
		return err
	}
	return cockroachDB.Close()
}

func CreateCockroachConnection(lc fx.Lifecycle) (*CockroachGormService, error) {

	dsn := os.Getenv("CRDB_DATABASE_URL") + "&application_name=$ docs_simplecrud_gorm"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	DB := CockroachGormService{
		Db: db,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			if err != nil {
				return err
			}
			log.Println("running migrations")
			db.AutoMigrate(&models.Product{})
			db.AutoMigrate(&models.ProductItem{})
			db.AutoMigrate(&models.Order{})
			return nil
		},
	})

	return &DB, nil
}
