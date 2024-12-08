package interfaces

import "aula2gobases/internal/model"

type IProductsRepository interface {
	Save(product model.Product) (model.Product, error)
	GetAll() ([]model.Product, error)
	GetById(id int) (model.Product, error)
	DeleteById(id int) (string, error)
	UpdateById(id int, product model.Product)
	UpdatePriceById(id int, product model.Product)
}

