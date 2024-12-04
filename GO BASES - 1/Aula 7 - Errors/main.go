package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Message	string
}

func (c CustomError) Error() string {
	return c.Message
}

func main() {
	hour := 10
	valuePerHour := 50.00
	res, err := getSalaryTaxes(valuePerHour, hour)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(res)
}

func getSalaryTaxes(valuePerHour float64, hours int) (float64, error) {
	salary := (valuePerHour * float64(hours)) * 0.9
	if hours < 80 {
		return 0, errors.New("Error: the worker cannot have worked less than 80 hours per month")
	}
	return	salary, nil
	
}