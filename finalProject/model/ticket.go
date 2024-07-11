package model

var Tickets []Ticket

// Ticket struct

// Ticket represents a flight ticket.
type Ticket struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Destination string  `json:"destination"`
	Country     string  `json:"country"`
	FlightTime  string  `json:"flightTime"`
	Price       float64 `json:"price"`
}

func (t *Ticket) CreateNewTicket() {
	Tickets = append(Tickets, *t)
}

func (t *Ticket) UpdateTicket(id int) {
	for i, ticket := range Tickets {
		if ticket.Id == id {
			Tickets[i] = *t
		}
	}
}

func (t *Ticket) DeleteTicket(id int) {
	for i, ticket := range Tickets {
		if ticket.Id == id {
			Tickets = append(Tickets[:i], Tickets[i+1:]...)
		}
	}
}

func GetTicketById(id int) Ticket {
	for _, ticket := range Tickets {
		if ticket.Id == id {
			return ticket
		}
	}
	return Ticket{}
}

func GetAllTickets() []Ticket {
	return Tickets
}
