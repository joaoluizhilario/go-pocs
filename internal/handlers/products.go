package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaoluizhilario/go-cockroach-poc/internal/models"
	"github.com/joaoluizhilario/go-cockroach-poc/internal/repositories"
)

type ProductAPIHandler struct {
	Repository repositories.ProductRepositoryInterface
}

// GetProducts ... Get all products
//
//	@Summary		Get all products
//	@Description	get all products
//	@Tags			Products
//	@Success		200	{array}		models.Product
//	@Failure		404	{object}	object
//	@Router			/product [get]
func (ah *ProductAPIHandler) ListProducts(c *fiber.Ctx) error {
	products := ah.Repository.ListProducts()
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
func (ah *ProductAPIHandler) CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	ah.Repository.CreateProduct(product)

	return c.Status(fiber.StatusOK).JSON(product)
}
