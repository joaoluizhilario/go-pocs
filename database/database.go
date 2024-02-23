package database

import (
	"log"
	"os"

	"github.com/joaoluizhilario/go-cockroach-poc/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {

	dsn := os.Getenv("CRDB_DATABASE_URL") + "&application_name=$ docs_simplecrud_gorm"

	// dsn := fmt.Sprintf(
	// 	"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Sao_Paulo",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// )

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.ProductItem{})
	db.AutoMigrate(&models.Order{})

	DB = Dbinstance{
		Db: db,
	}
}
