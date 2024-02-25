package repositories

import (
	"fmt"

	"github.com/joaoluizhilario/go-cockroach-poc/internal/database"
	"github.com/joaoluizhilario/go-cockroach-poc/internal/models"
)

type OrderRepository struct {
}

type OrderRepositoryInterface interface {
	ListOrders() *[]models.Order
	CreateOrder(order *models.Order) (*models.Order, error)
}

func (or *OrderRepository) ListOrders() *[]models.Order {
	orders := []models.Order{}
	database.DB.Db.Preload("Items").Preload("Items.Product").Find(&orders)
	return &orders
}

func (or *OrderRepository) CreateOrder(order *models.Order) (*models.Order, error) {
	tx := database.DB.Db.Begin()

	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	orderNew := models.Order{
		CustomerName: order.CustomerName,
		Items:        []models.ProductItem{},
	}

	if err := tx.Create(&orderNew).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, item := range order.Items {
		var product models.Product

		if err := tx.Model(models.Product{ID: item.ProductID}).First(&product).Error; err != nil {
			fmt.Println("Product not found")
			tx.Rollback()
			return nil, err
		}

		orderItem := models.ProductItem{
			OrderID:   orderNew.ID,
			ProductID: product.ID,
			Quantity:  item.Quantity,
			UnitPrice: product.Price,
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		order.Items = append(orderNew.Items, orderItem)
	}

	tx.Commit()

	return &orderNew, nil
}
