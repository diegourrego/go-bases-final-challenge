package main

import (
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	// Requerimiento 1:
	destination := "Brazil"
	total, _ := tickets.GetTotalTickets(destination)
	fmt.Printf("El total de boletos para %s es: %d\n", destination, total)

	//// Requerimiento 2:
	//period := "early morning"
	//numTicketsEarlyMorning, err := tickets.GetNumberOfTicketsByPeriod(period)
	//if err != nil {
	//	fmt.Println("ERROR GetNumberOfTicketsByPeriod", err)
	//}
	//fmt.Printf("El número de tiquetes en %s es: %d\n", period, numTicketsEarlyMorning)
	//
	//period = "morning"
	//numTicketsMorning, err := tickets.GetNumberOfTicketsByPeriod(period)
	//if err != nil {
	//	fmt.Println("ERROR GetNumberOfTicketsByPeriod", err)
	//}
	//fmt.Printf("El número de tiquetes en %s es: %d\n", period, numTicketsMorning)
	//
	//period = "afternoon"
	//numTicketsAfternoon, err := tickets.GetNumberOfTicketsByPeriod(period)
	//if err != nil {
	//	fmt.Println("ERROR GetNumberOfTicketsByPeriod", err)
	//}
	//fmt.Printf("El número de tiquetes en %s es: %d\n", period, numTicketsAfternoon)
	//
	//period = "night"
	//numTicketsNight, err := tickets.GetNumberOfTicketsByPeriod(period)
	//if err != nil {
	//	fmt.Println("ERROR GetNumberOfTicketsByPeriod", err)
	//}
	//fmt.Printf("El número de tiquetes en %s es: %d\n", period, numTicketsNight)
	//
	//// Requerimiento 3:
	//aDestination := "Finland"
	//average, err := tickets.AverageDestination(aDestination)
	//if err != nil {
	//	fmt.Println("error AverageDestination", err)
	//	return
	//}
	//fmt.Printf("The average for %s is: %f", aDestination, average)
}
