package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaoluizhilario/go-cockroach-poc/internal/models"
	"github.com/joaoluizhilario/go-cockroach-poc/internal/repositories"
)

type OrderAPIHandler struct {
	Repository *repositories.OrderRepository
}

func NewOrderAPIHandler(repository *repositories.OrderRepository) *OrderAPIHandler {
	return &OrderAPIHandler{
		Repository: repository,
	}
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
func (oh *OrderAPIHandler) CreateOrder(c *fiber.Ctx) error {
	orderBody := new(models.Order)

	if err := c.BodyParser(orderBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	orderNew, err := oh.Repository.CreateOrder(orderBody)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(orderNew)
}

// GetOrders ... Get all orders
//
//	@Summary		Get all orders
//	@Description	get all orders hot reload
//	@Tags			Orders
//	@Success		200	{array}		models.Order
//	@Failure		404	{object}	object
//	@Router			/order [get]
func (oh *OrderAPIHandler) ListOrders(c *fiber.Ctx) error {
	orders := oh.Repository.ListOrders()
	return c.Status(fiber.StatusOK).JSON(orders)
}
