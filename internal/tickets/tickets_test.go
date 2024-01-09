package tickets_test

import (
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

func TestGetTotalTickets(t *testing.T) {

	currentDir, err := os.Getwd()
	require.Nil(t, err)

	index := strings.Index(currentDir, "desafio-go-bases")
	if index == -1 {
		// La subcadena no se encontró, manejar este caso según tus necesidades
		fmt.Println("Subcadena no encontrada en la ruta.")
		return
	}
	err = os.Chdir(currentDir[:index+len("desafio-go-bases")])
	require.Nil(t, err)

	defer func() {
		err := os.Chdir(currentDir)
		require.Nil(t, err)
	}()

	t.Run("success-case01: should found the correct number of tickets Brazil = 45", func(t *testing.T) {
		numberOfTicketsObtained, err := tickets.GetTotalTickets("Brazil")
		var expectedNumberOfTickets int = 45
		require.Equal(t, expectedNumberOfTickets, numberOfTicketsObtained)
		require.Nil(t, err)
	})

	t.Run("success-case02: should found zero number of tickets for Antarctic", func(t *testing.T) {
		numberOfTicketsObtained, err := tickets.GetTotalTickets("Antarctic")
		var expectedNumberOfTickets int = 0
		require.Equal(t, expectedNumberOfTickets, numberOfTicketsObtained)
		require.Nil(t, err)
	})

	t.Run("failure - case01: should returns an error when send an empty string", func(t *testing.T) {
		numberOfTicketsObtained, err := tickets.GetTotalTickets("")
		var expectedNumberOfTickets int = 0
		require.Equal(t, expectedNumberOfTickets, numberOfTicketsObtained)
		require.EqualError(t, err, "the destination can't be an empty string")
	})
}

func TestTickets_GetCountByPeriod(t *testing.T) {
	currentDir, err := os.Getwd()
	require.Nil(t, err)

	index := strings.Index(currentDir, "desafio-go-bases")
	if index == -1 {
		// La subcadena no se encontró, manejar este caso según tus necesidades
		fmt.Println("Subcadena no encontrada en la ruta.")
		return
	}
	err = os.Chdir(currentDir[:index+len("desafio-go-bases")])
	require.Nil(t, err)

	defer func() {
		err := os.Chdir(currentDir)
		require.Nil(t, err)
	}()

	t.Run("success - case 01: should find correct number of tickets by early morning", func(t *testing.T) {
		period := "early morning"
		ticketsEarlyMorningExpected := 304
		ticketsEarlyMorningObtained, err := tickets.GetNumberOfTicketsByPeriod(period)
		require.Equal(t, ticketsEarlyMorningExpected, ticketsEarlyMorningObtained)
		require.Nil(t, err)
	})

	t.Run("success - case 02: should find correct number of tickets by morning", func(t *testing.T) {
		period := "morning"
		ticketsEarlyMorningExpected := 256
		ticketsEarlyMorningObtained, err := tickets.GetNumberOfTicketsByPeriod(period)
		require.Equal(t, ticketsEarlyMorningExpected, ticketsEarlyMorningObtained)
		require.Nil(t, err)
	})

	t.Run("success - case 03: should find correct number of tickets by afternoon", func(t *testing.T) {
		period := "afternoon"
		ticketsEarlyMorningExpected := 289
		ticketsEarlyMorningObtained, err := tickets.GetNumberOfTicketsByPeriod(period)
		require.Equal(t, ticketsEarlyMorningExpected, ticketsEarlyMorningObtained)
		require.Nil(t, err)
	})

	t.Run("success - case 04: should find correct number of tickets by early morning", func(t *testing.T) {
		period := "night"
		ticketsEarlyMorningExpected := 151
		ticketsEarlyMorningObtained, err := tickets.GetNumberOfTicketsByPeriod(period)
		require.Equal(t, ticketsEarlyMorningExpected, ticketsEarlyMorningObtained)
		require.Nil(t, err)
	})

	t.Run("failure - case01: should returns an error with an invalid period", func(t *testing.T) {
		period := "my test period"
		ticketsEarlyMorningExpected := 0
		ticketsEarlyMorningObtained, err := tickets.GetNumberOfTicketsByPeriod(period)
		require.Equal(t, ticketsEarlyMorningExpected, ticketsEarlyMorningObtained)
		require.EqualError(t, err, "invalid period")

	})
}

func TestAverageDestination(t *testing.T) {
	currentDir, err := os.Getwd()
	require.Nil(t, err)

	index := strings.Index(currentDir, "desafio-go-bases")
	if index == -1 {
		// La subcadena no se encontró, manejar este caso según tus necesidades
		fmt.Println("Subcadena no encontrada en la ruta.")
		return
	}
	err = os.Chdir(currentDir[:index+len("desafio-go-bases")])
	require.Nil(t, err)

	defer func() {
		err := os.Chdir(currentDir)
		require.Nil(t, err)
	}()

	t.Run("success - case01: should finds the correct average", func(t *testing.T) {
		averageObtained, err := tickets.AverageDestination("Brazil")
		expectedAverage := 0.045
		require.Equal(t, expectedAverage, averageObtained)
		require.Nil(t, err)
	})

	t.Run("success - case02: should returns zero if it not find tickets with the specified destination", func(t *testing.T) {
		expectedAverage := 0.0
		averageObtained, err := tickets.AverageDestination("Antarctic")
		require.Equal(t, expectedAverage, averageObtained)
		require.Nil(t, err)
	})

}
