package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Printf("\n**************** EXERCICIO 1 **********************\n")
	var userSalary float32 = 51000.10
	fmt.Printf("\nYour taxe is: $%v\n ", getSalaryTaxes(float32(userSalary)))

	fmt.Printf("\n**************** EXERCICIO 2 **********************\n")
	result := getStudentAvarage(6.00, 3.00, 9.00, 4.00)
	fmt.Printf("\n%.2f\n", result)

	fmt.Printf("\n**************** EXERCICIO 3 **********************\n")
	getMonthlySalaryMinutes(120, "B")

	fmt.Printf("\n**************** EXERCICIO 4 **********************\n")
	oper := operation("max")
	firstReturn, secondReturn := oper(1, 4, 4, 4)
	fmt.Printf("\n%s -  %v\n", secondReturn, firstReturn)

	fmt.Printf("\n**************** EXERCICIO 5 **********************\n")

	userOptAnimal := "ham"
	userAnimalQuantity := 4
	callFuncAninal, err := getFoodQuantityByAnimal(userOptAnimal)

	if err != nil {
		panic(err)
	}

	hamFoodTotal := callFuncAninal(userAnimalQuantity)

	var gramOrKg string
	if hamFoodTotal >= 1 {
		gramOrKg = fmt.Sprintf("%gKg", hamFoodTotal)
	} else {
		gramOrKg = fmt.Sprintf("%gg", hamFoodTotal)
	}

	fmt.Printf("\nYou need %v to feed %v %v\n", gramOrKg, userAnimalQuantity, userOptAnimal)

}

func getSalaryTaxes(salary float32) (taxes float32) {
	taxes = 0.00
	if salary > 50000.00 {
		taxes = salary * (17.00 / 100.00)
	} else if salary > 150000.00 {
		taxes = salary * ((17.00 + 10) / 100.00)
	}
	return taxes
}
func getStudentAvarage(studentsTests ...float64) (studentsAvarage float64) {
	studentsAvarage = 0.00
	validTests := 0.00
	for _, studentTest := range studentsTests {
		if studentTest > 0 {
			studentsAvarage = studentsAvarage + studentTest
			validTests++
		}

	}
	studentsAvarage = studentsAvarage / validTests
	return studentsAvarage
}

func getMonthlySalaryMinutes(minutes int64, category string) {
	const (
		categoryA = 1000.00
		categoryB = 1500.00
		categoryC = 3000.00
	)
	if category == "A" || category == "a" {
		salary := (minutes / 60) * categoryA
		fmt.Printf("\nSalario A - $ %v", salary)
	} else if category == "B" || category == "b" {
		salary := (float32(minutes) / 60.0) * categoryB
		salary = salary + (salary * (20.0 / 100.0))
		fmt.Printf("Salario B - $%v", salary)
	} else if category == "C" || category == "c" {
		salary := (float32(minutes) / 60) * categoryC
		salary = salary + (salary * (50.0 / 100.0))
		fmt.Printf("Salario C - $%v", salary)
	} else {
		fmt.Println("Nenhuma categoria encontrada!")
	}

}
func operation(option string) func(list ...int) (int, string) {
	switch option {
	case "min":
		return getMin
	case "avarage":
		return getAvarage
	case "max":
		return getMax

	}
	return nil
}

func getMin(listNum ...int) (int, string) {
	if len(listNum) == 0 {
		return 0, "Empty List"
	}
	n := listNum[0]
	for _, num := range listNum {
		if num < n {
			n = num
		}
	}
	return n, "Min Number"
}

func getAvarage(listNum ...int) (int, string) {
	if len(listNum) == 0 {
		return 0, "Empty List"
	}
	sum := 0
	for _, num := range listNum {
		sum += num
	}
	sum = sum / len(listNum)
	return sum, "Avarage Number"
}

func getMax(listNum ...int) (int, string) {

	if len(listNum) == 0 {
		return 0, "Empty List"
	}
	n := listNum[0]
	for _, num := range listNum {
		if num > n {
			n = num
		}
	}
	return n, "Max Number"
}

func getFoodQuantityByAnimal(animalOpt string) (func(quantityAnimal int) float64, error) {
	switch animalOpt {
	case "dog":
		return getFoodForTheDog, nil
	case "cat":
		return getFoodForTheCat, nil
	case "ham":
		return getFoodForTheHamster, nil
	case "spi":
		return getFoodForTheSpider, nil
	default:
		return nil, errors.New("This animal don't exists!")
	}

}

func getFoodForTheDog(quantityAnimal int) float64 {
	return float64(quantityAnimal) * 10.00
}

func getFoodForTheCat(quantityAnimal int) float64 {
	return float64(quantityAnimal) * 5.00
}

func getFoodForTheHamster(quantityAnimal int) float64 {
	return float64(quantityAnimal) * 0.25
}

func getFoodForTheSpider(quantityAnimal int) float64 {
	return float64(quantityAnimal) * 0.15
}
