package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-api/mocks"
	"go-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts_Success(t *testing.T) {
	mockUseCase := &mocks.MockProductUsecase{
		MockGetProducts: func() ([]models.Product, error) {
			return []models.Product{
				{ID: 1, Name: "Produto 1"},
				{ID: 2, Name: "Produto 2"},
			}, nil
		},
	}

	controller := &productController{productUseCase: mockUseCase}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/products", controller.GetProducts)

	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Produto 1")
	assert.Contains(t, w.Body.String(), "Produto 2")
}

func TestGetProducts_Error(t *testing.T) {
	mockUseCase := &mocks.MockProductUsecase{
		MockGetProducts: func() ([]models.Product, error) {
			return nil, errors.New("erro no banco de dados")
		},
	}

	controller := &productController{productUseCase: mockUseCase}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/products", controller.GetProducts)

	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

    fmt.Println("Resposta do servidor:", w.Body.String())

	
    var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Erro ao fazer o unmarshal da resposta: %v", err)
	}

	assert.Contains(t, response["error"], "erro no banco de dados")
}

func TestCreateProduct_Success(t *testing.T) {
	mockUseCase := &mocks.MockProductUsecase{
		MockCreateProduct: func(product models.Product) (models.Product, error) {
			product.ID = 1 // Simula a criação do produto
			return product, nil
		},
	}

	controller := NewProductController(mockUseCase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/products", controller.CreateProduct)

	product := models.Product{Name: "Produto Teste", Price: 100.0}
	jsonValue, _ := json.Marshal(product)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verifica status HTTP 201 Created
	assert.Equal(t, http.StatusCreated, w.Code)

	// Verifica resposta JSON
	var response models.Product
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, product.Name, response.Name)
	assert.Equal(t, product.Price, response.Price)
}

func TestCreateProduct_InvalidJSON(t *testing.T) {
	mockUseCase := &mocks.MockProductUsecase{}

	controller := NewProductController(mockUseCase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/products", controller.CreateProduct)

	invalidJSON := `{"name": "Produto Teste", "price": "abc"}` // 'price' inválido
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer([]byte(invalidJSON)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verifica status HTTP 400 Bad Request
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Verifica que a resposta contém "error"
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "error") // Só verifica se tem um erro, sem fixar a mensagem exata
}

func TestCreateProduct_ErrorCreating(t *testing.T) {
	mockUseCase := &mocks.MockProductUsecase{
		MockCreateProduct: func(product models.Product) (models.Product, error) {
			return models.Product{}, errors.New("erro ao criar produto")
		},
	}

	controller := NewProductController(mockUseCase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/products", controller.CreateProduct)

	product := models.Product{Name: "Produto Teste", Price: 100.0}
	jsonValue, _ := json.Marshal(product)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verifica status HTTP 500 Internal Server Error
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Verifica mensagem de erro
	expectedError := `{"error":"erro ao criar produto"}`
	assert.JSONEq(t, expectedError, w.Body.String())
}
