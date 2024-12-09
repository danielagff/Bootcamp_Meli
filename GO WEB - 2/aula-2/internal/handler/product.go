package handler

import (
	"aula2gobases/internal/model"
	"github.com/go-chi/chi"
	serverresponses "aula2gobases/internal/model/serverResponses"
	"aula2gobases/internal/service/interfaces"
	"encoding/json"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	Service interfaces.IProductsService
}

func NewProductHandler(service interfaces.IProductsService) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var requestBody model.Product

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(serverresponses.ServerResponse{
			Message: "json invalido",
		})
		return
	}

	product, err := ph.Service.CreateProduct(requestBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(serverresponses.ServerResponse{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(serverresponses.ServerResponse{
		Data:    product,
		Message: "Sucesso.",
	})
}

func (ph *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := ph.Service.GetAllProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(serverresponses.ServerResponse{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(serverresponses.ServerResponse{
		Data:    products,
		Message: "Sucesso.",
	})
}

func (ph *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(serverresponses.ServerResponse{
			Message: "ID invalido",
		})
		return
	}

	product, err := ph.Service.GetProductById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(serverresponses.ServerResponse{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(serverresponses.ServerResponse{
		Data:    product,
		Message: "Sucesso.",
	})
}

func (ph *ProductHandler) GetHigherPriceProductsByPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	priceStr := r.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(serverresponses.ServerResponse{
			Message: "Parametro invalido.",
		})
		return
	}

	products, err := ph.Service.GetHigherPriceProductsByPrice(price)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(serverresponses.ServerResponse{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(serverresponses.ServerResponse{
		Data:    products,
		Message: "Sucesso.",
	})
}

func (ph *ProductHandler) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(serverresponses.ServerResponse{
			Message: "ID invalido.",
		})
		return
	}

	msg, err := ph.Service.DeleteProductById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(serverresponses.ServerResponse{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(serverresponses.ServerResponse{
		Message: msg,
	})
}

func (ph *ProductHandler) UpdateProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(serverresponses.ServerResponse{
			Message: "ID invalido.",
		})
		return
	}

	var requestBody model.Product
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(serverresponses.ServerResponse{
			Message: "Json invalido",
		})
		return
	}

	msg, err := ph.Service.UpdateProductById(id, requestBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(serverresponses.ServerResponse{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(serverresponses.ServerResponse{
		Message: msg,
	})
}

func (ph *ProductHandler) UpdatePriceById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(serverresponses.ServerResponse{
			Message: "ID invalido.",
		})
		return
	}

	var requestBody model.Product
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(serverresponses.ServerResponse{
			Message: "ID invalido.",
		})
		return
	}

	product, err := ph.Service.UpdatePriceById(id, requestBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(serverresponses.ServerResponse{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(serverresponses.ServerResponse{
		Data:    product,
		Message: "Sucesso.",
	})
}