package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joaoluizhilario/go-cockroach-poc/internal/database"
	"github.com/joaoluizhilario/go-cockroach-poc/internal/models"
)

// GetProducts ... Get all products
//
//	@Summary		Get all products
//	@Description	get all products
//	@Tags			Products
//	@Success		200	{array}		models.Product
//	@Failure		404	{object}	object
//	@Router			/product [get]
func ListProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.DB.Db.Find(&products)

	return c.Status(fiber.StatusOK).JSON(products)
}

// CreateProduct ... Create Product
//
//	@Summary		Create new product based on parameters
//	@Description	Create new product
//	@Tags			Products
//	@Accept			json
//	@Param			product	body		models.Product	true	"Product Data"
//	@Success		200		{object}	object
//	@Failure		400,500	{object}	object
//	@Router			/product [post]
func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&product)

	return c.Status(fiber.StatusOK).JSON(product)
}

// CreateOrder ... Create Order
//
//	@Summary		Create new order based on parameters
//	@Description	Create new order
//	@Tags			Orders
//	@Accept			json
//	@Param			order	body		models.Order	true	"Order Data"
//	@Success		200		{object}	object
//	@Failure		400,500	{object}	object
//	@Router			/order [post]
func CreateOrder(c *fiber.Ctx) error {
	orderBody := new(models.Order)

	if err := c.BodyParser(orderBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	tx := database.DB.Db.Begin()

	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	order := models.Order{
		CustomerName: orderBody.CustomerName,
		Items:        []models.ProductItem{},
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range orderBody.Items {
		var product models.Product

		if err := tx.Model(models.Product{ID: item.ProductID}).First(&product).Error; err != nil {
			fmt.Println("Product not found")
			tx.Rollback()
			return err
		}

		orderItem := models.ProductItem{
			OrderID:   order.ID,
			ProductID: product.ID,
			Quantity:  item.Quantity,
			UnitPrice: product.Price,
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			return err
		}

		order.Items = append(order.Items, orderItem)
	}

	tx.Commit()

	return c.Status(fiber.StatusOK).JSON(orderBody)
}

// GetOrders ... Get all orders
//
//	@Summary		Get all orders
//	@Description	get all orders hot reload
//	@Tags			Orders
//	@Success		200	{array}		models.Order
//	@Failure		404	{object}	object
//	@Router			/order [get]
func ListOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.DB.Db.Preload("Items").Preload("Items.Product").Find(&orders)

	return c.Status(fiber.StatusOK).JSON(orders)
}
