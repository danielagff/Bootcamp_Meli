package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	Id          int
	Name        string
	Email       string
	Destination string
	Hour        string
	Price       float64
}

var TicketList []Ticket

func (t Ticket) addItem() {
	TicketList = append(TicketList, t)
}

func LoadAndReadCsv(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		csvItem, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Erro ao ler o Csv: ", err)
			return "", fmt.Errorf("error: cannot read this CSV - %v", err)
		}

		id, _ := strconv.Atoi(csvItem[0])

		price, _ := strconv.ParseFloat(csvItem[5], 64)

		ticket := Ticket{
			Id:          id,
			Name:        csvItem[1],
			Email:       csvItem[2],
			Destination: csvItem[3],
			Hour:        csvItem[4],
			Price:       price,
		}

		ticket.addItem()

	}

	return "Data load successfull!", nil
}

func GetTotalTickets(destination string) (string, error) {
	if destination == "" {
		return "", fmt.Errorf("not valid country name")
	}
	totalPerson := 0
	msg := ""
	for _, item := range TicketList {
		if item.Destination == destination {
			totalPerson++
		}
	}

	msg = fmt.Sprint(totalPerson, " are going to the same destination - ", destination)
	return msg, nil
}

func GetMornings() (int, int, int, int, error) {
	var startMoarning, moarning, afternoon, night int
	var errorsSlice []error
	for _, item := range TicketList {
		itemHour := includePatternInHour(item.Hour)
		if itemHour >= "00:00" && itemHour <= "06:59" {
			startMoarning++
		} else if itemHour >= "07:00" && itemHour <= "12:59" {
			moarning++
		} else if itemHour >= "13:00" && itemHour <= "19:59" {
			afternoon++
		} else if itemHour >= "20:00" && itemHour <= "23:59" {
			night++
		} else {
			errorsSlice = append(errorsSlice, fmt.Errorf("this - %v, is not a number", item.Hour))
		}
	}
	if len(errorsSlice) > 0 {
		msgError := ""
		for i, item := range errorsSlice {
			msgError += fmt.Sprintf("Error %v: %v\n", i+1, item)
		}
		return 0, 0, 0, 0, errors.New(msgError)
	}

	return startMoarning, moarning, afternoon, night, nil
}

func includePatternInHour(hour string) string {
	part := strings.Split(hour, ":")
	if len(part) != 2 {
		return hour
	}

	if len(part[0]) == 1 {
		part[0] = "0" + part[0]
	}
	if len(part[1]) == 1 {
		part[1] = "0" + part[1]
	}

	return part[0] + ":" + part[1]
}
func AverageDestination(country string) (string, error) {
	count := make(map[string]int)
	if country == "" {
		return "", fmt.Errorf("not valid Contry name %s", country)
	}


	totalTickets := 0

	for _, item := range TicketList {
		if item.Destination != "" {
			if item.Destination == country {
				count[item.Destination]++
			}
			totalTickets++
		}
	}

	percent := (float64(count[country]) / float64(totalTickets)) * 100
	msg := fmt.Sprintf("Country: %s, Percent: %.2f%%\n", country, percent)

	return msg, nil
}
