package usecase_test

import (
	"errors"
	"fmt"
	"go-api/mocks"
	"go-api/models"
	"go-api/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Teste para GetProducts - sucesso
func TestGetProducts_Success(t *testing.T) {
	mockRepo := new(mocks.MockProductRepository)
    mockProducts := []models.Product{
        {ID: 1, Name: "Product A", Price: 10.0},
        {ID: 2, Name: "Product B", Price: 20.0},
    }

    // Certifique-se de que o mock está configurado corretamente
    mockRepo.On("GetProducts").Return(mockProducts, nil)

    usecase := usecase.NewProductUsecase(mockRepo)
    products, err := usecase.GetProducts()

    // Verifique se não há erro
    assert.NoError(t, err)
    // Verifique se dois produtos foram retornados
    assert.Len(t, products, 2)
    // Verifique se o nome do primeiro produto é "Product A"
    assert.Equal(t, "Product A", products[0].Name)
    mockRepo.AssertExpectations(t)

}

//Teste para GetProducts (erro no repositório)
func TestGetProducts_Error(t *testing.T) {
	mockRepo := new(mocks.MockProductRepository)
	// Simulando um erro no repositório
	mockRepo.On("GetProducts").Return([]models.Product{}, fmt.Errorf("erro ao obter produtos"))

	usecase := usecase.NewProductUsecase(mockRepo)
	products, err := usecase.GetProducts()

	// Verificando se um erro é retornado
	assert.Error(t, err)
	// Verificando se a lista de produtos é nil ou vazia
	
	assert.Empty(t, products) 
	// Verificando se a mensagem de erro está correta
	assert.Equal(t, "erro ao obter produtos", err.Error())
	mockRepo.AssertExpectations(t)
	
}

// Teste para CreateProduct (sucesso)
func TestCreateProduct_Success(t *testing.T) {
	mockRepo := new(mocks.MockProductRepository)
	mockProduct := models.Product{Name: "Novo Produto", Price: 50.0}
	mockRepo.On("CreateProduct", mockProduct).Return(1, nil)

	usecase := usecase.NewProductUsecase(mockRepo)
	createdProduct, err := usecase.CreateProduct(mockProduct)

	assert.NoError(t, err)
	assert.Equal(t, 1, createdProduct.ID)
	mockRepo.AssertExpectations(t)
}

// Teste para CreateProduct (erro ao criar)
func TestCreateProduct_Error(t *testing.T) {
	mockRepo := new(mocks.MockProductRepository)
	mockProduct := models.Product{Name: "Novo Produto", Price: 50.0}
	mockRepo.On("CreateProduct", mockProduct).Return(0, errors.New("falha ao criar produto"))

	usecase := usecase.NewProductUsecase(mockRepo)
	createdProduct, err := usecase.CreateProduct(mockProduct)

	assert.Error(t, err)
	assert.Equal(t, 0, createdProduct.ID)
	assert.Equal(t, "falha ao criar produto", err.Error())
	mockRepo.AssertExpectations(t)
}

// Teste para GetProductById (produto encontrado)
func TestGetProductById_Success(t *testing.T) {
	mockRepo := new(mocks.MockProductRepository)
	mockProduct := models.Product{ID: 1, Name: "Produto Teste", Price: 100.0}

	mockRepo.On("GetProductById", 1).Return(&mockProduct, nil)

	usecase := usecase.NewProductUsecase(mockRepo)
	product, err := usecase.GetProductById(1)

	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "Produto Teste", product.Name)
	mockRepo.AssertExpectations(t)
}

// Teste para GetProductById (produto não encontrado)
func TestGetProductById_NotFound(t *testing.T) {
	mockRepo := new(mocks.MockProductRepository)

	mockRepo.On("GetProductById", 2).Return(nil, nil)

	usecase := usecase.NewProductUsecase(mockRepo)
	product, err := usecase.GetProductById(2)

	assert.NoError(t, err)
	assert.Nil(t, product)
	mockRepo.AssertExpectations(t)
}

// Teste para GetProductById (erro no banco)
func TestGetProductById_Error(t *testing.T) {
	mockRepo := new(mocks.MockProductRepository)

	mockRepo.On("GetProductById", 3).Return(nil, errors.New("erro no banco"))

	usecase := usecase.NewProductUsecase(mockRepo)
	product, err := usecase.GetProductById(3)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "erro no banco", err.Error())
	mockRepo.AssertExpectations(t)
}