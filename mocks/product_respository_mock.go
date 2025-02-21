package mocks

import (
	"go-api/models"

	"github.com/stretchr/testify/mock"
)

//mockProductRepository é um mock da interface ProductRepository
type MockProductRepository struct {
	mock.Mock
}

//Mock para GetProducts
func (m *MockProductRepository) GetProducts() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

//Mock para CreateProduct
func (m *MockProductRepository) CreateProduct(product models.Product) (int, error) {
	args := m.Called(product)
	return args.Int(0), args.Error(1)
}

//Mock para GetProductById
func (m *MockProductRepository) GetProductById(id int) (*models.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}