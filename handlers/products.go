package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaoluizhilario/go-cockroach-poc/database"
	"github.com/joaoluizhilario/go-cockroach-poc/models"
)

func ListProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.DB.Db.Find(&products)

	return c.Status(fiber.StatusOK).JSON(products)
}

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

func CreateOrder(c *fiber.Ctx) error {
	orderBody := new(models.Order)

	if err := c.BodyParser(orderBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// tx := database.DB.Db.Begin()

	// if tx.Error != nil {
	// 	return tx.Error
	// }

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		tx.Rollback()
	// 	}
	// }()

	// order := models.Order{
	// 	CustomerName: orderBody.CustomerName,
	// 	Items:        []models.ProductItem{},
	// }

	// for _, item := range orderBody.Items {
	// 	var product models.Product

	// 	if err := tx.First(&product, "id = ?", item.ProductID).Error; err != nil {
	// 		tx.Rollback()
	// 		return err
	// 	}

	// 	orderItem := models.ProductItem{
	// 		OrderID:   order.ID,
	// 		ProductID: product.ID,
	// 		Quantity:  item.Quantity,
	// 		UnitPrice: item.UnitPrice,
	// 	}

	// 	if err := tx.Create(&orderItem).Error; err != nil {
	//         // Se ocorrer um erro ao criar o ProductItem, reverta a transação e retorne o erro
	//         tx.Rollback()
	//         return err
	//     }

	// 	order.Items = append(order.Items, orderItem)
	// }

	// if err := tx.Create(&order).Error; err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// tx.Commit()

	database.DB.Db.Create(&orderBody)

	return c.Status(fiber.StatusOK).JSON(orderBody)
}

func ListOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.DB.Db.Preload("ProductItem").Preload("ProductItem.Product").Find(&orders)

	return c.Status(fiber.StatusOK).JSON(orders)
}
