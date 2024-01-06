package tickets

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Ticket struct {
	ID                 int
	Name               string
	Email              string
	DestinationCountry string
	FlightTime         time.Time
	Price              int
}

type Tickets []Ticket

func (ts Tickets) GetTickets() Tickets {
	return ts
}

func (ts Tickets) GetTicketsByDestination(destinationCountry string) Tickets {
	var ticketsRequested = Tickets{}
	for _, ticket := range ts {
		if ticket.DestinationCountry == destinationCountry {
			ticketsRequested = append(ticketsRequested, ticket)
		}
	}
	return ticketsRequested
}

//func (ts Tickets) GetCountByPeriod(time string) (int, error) {
//
//}

func NewTickets(tickets ...Ticket) Tickets {
	var ts Tickets

	for _, ticket := range tickets {
		ts = append(ts, ticket)
	}

	return ts
}

func (ts Tickets) GetTotal() int {
	return len(ts)
}

func GetFileData() (Tickets, error) {
	file, err := os.Open("./tickets.csv")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return Tickets{}, err
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
		fmt.Println("Error Reading the file:", err)
		return Tickets{}, err
	}

	tickets := NewTickets()

	// Recorrer el listado
	for _, ticketData := range list {
		// Conversiones necesarias para crear Ticket
		id, err := strconv.Atoi(ticketData[0])
		if err != nil {
			fmt.Println("Error convirtiendo id a entero", err)
			return Tickets{}, err
		}

		hour, err := time.Parse("15:04", ticketData[4])
		if err != nil {
			fmt.Println("Error convirtiendo hour a time", err)
			return Tickets{}, err
		}

		price, err := strconv.Atoi(ticketData[5])
		if err != nil {
			fmt.Println("Error convirtiendo price a float", err)
			return Tickets{}, err
		}

		tickets = append(tickets, Ticket{
			ID:                 id,
			Name:               ticketData[1],
			Email:              ticketData[2],
			DestinationCountry: ticketData[3],
			FlightTime:         hour,
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
	return len(ticketsFounded), nil
}

// GetMornings ejemplo 2
func GetMornings(time string) (int, error) {
	return 0, nil
}

// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {
	return 0, nil
}
