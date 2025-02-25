package repository

import "app/pkg/models"

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]models.Vehicle, err error)

	// Create a new vehicle
	Create(v models.VehicleDoc) (models.Vehicle, error)

	//Find By color And Year
	FindByColorAndYear(color string, year int) (v map[int]models.Vehicle, err error)
}
