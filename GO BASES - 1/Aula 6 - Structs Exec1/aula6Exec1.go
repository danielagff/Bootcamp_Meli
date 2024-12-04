package main

import "fmt"

type Students struct {
	ID      int
	Name    string
	Surname string
	DNI     string
	Date    string
}

func (a *Students) obterDetalhesAluno() {
	fmt.Printf("ID : %v\nName: %vSurname: %v\nData:%v\nDNI:%v\n ", a.ID, a.Name, a.Surname, a.Date, a.DNI)
}

func main() {
	newStudent := Students {
		ID: 1,
		Name: "Daniel",
		Surname: "Filho",
		DNI: "seilaoqeisso",
		Date: "12/09/2024",
	}

	newStudent.obterDetalhesAluno()
}
