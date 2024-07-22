package service

import (
	"app/internal/repository"
	"app/pkg/models"
	"errors"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp repository.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp repository.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]models.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) Create(v models.VehicleDoc) (vr models.Vehicle, err error) {

	vr, err = s.rp.Create(v)

	if err != nil {
		return models.Vehicle{}, errors.New("user not created")
	}

	return
}

func (s *VehicleDefault) FindByColorAndYear(color string, year int) (v map[int]models.Vehicle, err error) {
	v, err = s.rp.FindByColorAndYear(color, year)
	return
}
