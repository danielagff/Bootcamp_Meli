package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

func main() {
	name := "teste"
	id := "a"
	phone := "11 987319827"
	address := "Rua Casca 202"

	var sliceError []string

	path := "/Users/dffilho/Documents/Bootcamp_Meli/Aula 8 - Panics/customers.txt"
	// parte desnecessaria apenas para atender o criterio do exercicio
	data, err := os.Open(path)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer data.Close()

	dataRead, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	isDataEmpty, err2 := verifyDataIsEmpty(name, id, phone, address)
	if isDataEmpty {
		sliceError = append(sliceError, fmt.Sprintf("%v", err2))
		fmt.Println(sliceError)
		return
	}

	existsByName := getClientByName(name, dataRead)

	if existsByName {
		defer func() {
			if rec := recover(); rec != nil {
				sliceError = append(sliceError, fmt.Sprintf("%v", rec))
				fmt.Println(sliceError)
			}
		}()

		panic("Error: client already exists")
	}

	res, err3 := addData(path, name, id, phone, address)
	if err3 != nil {
		sliceError = append(sliceError, fmt.Sprintf("%v", err3))
		fmt.Println(sliceError)
		return
	}
	fmt.Println(res)
}

func getClientByName(name string, dataRead []byte) bool {
	return bytes.Contains(dataRead, []byte(name))
}

func verifyDataIsEmpty(name, id, phone, address string) (bool, error) {
	if name == "" || id == "" || phone == "" || address == "" {
		return true, errors.New("Data is null")
	}
	return false, nil
}

func addData(archivePath, name, id, phone, address string) (string, error) {
	formatData := fmt.Sprintf("\n%v | %v | %v | %v", name, id, phone, address)

	file, err := os.OpenFile(archivePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", errors.New("Error: Can't write in this file!")
	}
	defer file.Close()

	_, err2 := file.WriteString(formatData)
	if err2 != nil {
		return "", errors.New("Error: some error writing the data!")
	}

	return "Line included successfully!!!", nil
}