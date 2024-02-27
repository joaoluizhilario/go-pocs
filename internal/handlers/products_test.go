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

type ProductRepositoryMock struct{}

func (pr *ProductRepositoryMock) ListProducts() *[]models.Product {
	products := []models.Product{
		{
			ID:    1,
			Name:  "Teste",
			Price: 123.4,
		},
	}
	return &products
}

func (pr *ProductRepositoryMock) CreateProduct(product *models.Product) {
	fmt.Println(product.Name, "Created by mock")
}

func TestProductAPIHandler_ListProducts(t *testing.T) {

	expectedResponse := []models.Product{
		{
			ID:    1,
			Name:  "Teste",
			Price: 123.4,
		},
	}

	repository := ProductRepositoryMock{}
	handler := ProductAPIHandler{
		Repository: &repository,
	}

	app := fiber.New()
	app.Get("/product", handler.ListProducts)

	resp, err := app.Test(httptest.NewRequest("GET", "/product", nil))
	if err != nil {
		t.Errorf("Erro ao realizar requisicao")
	}

	var products []models.Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		t.Errorf("Erro ao decodificar JSON: %v", err)
	}

	for i, product := range expectedResponse {
		expectedId := expectedResponse[i].ID
		expectedName := expectedResponse[i].Name
		expectedPrice := expectedResponse[i].Price

		assert.Equalf(t, product.ID, expectedId, "ID assetion")
		assert.Equalf(t, product.Name, expectedName, "Name assetion")
		assert.Equalf(t, product.Price, expectedPrice, "ID assetion")
	}
}

func TestProductAPIHandler_CreateProduct(t *testing.T) {

	repository := ProductRepositoryMock{}
	handler := ProductAPIHandler{
		Repository: &repository,
	}

	app := fiber.New()
	app.Post("/product", handler.CreateProduct)

	bodyString := `{"id": 1, "name": "Teste", "price": 123.4}`
	var expectedResponse map[string]interface{}
	if err := json.Unmarshal([]byte(bodyString), &expectedResponse); err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}

	req := httptest.NewRequest("POST", "/product", strings.NewReader(bodyString))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.FailNow()
	}

	var product map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		t.FailNow()
	}

	assert.Equalf(t, resp.StatusCode, http.StatusOK, "Status code 200 OK")

	assert.Equalf(t, product["id"], expectedResponse["id"], "ID assertion")
	assert.Equalf(t, product["name"], expectedResponse["name"], "Name assertion")
	assert.Equalf(t, product["price"], expectedResponse["price"], "Price assertion")
}
