package tickets_test

import (
	"strings"
	"testing"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func TestGetTotalTickets(t *testing.T) {
	tickets.TicketList = []tickets.Ticket{
		{Id: 1, Name: "t", Email: "teste@rmail.cowm", Destination: "Brazil", Hour: "10:05", Price: 785},
	}

	tests := []struct {
		Name        string
		Input       string
		ExpectedOut string
	}{
		{Name: "Teste de Sucesso", Input: "Brazil", ExpectedOut: "1 are going to the same destination - Brazil"},
		{Name: "Teste de Erro", Input: "", ExpectedOut: "not valid country name"},
	}

	for _, itemTest := range tests {
		t.Run(itemTest.Name, func(t *testing.T) {
			result, err := tickets.GetTotalTickets(itemTest.Input)
			if err != nil {
				if err.Error() != itemTest.ExpectedOut {
					t.Errorf("Esperava: %v, Recebeu: %v", itemTest.ExpectedOut, err.Error())
				}
			} else {
				if result != itemTest.ExpectedOut {
					t.Errorf("Esperava: %v, Recebeu: %v", itemTest.ExpectedOut, result)
				}
			}
		})
	}
}

func TestGetMorningsSucess(t *testing.T) {
	tickets.TicketList = []tickets.Ticket{
		{Id: 1, Name: "t", Email: "teste@rmail.com", Destination: "Brazil", Hour: "00:00", Price: 785},
	}

	tests := []struct {
		Name         string
		Input        string
		ExpectedOut1 int
		ExpectedOut2 int
		ExpectedOut3 int
		ExpectedOut4 int
		ErrorMsg     string
	}{
		{Name: "Teste de Sucesso", ExpectedOut1: 0, ExpectedOut2: 1, ExpectedOut3: 0, ExpectedOut4: 0, ErrorMsg: ""},
	}
	for _, itemTest := range tests {
		t.Run(itemTest.Name, func(t *testing.T) {
			res1, res2, res3, res4, _ := tickets.GetMornings()
			if res1 != itemTest.ExpectedOut1 && res2 != itemTest.ExpectedOut2 &&
				res3 != itemTest.ExpectedOut3 && res4 != itemTest.ExpectedOut4 {
				t.Errorf("Esperava: %v,%v,%v,%v Recebeu: %v,%v,%v,%v", itemTest.ExpectedOut1,
					itemTest.ExpectedOut2, itemTest.ExpectedOut3, itemTest.ExpectedOut4,
					res1, res2, res3, res4)
			}

		})
	}
}

func TestGetMorningsError(t *testing.T) {
	tickets.TicketList = []tickets.Ticket{
		{Id: 1, Name: "t", Email: "teste@rmail.com", Destination: "Brazil", Hour: "x", Price: 785},
	}

	tests := []struct {
		Name         string
		ExpectedOut1 int
		ExpectedOut2 int
		ExpectedOut3 int
		ExpectedOut4 int
		ErrorMsg     string
	}{
		{Name: "Teste de Erro", ExpectedOut1: 0, ExpectedOut2: 1, ExpectedOut3: 0, ExpectedOut4: 0, ErrorMsg: "Error 1: this - x, is not a number"},
	}

	for _, itemTest := range tests {
		t.Run(itemTest.Name, func(t *testing.T) {
			_, _, _, _, err := tickets.GetMornings()

			if err == nil {
				t.Errorf("Esperava erro, recebeu nil")
				return
			}

			if !strings.Contains(err.Error(), itemTest.ErrorMsg) {
				t.Errorf("Esperava: %v, Recebeu: %v", itemTest.ErrorMsg, err.Error())
			}
		})
	}
}

func TestAverageDestination(t *testing.T) {
	tickets.TicketList = []tickets.Ticket{
		{Id: 1, Name: "t", Email: "teste@rmail.com", Destination: "Brazil", Hour: "00:00", Price: 785},
	}

	tests := []struct {
		Name        string
		Input 		string
		ExpectedOut string
	}{
		{Name: "Teste de Sucesso", Input: "Brazil", ExpectedOut: "Country: Brazil, Percent: 100.00%"},
		{Name: "Teste de Erro", Input: "", ExpectedOut: "not valid Contry name"},
	}

	for _, itemTest := range tests {
		t.Run(itemTest.Name, func(t *testing.T) {
			result, err := tickets.GetTotalTickets(itemTest.Input)
			if err != nil {
				if strings.Contains(err.Error(),itemTest.ExpectedOut) {
					t.Errorf("Esperava: %v, Recebeu: %v", itemTest.ExpectedOut, err.Error())
				}
			} else {
				if strings.Contains(result,itemTest.ExpectedOut) {
					t.Errorf("Esperava: %v, Recebeu: %v", itemTest.ExpectedOut, result)
				}
			}
		})
	}
}
