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
	ID                 int
	Name               string
	Email              string
	DestinationCountry string
	FlightTime         string
	Price              int
}

type Tickets []Ticket

func (ts Tickets) GetTickets() Tickets {
	return ts
}

func (ts Tickets) GetTicketsByDestination(destinationCountry string) int {
	var ticketsRequested = Tickets{}
	for _, ticket := range ts {
		if ticket.DestinationCountry == destinationCountry {
			ticketsRequested = append(ticketsRequested, ticket)
		}
	}
	return len(ticketsRequested)
}

func (ts Tickets) GetCountByPeriod(time string) (int, error) {

	var earlyMorningSlice []int
	var morningSlice []int
	var afternoonSlice []int
	var nightSlice []int

	for _, ticket := range ts {
		hourArray := strings.Split(ticket.FlightTime, ":")
		hourInt, err := strconv.Atoi(hourArray[0])
		if err != nil {
			fmt.Println("Error en GetCountByPeriod", err)
			return 0, err
		}

		switch {
		case hourInt >= 0 && hourInt <= 6:
			earlyMorningSlice = append(earlyMorningSlice, hourInt)
		case hourInt >= 7 && hourInt <= 12:
			morningSlice = append(morningSlice, hourInt)
		case hourInt >= 13 && hourInt <= 19:
			afternoonSlice = append(afternoonSlice, hourInt)
		case hourInt >= 20 && hourInt <= 23:
			nightSlice = append(nightSlice, hourInt)
		}
	}

	switch time {
	case "early morning":
		return len(earlyMorningSlice), nil
	case "morning":
		return len(morningSlice), nil
	case "afternoon":
		return len(afternoonSlice), nil
	case "night":
		return len(nightSlice), nil
	default:
		return 0, errors.New("invalid time")
	}

}

func (ts Tickets) GetAverage(destination string) (float64, error) {
	ticketsObtained := ts.GetTicketsByDestination(destination)
	totalOfTickets := float64(len(ts))
	if totalOfTickets == 0 {
		return 0.0, errors.New("len of tickets equal zero. can not divide it")
	}
	return float64(ticketsObtained) / totalOfTickets, nil
}

func NewTickets(tickets ...Ticket) Tickets {
	var ts Tickets

	for _, ticket := range tickets {
		ts = append(ts, ticket)
	}

	return ts
}

func GetFileData() (Tickets, error) {
	file, err := os.Open("./tickets.csv")
	if err != nil {
		return nil, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
			return
		}
	}()

	reader := csv.NewReader(file)
	list, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	tickets := NewTickets()

	// Recorrer el listado
	for _, ticketData := range list {
		// Conversiones necesarias para crear Ticket
		id, err := strconv.Atoi(ticketData[0])
		if err != nil {
			return nil, err
		}

		price, err := strconv.Atoi(ticketData[5])
		if err != nil {
			return nil, err
		}

		tickets = append(tickets, Ticket{
			ID:                 id,
			Name:               ticketData[1],
			Email:              ticketData[2],
			DestinationCountry: ticketData[3],
			FlightTime:         ticketData[4],
			Price:              price,
		})

	}
	return tickets, nil
}

// GetTotalTickets ejemplo 1
func GetTotalTickets(destination string) (int, error) {
	tickets, err := GetFileData()
	if err != nil {
		fmt.Println("ERROR in GetFileData", err)
		return 0, err
	}
	ticketsFounded := tickets.GetTicketsByDestination(destination)
	return ticketsFounded, nil
}

// GetNumberOfTicketsByPeriod GetMornings ejemplo 2
func GetNumberOfTicketsByPeriod(time string) (int, error) {
	tickets, err := GetFileData()
	if err != nil {
		fmt.Println("ERROR in GetFileData", err)
		return 0, err
	}
	numberTickets, err := tickets.GetCountByPeriod(time)
	if err != nil {
		fmt.Println("ERROR en GetMornings", err)
		return 0, err
	}
	return numberTickets, nil
}

// AverageDestination ejemplo 3
func AverageDestination(destination string) (float64, error) {
	totalTickets, err := GetFileData()
	if err != nil {
		fmt.Println("ERROR in GetFileData", err)
		return 0, err
	}

	ticketsByDestination, err := GetTotalTickets(destination)
	if err != nil {
		fmt.Printf("Error en GetTotalTickets")
		return 0, err
	}
	total := float64(len(totalTickets))
	return float64(ticketsByDestination) / total, nil

}
