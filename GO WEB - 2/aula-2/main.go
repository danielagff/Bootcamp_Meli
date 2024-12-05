package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"golang.org/x/exp/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var ListOfProducts []*Product

func LoadJson() error {
	jsonFile, err := os.ReadFile("products.json")
	if err != nil {
		fmt.Println(err)
		return err
	}

	json.NewDecoder(bytes.NewBuffer(jsonFile)).Decode(&ListOfProducts)

	return nil
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var requestBody Product
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		panic(err)
	}
	if _, err := time.Parse("02/01/2006", requestBody.Expiration); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	requestBody.ID = max() + 1
	requestBody.CodeValue = getProductCodeValue()
	
	ListOfProducts = append(ListOfProducts, &requestBody)

	json.NewEncoder(w).Encode(requestBody)
}

func getProductCodeValue() string {
	existsCode := make(map[string]bool)
	for _, product := range ListOfProducts {
		existsCode[product.CodeValue] = true
	}
	return generateProductCodeValue(existsCode)
}

func generateProductCodeValue(existingCodes map[string]bool) string {
	rand.Seed(uint64(time.Now().UnixNano()))
	const alphanum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for {
		code := ""
		for i := 1; i <= 7; i++ {
			code += string(alphanum[rand.Intn(len(alphanum))])
		}

		if !existingCodes[code] {
			existingCodes[code] = true
			return code
		}
	}
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ListOfProducts)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	product := searchProductById(int(id))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}

func GetHigherPriceProductsByPrice(w http.ResponseWriter, r *http.Request) {
	priceQParam := r.URL.Query().Get("price")
	price, _ := strconv.ParseFloat(priceQParam, 64)
	listOfProducts := searchProductsByHigherPrice(price)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listOfProducts)
}

func searchProductById(id int) *Product {
	for _, product := range ListOfProducts {
		if id == product.ID {
			return product
		}
	}
	return nil
}

func max() int {
	higherIdProduct := ListOfProducts[0]

	for _, product := range ListOfProducts {
		if product.ID > higherIdProduct.ID {
			higherIdProduct = product
		}
	}

	return higherIdProduct.ID
}

func searchProductsByHigherPrice(price float64) []Product {
	var itemsFiltered []Product
	for _, product := range ListOfProducts {
		if price > product.Price {
			itemsFiltered = append(itemsFiltered, *product)
		}
	}
	return itemsFiltered
}

func main() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	LoadJson()

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)

		w.Write([]byte(`pong`))
	})

	router.Get("/products/", GetProducts)

	router.Get("/products/{id}", GetProductById)

	router.Get("/products/search", GetHigherPriceProductsByPrice)

	router.Post("/products", CreateProduct)

	if err := http.ListenAndServe(":8081", router); err != nil {
		panic(err)
	}

}
