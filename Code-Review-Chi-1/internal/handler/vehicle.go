package handler

import (
	"app/internal/service"
	"app/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv service.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv service.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]models.VehicleDoc)
		for key, value := range v {
			data[key] = models.VehicleDoc{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) CreateVehicle(w http.ResponseWriter, r *http.Request) {

	var vehicle models.VehicleDoc

	err := json.NewDecoder(r.Body).Decode(&vehicle)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, err.Error())
	}

	data, err := h.sv.Create(vehicle)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, err.Error())
	}

	response.JSON(w, http.StatusCreated, data)

}

func (h *VehicleDefault) GetByColorAndYear(w http.ResponseWriter, r *http.Request) {

	color := chi.URLParam(r, "color")
	year := chi.URLParam(r, "year")
	yearInt, err := strconv.Atoi(year)

	if err != nil {
		response.JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.sv.FindByColorAndYear(color, yearInt)

	if err != nil {
		response.JSON(w, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"message": "success",
		"data":    data,
	})
}
