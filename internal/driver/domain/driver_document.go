package domain

import "time"

type DocumentType string

const (
	DocumentNID        DocumentType = "NID"
	DocumentLicense    DocumentType = "LICENSE"
	DocumentVehicleReg DocumentType = "VEHICLE_REGISTRATION"
)

type DriverDocument struct {
	ID           string
	DriverID     string

	Type         DocumentType
	DocumentURL  string
	Verified     bool

	CreatedAt    time.Time
}
