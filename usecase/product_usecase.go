package usecase

import (
	"go-api/models"
	"go-api/repository"
)

type ProductUsecase struct{
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase{
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]models.Product, error){
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product models.Product) (models.Product, error){
	productID, err := pu.repository.CreateProduct(product)
	if err != nil{
		return models.Product{}, err
	}
	product.ID = productID

	return product, nil
}

func (pu *ProductUsecase) GetProductById(id int) (*models.Product, error){
	product, err := pu.repository.GetProductById(id)
	if err != nil{
		return nil, err
	}
	return product, nil
}