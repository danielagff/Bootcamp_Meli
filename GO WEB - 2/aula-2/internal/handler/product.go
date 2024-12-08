package handler

import (
	"aula2gobases/internal/model"
	"aula2gobases/internal/service"
	"encoding/json"
	"net/http"
)

type ProductHandler struct {
	Service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var requestBody model.Product

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		panic(err)
	}
	product, err := ph.Service.CreateProduct(requestBody)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(201)

	json.NewEncoder(w).Encode(product)
}
