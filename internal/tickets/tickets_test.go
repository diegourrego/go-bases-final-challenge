package tickets_test

import (
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetTotalTickets(t *testing.T) {
	var mockTickets = tickets.NewTickets(
		tickets.Ticket{ID: 1, Name: "Eduardo Rojas", Email: "eduardo@mail.com",
			FlightTime: "12:20", DestinationCountry: "Finland", Price: 567},
		tickets.Ticket{ID: 2, Name: "Martha Roa", Email: "martha@mail.com",
			FlightTime: "02:00", DestinationCountry: "Brazil", Price: 237},
		tickets.Ticket{ID: 3, Name: "Jorge Español", Email: "jorge@mail.com",
			FlightTime: "21:27", DestinationCountry: "Brazil", Price: 567},
	)

	t.Run("success-case01: should found the correct number of tickets Brazil=2", func(t *testing.T) {
		var expectedNumberOfTickets int = 2
		numberOfTicketsObtained := mockTickets.GetTicketsByDestination("Brazil")
		require.Equal(t, expectedNumberOfTickets, numberOfTicketsObtained)
	})

	t.Run("success-case02: should found zero tickets with unknown destination", func(t *testing.T) {
		var expectedNumberOfTickets int = 0
		numberOfTicketsObtained := mockTickets.GetTicketsByDestination("Mexico")
		require.Equal(t, expectedNumberOfTickets, numberOfTicketsObtained)
	})
}

func TestTickets_GetCountByPeriod(t *testing.T) {
	var mockTickets = tickets.NewTickets(
		tickets.Ticket{ID: 1, Name: "Eduardo Rojas", Email: "eduardo@mail.com",
			FlightTime: "12:20", DestinationCountry: "Finland", Price: 567},
		tickets.Ticket{ID: 2, Name: "Martha Roa", Email: "martha@mail.com",
			FlightTime: "02:00", DestinationCountry: "Brazil", Price: 237},
		tickets.Ticket{ID: 3, Name: "Jorge Español", Email: "jorge@mail.com",
			FlightTime: "21:27", DestinationCountry: "Brazil", Price: 567},
	)

	t.Run("success - case 01: should find correct number of tickets by early morning", func(t *testing.T) {
		period := "early morning"
		ticketsEarlyMorningExpected := 1
		ticketsEarlyMorningObtained, err := mockTickets.GetCountByPeriod(period)
		require.Equal(t, ticketsEarlyMorningExpected, ticketsEarlyMorningObtained)
		require.Nil(t, err)
	})

	t.Run("success - case 02: should find correct number of tickets by morning", func(t *testing.T) {
		period := "morning"
		ticketsEarlyMorningExpected := 1
		ticketsEarlyMorningObtained, err := mockTickets.GetCountByPeriod(period)
		require.Equal(t, ticketsEarlyMorningExpected, ticketsEarlyMorningObtained)
		require.Nil(t, err)
	})

	t.Run("success - case 03: should find correct number of tickets by afternoon", func(t *testing.T) {
		period := "afternoon"
		ticketsEarlyMorningExpected := 0
		ticketsEarlyMorningObtained, err := mockTickets.GetCountByPeriod(period)
		require.Equal(t, ticketsEarlyMorningExpected, ticketsEarlyMorningObtained)
		require.Nil(t, err)
	})

	t.Run("success - case 04: should find correct number of tickets by early morning", func(t *testing.T) {
		period := "night"
		ticketsEarlyMorningExpected := 1
		ticketsEarlyMorningObtained, err := mockTickets.GetCountByPeriod(period)
		require.Equal(t, ticketsEarlyMorningExpected, ticketsEarlyMorningObtained)
		require.Nil(t, err)
	})

	t.Run("failure - case01: should returns an error with an invalid period", func(t *testing.T) {
		period := "my test period"
		ticketsEarlyMorningExpected := 0
		ticketsEarlyMorningObtained, err := mockTickets.GetCountByPeriod(period)
		require.Equal(t, ticketsEarlyMorningExpected, ticketsEarlyMorningObtained)
		require.EqualError(t, err, "invalid time")

	})
}

func TestAverageDestination(t *testing.T) {
	var mockTickets = tickets.NewTickets(
		tickets.Ticket{ID: 1, Name: "Eduardo Rojas", Email: "eduardo@mail.com",
			FlightTime: "12:20", DestinationCountry: "Finland", Price: 567},
		tickets.Ticket{ID: 2, Name: "Martha Roa", Email: "martha@mail.com",
			FlightTime: "02:00", DestinationCountry: "Brazil", Price: 237},
		tickets.Ticket{ID: 3, Name: "Jorge Español", Email: "jorge@mail.com",
			FlightTime: "21:27", DestinationCountry: "Brazil", Price: 567},
	)

	var mockEmptyTickets = tickets.NewTickets()

	t.Run("success - case01: should finds the correct average", func(t *testing.T) {
		expectedAverage := 0.6666666666666666
		averageObtained, err := mockTickets.GetAverage("Brazil")
		require.Equal(t, expectedAverage, averageObtained)
		require.Nil(t, err)
	})

	t.Run("success - case02: should returns zero if it not find tickets with the specified destination", func(t *testing.T) {
		expectedAverage := 0.0
		averageObtained, err := mockTickets.GetAverage("Mexico")
		require.Equal(t, expectedAverage, averageObtained)
		require.Nil(t, err)
	})

	t.Run("failure - case01: should returns an error if len of tickets is zero", func(t *testing.T) {
		expectedAverage := 0.0
		averageObtained, err := mockEmptyTickets.GetAverage("Mexico")
		require.Equal(t, expectedAverage, averageObtained)
		require.Errorf(t, err, "len of tickets equal zero. can not divide it")
	})
}
