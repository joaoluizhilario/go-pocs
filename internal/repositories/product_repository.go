package repositories

import (
	"github.com/joaoluizhilario/go-cockroach-poc/internal/database"
	"github.com/joaoluizhilario/go-cockroach-poc/internal/models"
)

type ProductRepositoryInterface interface {
	ListProducts() *[]models.Product
	CreateProduct(product *models.Product)
}

type ProductRepository struct {
	DB *database.CockroachGormService
}

func (pr *ProductRepository) ListProducts() *[]models.Product {
	products := []models.Product{}
	pr.DB.Db.Find(&products)
	return &products
}

func (pr *ProductRepository) CreateProduct(product *models.Product) {
	pr.DB.Db.Create(&product)
}

func NewProductRepository(dbService *database.CockroachGormService) *ProductRepository {
	return &ProductRepository{DB: dbService}
}
