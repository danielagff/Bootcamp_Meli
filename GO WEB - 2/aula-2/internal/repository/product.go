package repository

import (
	"aula2gobases/docs/db"
	"aula2gobases/internal/model"
	"fmt"
)

type ProductRepository struct {
	DataBase *db.DataBase
}

func NewProductRepository(db *db.DataBase) *ProductRepository {
	return &ProductRepository{DataBase: db}
}

func (r *ProductRepository) Save(product model.Product) (model.Product, error) {
	id := r.getMaxId() + 1
	product.ID = id
	r.DataBase.Products[id] = product
	err := r.DataBase.Commit()
	if err != nil {
		return model.Product{}, err
	}
	return r.DataBase.Products[id], nil
}

func (r *ProductRepository) GetAll() ([]model.Product, error) {
	var listOfProducts []model.Product
	for _, product := range r.DataBase.Products {
		listOfProducts = append(listOfProducts, product)
	}
	return listOfProducts, nil
}

func (r *ProductRepository) GetById(id int) (model.Product, error) {
	product, exists := r.DataBase.Products[id]
	if !exists {
		return model.Product{}, fmt.Errorf("produto com o id: %d não encontrado", id)
	}
	return product, nil
}

func (r *ProductRepository) DeleteById(id int) (string, error) {
	if _, exists := r.DataBase.Products[id]; !exists {
		return "", fmt.Errorf("produto com o id: %d não encontrado", id)
	}
	delete(r.DataBase.Products, id)
	return "Product deleted!", nil
}

func (r *ProductRepository) UpdateById(id int, product model.Product) (string, error) {
	if _, exists := r.DataBase.Products[id]; !exists {
		return "", fmt.Errorf("produto com o id: %d não encontrado", id)
	}
	r.DataBase.Products[id] = product
	return "Product Updated!", nil
}

func (r *ProductRepository) UpdatePriceById(id int, product model.Product) (model.Product, error) {
	if _, exists := r.DataBase.Products[id]; !exists {
		return product, fmt.Errorf("produto com o id %d não encontrado", id)
	}
	productInDB := r.DataBase.Products[id]
	productInDB.Price = product.Price
	r.DataBase.Products[id] = productInDB
	return r.DataBase.Products[id], nil
}

func (r *ProductRepository) getMaxId() int {
	maxID := 0
	for _, product := range r.DataBase.Products {
		if product.ID > maxID {
			maxID = product.ID
		}
	}
	return maxID
}
