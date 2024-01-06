package main

import (
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	// Requerimiento 1:
	destination := "Brazil"
	total, _ := tickets.GetTotalTickets(destination)
	fmt.Printf("El total de boletos para %s es: %d", destination, total)

	// Requerimiento 2:

}
