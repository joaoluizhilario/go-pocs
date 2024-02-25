package repositories

import (
	"github.com/joaoluizhilario/go-cockroach-poc/internal/database"
	"github.com/joaoluizhilario/go-cockroach-poc/internal/models"
)

type ProductRepositoryInterface interface {
	ListProducts() *[]models.Product
	CreateProduct(product *models.Product)
}

type ProductRepository struct{}

func (pr *ProductRepository) ListProducts() *[]models.Product {
	products := []models.Product{}
	database.DB.Db.Find(&products)
	return &products
}

func (pr *ProductRepository) CreateProduct(product *models.Product) {
	database.DB.Db.Create(&product)
}
