package service

import "app/pkg/models"

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]models.Vehicle, err error)
	Create(v models.VehicleDoc) (vr models.Vehicle, err error)
	FindByColorAndYear(color string, year int) (v map[int]models.Vehicle, err error)
}
