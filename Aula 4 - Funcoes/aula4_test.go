package main

import (
	"bytes"
	"os"
	"testing"
)

func TestGetSalaryTaxes(t *testing.T) {
	res := getSalaryTaxes(51000.10)
	expected := 8670.018

	if res != float32(expected) {
		t.Errorf("Esperado= %v; obteve %v", expected, res)
	}

}

func TestGetStudentAvarage(t *testing.T) {
	res := getStudentAvarage(6.00, 3.00, 9.00, 4.00)
	expected := 5.50

	if res != expected {
		t.Errorf("Esperado= %v; obteve %v", expected, res)
	}

}

func TestGetMonthlySalaryMinutes(t *testing.T) {

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	getMonthlySalaryMinutes(120, "B")

	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stdout = old

	output := buf.String()

	expected := "Salario B - $3600"
	if output != expected {
		t.Errorf("Esperado %q mas obteve %q", expected, output)
	}

}

func TestOperation(t *testing.T) {
	oper := operation("max")
	firstReturn, secondReturn := oper(1, 4, 4, 4)
	isExpected := false

	if firstReturn == 4 && secondReturn == "Max Number" {
		isExpected = true
	}
	if !isExpected {
		t.Errorf("Esperado= 4; obteve %v\nEsperado=Max Number; obteve %v", firstReturn, secondReturn)
	}
}

func TestGetFoodQuantityByAnimal(t *testing.T) {
	getFoodForTheHamster, err := getFoodQuantityByAnimal("ham")
	if err != nil {
		t.Fatalf("Esperava nenhum erro ao obter a função closure, mas obteve: %v", err)
	}

	res := getFoodForTheHamster(4)
	expected := float64(1)
	if res != expected {
		t.Errorf("Esperado= %v; obteve %v", expected, res)
	}
}
