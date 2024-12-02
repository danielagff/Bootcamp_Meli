package main

import (
	product "aula5Struct/structs"
)

func main() {
	var productDaniel product.Product = product.Product{ID: 1, Name: "Cell", Price: 2.22, Description: "Sei la", Category: "Eletronics"}
	productDaniel.Save()
	product.GetAllProducts()
}
