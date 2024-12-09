package interfaces

import "aula2gobases/internal/model"

type IProductsService interface {
	CreateProduct(product model.Product) (model.Product, error)
	GetAllProducts() ([]model.Product, error)
	GetProductById(id int) (model.Product, error)
	GetHigherPriceProductsByPrice(price float64) ([]model.Product, error)
	DeleteProductById(id int) (string, error)
	UpdateProductById(id int, product model.Product) (string, error)
	UpdatePriceById(id int, product model.Product) (model.Product, error)
	
}

