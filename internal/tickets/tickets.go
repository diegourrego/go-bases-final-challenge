package tickets

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID                 int
	Name               string
	Email              string
	DestinationCountry string
	FlightTime         string
	Price              int
}

type Tickets struct {
	Tickets []Ticket
}

func NewTickets() *Tickets {
	return &Tickets{}
}

func (t *Tickets) AddTicket(ticket Ticket) {
	t.Tickets = append(t.Tickets, ticket)
}

func (t *Tickets) GetTickets() []Ticket {
	return t.Tickets
}

func (t *Tickets) GetTicket(id int) Ticket {
	for _, ticket := range t.Tickets {
		if ticket.ID == id {
			return ticket
		}
	}
	return Ticket{}
}

func (t *Tickets) LoadFile() {
	file, err := os.Open("./tickets.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	reader := csv.NewReader(file)
	list, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range list {
		ticket := Ticket{}
		ticket.ID = int(item[0][0])
		ticket.Name = item[1]
		ticket.Email = item[2]
		ticket.DestinationCountry = item[3]
		ticket.FlightTime = item[4]
		ticket.Price = int(item[5][0])
		t.AddTicket(ticket)
	}

}

// GetTotalTickets ejemplo 1
func GetTotalTickets(destination string) (int, error) {

	if destination == "" {
		return 0, errors.New("the destination can't be an empty string")
	}

	tickets := Tickets{}
	tickets.LoadFile() // Aquí cargo la información dentro del slice
	var totalTickets []Ticket

	for _, ticket := range tickets.GetTickets() {
		if ticket.DestinationCountry == destination {
			totalTickets = append(totalTickets, ticket)
		}
	}

	return len(totalTickets), nil
}

// GetNumberOfTicketsByPeriod GetMornings ejemplo 2
func GetNumberOfTicketsByPeriod(period string) (int, error) {
	tickets := Tickets{}
	tickets.LoadFile()

	periodFlights := make(map[string]int)

	for _, ticket := range tickets.GetTickets() {
		hourArray := strings.Split(ticket.FlightTime, ":")
		hourInt, err := strconv.Atoi(hourArray[0])
		if err != nil {
			return 0, err
		}

		switch {
		case hourInt >= 0 && hourInt <= 6:
			periodFlights["early morning"]++
		case hourInt >= 7 && hourInt <= 12:
			periodFlights["morning"]++
		case hourInt >= 13 && hourInt <= 19:
			periodFlights["afternoon"]++
		case hourInt >= 20 && hourInt <= 23:
			periodFlights["night"]++
		}
	}

	count, found := periodFlights[period]
	if !found {
		return 0, errors.New("invalid period")
	}
	return count, nil

}

//// AverageDestination ejemplo 3
//func AverageDestination(destination string) (float64, error) {
//	totalTickets, err := GetFileData()
//	if err != nil {
//		fmt.Println("ERROR in GetFileData", err)
//		return 0, err
//	}
//
//	ticketsByDestination, err := GetTotalTickets(destination)
//	if err != nil {
//		fmt.Printf("Error en GetTotalTickets")
//		return 0, err
//	}
//	total := float64(len(totalTickets))
//	return float64(ticketsByDestination) / total, nil
//
//}
