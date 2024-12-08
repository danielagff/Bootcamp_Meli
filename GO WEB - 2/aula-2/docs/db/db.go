package db

import (
	"aula2gobases/internal/model"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type DataBase struct {
	Products map[int]model.Product
}

func NewDataBase() *DataBase {
	return &DataBase{
		Products: make(map[int]model.Product),
	}
}

func (dataBase *DataBase) LoadJson() (string, error) {
	var products []model.Product
	jsonFile, err := os.ReadFile("/Users/dffilho/Documents/Bootcamp_Meli/GO WEB - 2/aula-2/docs/db/products.json")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	json.NewDecoder(bytes.NewBuffer(jsonFile)).Decode(&products)

	for _, product := range products {
		dataBase.Products[product.ID] = product
	}

	return "bd initialized", nil
}

func (dataBase *DataBase) Commit() error {
	filePath := "/Users/dffilho/Documents/Bootcamp_Meli/GO WEB - 2/aula-2/docs/db/products.json"
	
	productsSlice := make([]model.Product, 0, len(dataBase.Products))
	for _, product := range dataBase.Products {
		productsSlice = append(productsSlice, product)
	}

	sort.Slice(productsSlice, func(i, j int) bool {
		return productsSlice[i].ID < productsSlice[j].ID
	})

	data, err := json.MarshalIndent(productsSlice, "", "  ")
	if err != nil {
		return fmt.Errorf("erro ao converter produtos para JSON: %v", err)
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo para escrita: %v", err)
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		return fmt.Errorf("erro ao escrever no arquivo: %v", err)
	}

	return nil
}