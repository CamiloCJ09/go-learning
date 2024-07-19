package handler

import (
	"app/internal/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// HandlerTicketDefault represents the default handler of the tickets
type HandlerTicketDefault struct {
	// sv represents the service of the tickets
	sv service.ServiceTicket
}

type HandlerTicket interface {
	GetTotalTickets(w http.ResponseWriter, r *http.Request)
	GetTicketsAmountByDestinationCountry(w http.ResponseWriter, r *http.Request)
	GetPercentageTicketsByDestinationCountry(w http.ResponseWriter, r *http.Request)
}

// NewHandlerTicketDefault creates a new default handler of the tickets
func NewHandlerTicketDefault(sv service.ServiceTicket) HandlerTicket {
	return &HandlerTicketDefault{
		sv: sv,
	}
}

// GetTotalTickets returns the total number of tickets
func (h *HandlerTicketDefault) GetTotalTickets(w http.ResponseWriter, r *http.Request) {
	total, err := h.sv.GetTotalTickets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"total": total})
}

// GetTicketsAmountByDestinationCountry returns the amount of tickets filtered by destination country
func (h *HandlerTicketDefault) GetTicketsAmountByDestinationCountry(w http.ResponseWriter, r *http.Request) {

	destination := chi.URLParam(r, "destination")

	amount, err := h.sv.GetTicketsAmountByDestinationCountry(destination)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"amount": amount})

}

// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
func (h *HandlerTicketDefault) GetPercentageTicketsByDestinationCountry(w http.ResponseWriter, r *http.Request) {

	destination := chi.URLParam(r, "destination")

	percentage, err := h.sv.GetPercentageTicketsByDestinationCountry(destination)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"percentage": percentage})
}
