package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"ride/internal/auth/service"
)

type contextKey string

const userIDKey contextKey = "user_id"

func AuthMiddleware(jwtSvc service.JWTService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "missing authorization header", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid auth header", http.StatusUnauthorized)
				return
			}

			claims, err := jwtSvc.ValidateToken(parts[1])
			if err != nil {
				log.Println("JWT INVALID:", err)
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}
			log.Println("JWT SUBJECT:", claims.Subject)

			ctx := context.WithValue(r.Context(), userIDKey, claims.Subject)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserID(ctx context.Context) string {
	id, _ := ctx.Value(userIDKey).(string)
	return id
}
