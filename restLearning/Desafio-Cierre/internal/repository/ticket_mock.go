package repository

import (
	"app/internal/domain"
)

// NewRepositoryTicketMock creates a new repository for tickets in a map
func NewRepositoryTicketMock() *RepositoryTicketMock {
	return &RepositoryTicketMock{}
}

// RepositoryTicketMock implements the repository interface for tickets
type RepositoryTicketMock struct {
	// FuncGet represents the mock for the Get function
	FuncGet func() (t map[int]domain.TicketAttributes, err error)
	// FuncGetTicketsByDestinationCountry
	FuncGetTicketsByDestinationCountry func(country string) (t map[int]domain.TicketAttributes, err error)

	// Spy verifies if the methods were called
	Spy struct {
		// Get represents the spy for the Get function
		Get int
		// GetTicketsByDestinationCountry represents the spy for the GetTicketsByDestinationCountry function
		GetTicketsByDestinationCountry int
	}
}

// GetAll returns all the tickets
func (r *RepositoryTicketMock) Get() (t map[int]domain.TicketAttributes, err error) {
	// spy
	r.Spy.Get++

	// mock
	t, err = r.FuncGet()
	return
}

// GetTicketsByDestinationCountry returns the tickets filtered by destination country
func (r *RepositoryTicketMock) GetTicketsByDestinationCountry(country string) (t map[int]domain.TicketAttributes, err error) {
	// spy
	r.Spy.GetTicketsByDestinationCountry++

	// mock
	t, err = r.FuncGetTicketsByDestinationCountry(country)
	return
}

// SetTickets sets the tickets
func (r *RepositoryTicketMock) SetTickets(tickets map[int]domain.TicketAttributes) {
	// mock
}
