package usecase

import (
	"context"
	"errors"
	"ride/internal/ride/domain"
	"ride/internal/ride/repository"
	"time"
)


var ErrNoDriverAvailable = errors.New("No driver available")


type RideUsecase interface {
	RequestRide(ctx context.Context, userID string, fare int64) (*domain.Ride, error)
}

type rideUsecase struct {
	rideRepo repository.RideRepository
	driverRepo repository.DriverRepository
}

func NewRideUsecase(
	rideRepo repository.RideRepository,
	driverRepo repository.DriverRepository,
) RideUsecase {
	return &rideUsecase{
		rideRepo: rideRepo,
		driverRepo: driverRepo,
	}
}

func (u *rideUsecase) RequestRide(
	ctx context.Context,
	userID string,
	fare int64,
) (*domain.Ride, error) {

	drivers, err := u.driverRepo.FindAvailable(ctx, 10)
	if err != nil || len(drivers) == 0 {
		return nil, ErrNoDriverAvailable
	}

	driverCh := make(chan domain.Driver, 1)
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	for _, d := range drivers {
		driver := d
		go func() {
			if driver.IsOnline && !driver.IsBusy {
				select {
				case driverCh <- driver:
				case <-ctx.Done():
				}
			}
		}()
	}

	var selected domain.Driver
	select {
	case selected = <-driverCh:
	case <-ctx.Done():
		return nil, ErrNoDriverAvailable
	}

	ride := &domain.Ride{
		UserID:   userID,
		DriverID: selected.ID,
		Status:   domain.RideRequested,
		Fare:     fare,
	}

	if err := u.rideRepo.Create(ctx, ride); err != nil {
		return nil, err
	}

	_ = u.driverRepo.MarkBusy(ctx, selected.ID)

	return ride, nil
}
