package repository

import (
	"database/sql"
	"fmt"
	"go-api/models"

)

type ProductRepository interface {
	GetProducts() ([]models.Product, error)
	CreateProduct(product models.Product) (int, error)
	GetProductById(id int) (*models.Product, error)
}

type ProductRepositoryImpl struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{connection: connection}
}

func (pr *ProductRepositoryImpl) GetProducts() ([]models.Product, error) {
	query := "SELECT id, product_name, price FROM products"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.Product{}, err
	}

	var productList []models.Product 
	var productobj models.Product

	for rows.Next() {
		err = rows.Scan(
			&productobj.ID,
			&productobj.Name,
			&productobj.Price)
		if err != nil {
			fmt.Println(err)
			return []models.Product{}, err
		}
		productList = append(productList, productobj)
	}
	rows.Close()

	return productList, nil
}

func (pr *ProductRepositoryImpl) CreateProduct(product models.Product) (int, error){
	var id int
	query, err := pr.connection.Prepare("INSERT INTO products"+
	"(product_name, price)"+
	"VALUES ($1, $2) RETURNING id")
	if err != nil{
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil
}

func (pr *ProductRepositoryImpl) GetProductById(id int) (*models.Product, error){
	query, err := pr.connection.Prepare("SELECT * FROM products WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var product models.Product

	err = query.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)
	if err != nil{
		if err == sql.ErrNoRows{
			return nil, nil
		}
		return nil, err
	}
	query.Close()
	return &product, nil
}