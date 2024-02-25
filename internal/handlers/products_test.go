package handlers

// import (
// 	"fmt"

// 	"github.com/joaoluizhilario/go-cockroach-poc/internal/models"
// 	"github.com/stretchr/testify/mock"
// )

// type ProductRepositoryMock struct{}

// func (pr *ProductRepositoryMock) ListProducts() *[]models.Product {
// 	products := []models.Product{
// 		models.Product{
// 			ID:    1,
// 			Name:  "Teste",
// 			Price: 123.4,
// 		},
// 	}
// 	return &products
// }

// func (pr *ProductRepositoryMock) CreateProduct(product *models.Product) {
// 	fmt.Println("Created by mock", product.ID, product.Name, product.Price)
// }

// type CtxMock struct {
// 	mock.Mock
// }

// // Mock de método JSON
// func (m *CtxMock) JSON(v interface{}) error {
// 	args := m.Called(v)
// 	return args.Error(0)
// }

// func TestProductAPIHandler_ListProducts(t *testing.T) {

// 	repository := ProductRepositoryMock{}
// 	handler := ProductAPIHandler{
// 		Repository: &repository,
// 	}

// 	ctx := new(CtxMock)
// 	expectedResponse := map[string]interface{}{"id": 1, "name": "Teste", "price": 123.4}
// 	ctx.On("JSON", expectedResponse).Return(nil)

// 	handler.ListProducts(ctx)

// 	ctx.AssertExpectations(t)
// }
