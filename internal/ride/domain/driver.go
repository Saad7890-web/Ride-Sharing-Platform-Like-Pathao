package domain

type Driver struct {
	ID        string
	IsOnline  bool
	IsBusy    bool
	Latitude  float64
	Longitude float64
}
