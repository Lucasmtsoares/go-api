package mocks

import (
	"go-api/models"
	"go-api/usecase"
)

// Garantindo que MockProductUsecase implementa a interface corretamente
var _ usecase.IProductUsecase = (*MockProductUsecase)(nil)

type MockProductUsecase struct {
	MockGetProducts    func() ([]models.Product, error)
	MockCreateProduct  func(models.Product) (models.Product, error)
	MockGetProductById func(int) (*models.Product, error)
}

// Implementação do mock
func (m *MockProductUsecase) GetProducts() ([]models.Product, error) {
	return m.MockGetProducts()
}

func (m *MockProductUsecase) CreateProduct(product models.Product) (models.Product, error) {
	return m.MockCreateProduct(product)
}

func (m *MockProductUsecase) GetProductById(id int) (*models.Product, error) {
	return m.MockGetProductById(id)
}
