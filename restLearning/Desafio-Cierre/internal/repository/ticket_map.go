package repository

import (
	"app/internal/domain"
	"log"
)

// NewRepositoryTicketMap creates a new repository for tickets in a map
func NewRepositoryTicketMap(dbFile string, lastId int) RepositoryTicket {
	return &RepositoryTicketMap{
		dbFile: dbFile,
		lastId: lastId,
	}
}

// RepositoryTicket represents the repository interface for tickets
type RepositoryTicket interface {
	// GetAll returns all the tickets
	Get() (t map[int]domain.TicketAttributes, err error)

	// GetTicketByDestinationCountry returns the tickets filtered by destination country
	GetTicketsByDestinationCountry(country string) (t map[int]domain.TicketAttributes, err error)

	// SetTickets sets the tickets
	SetTickets(tickets map[int]domain.TicketAttributes)
}

// RepositoryTicketMap implements the repository interface for tickets in a map
type RepositoryTicketMap struct {
	// db represents the database in a map
	// - key: id of the ticket
	// - value: ticket
	db map[int]domain.TicketAttributes

	// dbFile represents the file where the database is stored
	dbFile string

	// lastId represents the last id of the ticket
	lastId int
}

// GetAll returns all the tickets
func (r *RepositoryTicketMap) Get() (t map[int]domain.TicketAttributes, err error) {
	// create a copy of the map
	t = make(map[int]domain.TicketAttributes, len(r.db))
	for k, v := range r.db {
		t[k] = v
	}

	return t, nil
}

func (r *RepositoryTicketMap) SetTickets(tickets map[int]domain.TicketAttributes) {
	r.db = tickets
}

// GetTicketsByDestinationCountry returns the tickets filtered by destination country
func (r *RepositoryTicketMap) GetTicketsByDestinationCountry(country string) (t map[int]domain.TicketAttributes, err error) {
	// create a copy of the map
	t = make(map[int]domain.TicketAttributes)
	for k, v := range r.db {
		if v.Country == country {
			log.Println(v)
			t[k] = v
		}
	}

	return t, nil
}
