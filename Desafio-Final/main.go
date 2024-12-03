package main

import (
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	file := "/Users/dffilho/Documents/desafio-go-bases/tickets.csv"
	_, err := tickets.LoadAndReadCsv(file)
	if err != nil {
		fmt.Println(err)
	}
	totalTickets, err := tickets.GetTotalTickets("Brazil")
	if err != nil {
		fmt.Println(err)
	}
	startMoarning, moarning, afternoon, night, err:= tickets.GetMornings()
	if err != nil {
		fmt.Println(err)
	}
	average, err := tickets.AverageDestination("Brazil")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Exercicio 1 - %v\n\n",totalTickets)

	fmt.Printf("Exercicio 2 - start Moarning: %v, moarning: %v, afternoon: %v, night: %v\n\n", startMoarning, moarning, afternoon, night)

	fmt.Printf("Exercicio 3 - %v\n\n",average)

}
