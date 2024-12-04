package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\n**************** EXERCICIO 1 **********************")
	// Exercicio 1
	var word string = "Sei la"
	fmt.Println("\nTamanho : ", len(word))
	for _, char := range word {

		fmt.Printf("%q\n", char)
	}
	// Exercicio 2
	fmt.Printf("\n**************** EXERCICIO 2 **********************")

	const (
		minAge         uint8   = 22
		isWorking      bool    = true
		minYearWorking uint8   = 1
		salary         float64 = 100000.00
	)
	var (
		userAge         uint8   = 20
		userIsWorking   bool    = false
		userYearWorking uint8   = 0
		userSalary      float64 = 110000.01
	)

	var (
		ageValidation          bool = userAge < minAge
		isWorkingValidation    bool = userIsWorking != isWorking
		yearsWorkingValidation bool = userYearWorking < minYearWorking
		userSalaryValidation   bool = userSalary < salary
	)

	var errorList []string

	if ageValidation {
		errorList = append(errorList, "você não atende aos requisitos minimos de idade!")
	}
	if isWorkingValidation {
		errorList = append(errorList, "você não está trabalhando!")
	}
	if yearsWorkingValidation {
		errorList = append(errorList, "você não tem o tempo minimo trabalhando!")
	}

	if len(errorList) > 0 {
		fmt.Println("\nEmprestimo não concedido pois:")
		for _, err := range errorList {
			fmt.Println(err)
		}

	} else {
		if userSalaryValidation {
			fmt.Println("Empréstimo concedido com juros!")
		} else {
			fmt.Println("Empréstimo concedido sem juros!")
		}
	}
	fmt.Printf("\n**************** EXERCICIO 3 **********************")

	var monthNumber uint8 = 1

	monthMap := map[uint8]string{
		1:  "Janeiro",
		2:  "Fevereiro",
		3:  "Março",
		4:  "Abril",
		5:  "Maio",
		6:  "Junho",
		7:  "Julho",
		8:  "Agosto",
		9:  "Setembro",
		10: "Outubro",
		11: "Novembro",
		12: "Dezembro",
	}

	if monthNumber > 12 || monthNumber < 1 {
		fmt.Printf("\nEsse mês não existe\n")
	} else {
		fmt.Printf("\n %s ", monthMap[monthNumber])
	}
	fmt.Printf("\n**************** EXERCICIO 4 **********************\n")
	
	var wishedEmployee string  = "Benjamin"
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var plus21yearsOld int = 0

	delete(employees,"Pedro")
	employees["Federico"] = 25

	fmt.Printf("%s tem %d anos",wishedEmployee,employees[wishedEmployee])

	for _, value := range employees {
		if(value > 21) {
			plus21yearsOld++
		}
	}
	fmt.Printf("\n%d funcionarios tem mais que 21 anos", plus21yearsOld)



}
