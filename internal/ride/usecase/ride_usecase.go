package usecase

import (
	"context"
	"errors"
	"time"

	tx "ride/internal/common/repository"
	"ride/internal/ride/domain"
	"ride/internal/ride/repository"
)

var ErrNoDriverAvailable = errors.New("no driver available")

type RideUsecase interface {
	RequestRide(ctx context.Context, userID string, fare int64) (*domain.Ride, error)
}

type rideUsecase struct {
	txManager tx.TransactionManager
	rideRepo  repository.RideRepository
	driverRepo repository.DriverRepository
}

func NewRideUsecase(
	txManager tx.TransactionManager,
	rideRepo repository.RideRepository,
	driverRepo repository.DriverRepository,
) RideUsecase {
	return &rideUsecase{
		txManager: txManager,
		rideRepo:  rideRepo,
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

	
	tx, err := u.txManager.Begin(ctx)
	if err != nil {
		return nil, err
	}

	
	defer tx.Rollback()

	
	ride := &domain.Ride{
		UserID:   userID,
		DriverID: selected.ID,
		Status:   domain.RideRequested,
		Fare:     fare,
	}

	if err := u.rideRepo.Create(ctx, tx, ride); err != nil {
		return nil, err
	}

	if err := u.driverRepo.MarkBusy(ctx, tx, selected.ID); err != nil {
		return nil, err
	}


	if err := u.rideRepo.AssignDriver(ctx, tx, ride.ID, selected.ID); err != nil {
		return nil, err
	}


	
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return ride, nil
}
