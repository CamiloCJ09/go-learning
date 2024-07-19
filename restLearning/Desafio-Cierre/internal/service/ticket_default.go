package service

import (
	"app/internal/repository"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp repository.RepositoryTicket
}

type ServiceTicket interface {
	// GetTotalAmountTickets returns the total amount of tickets
	GetTotalTickets() (total int, err error)

	// GetTicketsAmountByDestinationCountry returns the amount of tickets filtered by destination country
	GetTicketsAmountByDestinationCountry(country string) (amount int, err error)

	// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
	GetPercentageTicketsByDestinationCountry(country string) (percentage float64, err error)
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp repository.RepositoryTicket) ServiceTicket {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalTickets() (total int, err error) {
	tickets, err := s.rp.Get()
	return len(tickets), err
}

// GetTicketsAmountByDestinationCountry returns the amount of tickets filtered by destination country
func (s *ServiceTicketDefault) GetTicketsAmountByDestinationCountry(country string) (amount int, err error) {
	tickets, err := s.rp.GetTicketsByDestinationCountry(country)
	return len(tickets), err
}

// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
func (s *ServiceTicketDefault) GetPercentageTicketsByDestinationCountry(country string) (percentage float64, err error) {
	tickets, err := s.rp.Get()
	if err != nil {
		return
	}

	ticketsByCountry, err := s.rp.GetTicketsByDestinationCountry(country)
	if err != nil {
		return
	}

	percentage = float64(len(ticketsByCountry)) / float64(len(tickets)) * 100
	return
}
