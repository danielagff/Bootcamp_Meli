package structs

import (
	"fmt"
)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var listoOfProducts []Product

func (p Product) Save() {
	listoOfProducts = append(listoOfProducts, p)
	fmt.Println("Produto adicionado!")
}

func GetAllProducts() {
	fmt.Println(listoOfProducts)
}
