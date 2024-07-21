package repository

import (
	"app/pkg/models"
	"errors"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]models.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]models.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]models.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// Create a new vehicle
func (r *VehicleMap) Create(v models.VehicleDoc) (int, error) {

	if validateIdDontExists(v.ID, r.db) {
		return 0, errors.New("id already exists")
	}

	vehicle := models.CreateVehicle(v)

	r.db[v.ID] = vehicle

	return vehicle.Id, nil
}

func validateIdDontExists(id int, db map[int]models.Vehicle) bool {
	_, ok := db[id]
	return ok
}

func (r *VehicleMap) FindByColorAndYear(color string, year int) (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)
	for _, vehicle := range r.db {
		if vehicle.Color == color && vehicle.FabricationYear == year {
			v[vehicle.Id] = vehicle
		}
	}
	if len(v) == 0 {
		err = errors.New("no se encontraron vehículos con esos criterios")
	}
	return
}
