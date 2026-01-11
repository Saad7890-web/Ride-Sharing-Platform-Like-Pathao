package router

import (
	"net/http"
	"ride/internal/auth/handler"
	"ride/internal/auth/middleware"
	"ride/internal/auth/repository"
	"ride/internal/auth/service"
	"ride/internal/auth/usecase"
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

	protected := middleware.AuthMiddleware(jwtSvc)(
		middleware.PermissionMiddleware(userRepo, "ride:create")(
			http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request){
				w.Write([]byte("ACCESS GRANTED"))
			}),
		),
	)

	mux.Handle("/protected", protected)
	return mux
}