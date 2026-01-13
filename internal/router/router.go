package router

import (
	"net/http"

	"ride/internal/auth/handler"
	"ride/internal/auth/middleware"
	"ride/internal/auth/repository"
	"ride/internal/auth/service"
	"ride/internal/auth/usecase"

	rideHandler "ride/internal/ride/handler"
	rideRepoPg "ride/internal/ride/repository"
	rideUC "ride/internal/ride/usecase"

	"ride/internal/database"
	"ride/internal/health"

	"github.com/jackc/pgx/v5/pgxpool"
)

	




func New(db *pgxpool.Pool, jwtSvc service.JWTService) http.Handler {
	mux := http.NewServeMux()

	userRepo := repository.NewPostgresUserRepository(db)

	authUC := usecase.NewAuthUseCase(userRepo, jwtSvc)

	authHandler := handler.NewAuthHandler(authUC)

	

	mux.HandleFunc("/health", health.Handler)
	mux.HandleFunc("/auth/register", authHandler.Register)
	mux.HandleFunc("/auth/login", authHandler.Login)

	
	txManager := database.NewPostgresTxManager(db)


	rideRepo := rideRepoPg.NewPostgresRideRepository(db)
	driverRepo := rideRepoPg.NewPostgresDriverRepository(db)

	rideUC := rideUC.NewRideUsecase(
		txManager,
		rideRepo,
		driverRepo,
	)

	
	rideHandler := rideHandler.NewRideHandler(rideUC)
	


	requestRide := middleware.AuthMiddleware(jwtSvc)(
		middleware.PermissionMiddleware(userRepo, "ride:create")(
			http.HandlerFunc(rideHandler.RequestRide),
		),
	)

	mux.Handle("/rides/request", requestRide)
	return mux
}