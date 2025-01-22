package main

import (
	"go-api/controllers"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection , err := db.ConnectDB()
	if err != nil{
		panic(err)
	}

	//camada repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//camada use case
	ProductUseCase := usecase.NewProductUsecase(ProductRepository)
	//camada controller
	ProductController := controllers.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.Run(":8000")
}