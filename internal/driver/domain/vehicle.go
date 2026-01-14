package domain

type VehicleType string

const (
	VehicleBike VehicleType = "BIKE"
	VehicleCar  VehicleType = "CAR"
	VehicleCNG  VehicleType = "CNG"
)

type Vehicle struct {
	ID          string
	DriverID    string

	Type        VehicleType
	PlateNumber string
	Model       string
}
