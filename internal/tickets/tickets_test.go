package tickets_test

import (
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGetTotalTickets(t *testing.T) {

	currentDir, err := os.Getwd()
	require.Nil(t, err)

	err = os.Chdir("/Users/durrego/Boot_Practice/learningHUB/desafio-go-bases")
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
