package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joaoluizhilario/go-cockroach-poc/internal/models"
	"github.com/stretchr/testify/assert"
)

type OrderRepositoryMock struct{}

func (pr *OrderRepositoryMock) ListOrders() *[]models.Order {
	orders := []models.Order{
		{
			ID:           1,
			CustomerName: "Teste",
			TotalPrice:   123.4,
			Items: []models.ProductItem{
				{
					ProductID: 1,
					UnitPrice: 20.40,
					Quantity:  2,
				},
			},
		},
	}
	return &orders
}

func (or *OrderRepositoryMock) CreateOrder(order *models.Order) (*models.Order, error) {
	fmt.Println(order.CustomerName, "Order Created by mock")
	return &models.Order{
		ID:           1,
		CustomerName: "Customer Teste",
		TotalPrice:   123.4,
		Items: []models.ProductItem{
			{
				ProductID: 1,
				UnitPrice: 20.40,
				Quantity:  2,
			},
		},
	}, nil
}

func TestOrderAPIHandler_ListOrders(t *testing.T) {

	expectedResponse := []models.Order{
		{
			ID:           1,
			CustomerName: "Customer Teste",
			TotalPrice:   123.4,
			Items: []models.ProductItem{
				{
					ProductID: 1,
					UnitPrice: 20.40,
					Quantity:  2,
				},
			},
		},
	}

	repository := OrderRepositoryMock{}
	handler := OrderAPIHandler{
		Repository: &repository,
	}

	app := fiber.New()
	app.Get("/order", handler.ListOrders)

	resp, err := app.Test(httptest.NewRequest("GET", "/order", nil))
	if err != nil {
		t.Errorf("Erro ao realizar requisicao")
	}

	var orders []models.Order
	if err := json.NewDecoder(resp.Body).Decode(&orders); err != nil {
		t.Errorf("Erro ao decodificar JSON: %v", err)
	}

	for i, order := range expectedResponse {
		expectedId := expectedResponse[i].ID
		expectedCustomerName := expectedResponse[i].CustomerName
		expectedTotalPrice := expectedResponse[i].TotalPrice

		assert.Equalf(t, order.ID, expectedId, "ID assetion")
		assert.Equalf(t, order.CustomerName, expectedCustomerName, "CustomerName assetion")
		assert.Equalf(t, order.TotalPrice, expectedTotalPrice, "ID assetion")

		for j, item := range order.Items {
			expctedProductId := expectedResponse[i].Items[j].ProductID
			expectedUnitPrice := expectedResponse[i].Items[j].UnitPrice
			expectedQuantity := expectedResponse[i].Items[j].Quantity

			assert.Equalf(t, item.ProductID, expctedProductId, "Order ProductItem ProductID assetion")
			assert.Equalf(t, item.UnitPrice, expectedUnitPrice, "Order ProductItem UnitPrice assetion")
			assert.Equalf(t, item.Quantity, expectedQuantity, "Order ProductItem  Quantity assetion")
		}
	}
}

func TestOrderAPIHandler_CreateOrder(t *testing.T) {

	repository := OrderRepositoryMock{}
	handler := OrderAPIHandler{
		Repository: &repository,
	}

	app := fiber.New()
	app.Post("/order", handler.CreateOrder)

	bodyString := `{"customer_name": "Customer Teste", "items":[{"product_id": 1, "quantity": 3, "price": 123.4}, {"product_id": 2, "quantity": 3, "price": 13.4}, {"product_id": 3, "quantity": 3, "price": 23.4}]}`
	var expectedResponse map[string]interface{}
	if err := json.Unmarshal([]byte(bodyString), &expectedResponse); err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}

	req := httptest.NewRequest("POST", "/order", strings.NewReader(bodyString))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.FailNow()
	}

	var order map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&order); err != nil {
		t.FailNow()
	}

	assert.Equalf(t, resp.StatusCode, http.StatusOK, "Status code 200 OK")

	assert.NotNil(t, order["id"])
	assert.Equalf(t, order["customer_name"], expectedResponse["customer_name"], "CustomerName assertion")
	// assert.Equalf(t, order["total_price"], expectedResponse["price"], "TotalPrice assertion")

	var total float32 = 0.0

	if sliceItems, ok := order["items"].([]interface{}); ok {
		for _, item := range sliceItems {
			if mapItem, ok := item.(map[string]interface{}); ok {
				if floatValue, ok := mapItem["price"].(float32); ok {
					total += floatValue
				}
			}
		}
	}

	fmt.Println("Valor do total:", total)
}
