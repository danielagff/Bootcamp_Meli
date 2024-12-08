package main

import (
	"aula2gobases/docs/db"
	"aula2gobases/internal/handler"
	"aula2gobases/internal/repository"
	"aula2gobases/internal/service"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)


func main() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	dataBase := db.NewDataBase()
	dataBase.LoadJson()
	productRepository := repository.NewProductRepository(dataBase)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	//router.Get("/products/", GetProducts)

	//router.Get("/products/{id}", GetProductById)

	//router.Get("/products/search", GetHigherPriceProductsByPrice)

	router.Post("/products", productHandler.CreateProduct)

	if err := http.ListenAndServe(":8081", router); err != nil {
		panic(err)
	}

}
