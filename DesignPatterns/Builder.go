package DesignPatterns

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

//Product
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

//A Builder of type car
type CarBuilder struct {
	v VehicleProduct
}