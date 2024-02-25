package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joaoluizhilario/go-cockroach-poc/internal/models"
	"go.uber.org/fx"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresGormService struct {
	Db *gorm.DB
}

func (g *PostgresGormService) Connect() error {
	return nil
}

func (g *PostgresGormService) Disconnect() error {
	cockroachDB, err := g.Db.DB()
	if err != nil {
		return err
	}
	return cockroachDB.Close()
}

func CreatePostgresConnection(lc fx.Lifecycle) (*PostgresGormService, error) {

	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	DB := PostgresGormService{
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
