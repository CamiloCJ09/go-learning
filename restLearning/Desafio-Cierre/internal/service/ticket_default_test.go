package service_test

import (
	"app/internal/domain"
	"app/internal/repository"
	"app/internal/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for ServiceTicketDefault.GetTotalAmountTickets
func TestServiceTicketDefault_GetTotalAmountTickets(t *testing.T) {
	t.Run("success to get total tickets", func(t *testing.T) {
		// arrange
		// - repository: mock
		rp := repository.NewRepositoryTicketMock()
		// - repository: set-up
		rp.FuncGet = func() (t map[int]domain.TicketAttributes, err error) {
			t = map[int]domain.TicketAttributes{
				1: {
					Name:    "John",
					Email:   "johndoe@gmail.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
			}
			return
		}

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		total, err := sv.GetTotalTickets()

		// assert
		expectedTotal := 1
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})

	t.Run("error to get total tickets", func(t *testing.T) {
		// arrange
		// - repository: mock
		rp := repository.NewRepositoryTicketMock()
		// - repository: set-up
		rp.FuncGet = func() (t map[int]domain.TicketAttributes, err error) {
			err = errors.New("error")
			return
		}

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		total, err := sv.GetTotalTickets()

		// assert
		require.Error(t, err)
		require.Zero(t, total)
	})
}

// Tests for ServiceTicketDefault.GetTicketsAmountByDestinationCountry
func TestServiceTicketDefault_GetTicketsAmountByDestinationCountry(t *testing.T) {
	t.Run("success to get tickets by destination country", func(t *testing.T) {
		// arrange
		// - repository: mock
		rp := repository.NewRepositoryTicketMock()
		// - repository: set-up
		rp.FuncGetTicketsByDestinationCountry = func(country string) (t map[int]domain.TicketAttributes, err error) {
			t = map[int]domain.TicketAttributes{
				1: {
					Name:    "John",
					Email:   "email@email.com:",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
			}
			return
		}

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act

		amount, err := sv.GetTicketsAmountByDestinationCountry("USA")

		// assert

		expectedAmount := 1
		require.NoError(t, err)
		require.Equal(t, expectedAmount, amount)
	})

	t.Run("error to get tickets by destination country", func(t *testing.T) {
		// arrange
		// - repository: mock
		rp := repository.NewRepositoryTicketMock()
		// - repository: set-up
		rp.FuncGetTicketsByDestinationCountry = func(country string) (t map[int]domain.TicketAttributes, err error) {
			err = errors.New("error")
			return
		}

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		amount, err := sv.GetTicketsAmountByDestinationCountry("USA")

		// assert
		require.Error(t, err)
		require.Zero(t, amount)

	})
}

// Tests for ServiceTicketDefault.GetPercentageTicketsByDestinationCountry
func TestServiceTicketDefault_GetPercentageTicketsByDestinationCountry(t *testing.T) {
	t.Run("success to get percentage of tickets by destination country", func(t *testing.T) {
		// arrange
		// - repository: mock
		rp := repository.NewRepositoryTicketMock()
		// - repository: set-up
		rp.FuncGet = func() (t map[int]domain.TicketAttributes, err error) {
			t = map[int]domain.TicketAttributes{
				1: {
					Name:    "John",
					Email:   "email@email.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
				2: {
					Name:    "Jane",
					Email:   "jane@email.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},

				3: {
					Name:    "Alice",
					Email:   "alice@email.com",
					Country: "Serbia",
					Hour:    "10:00",
					Price:   100,
				},
			}
			return
		}

		rp.FuncGetTicketsByDestinationCountry = func(country string) (t map[int]domain.TicketAttributes, err error) {
			t = map[int]domain.TicketAttributes{
				1: {
					Name:    "John",
					Email:   "email@email.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
				2: {
					Name:    "Jane",
					Email:   "jane@email.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
			}
			return
		}

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		percentage, err := sv.GetPercentageTicketsByDestinationCountry("USA")

		// assert
		expectedPercentage := 66.66666666666666
		require.NoError(t, err)
		require.Equal(t, expectedPercentage, percentage)
	})

}
