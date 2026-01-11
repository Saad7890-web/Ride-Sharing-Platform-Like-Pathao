package middleware

import (
	"log"
	"net/http"

	"ride/internal/auth/repository"
)

func PermissionMiddleware(
	userRepo repository.UserRepository,
	requiredPermission string,
) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID := GetUserID(r.Context())
			log.Println("USER ID FROM TOKEN:", userID)
			if userID == "" {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			user, err := userRepo.FindByID(userID)
			if err != nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			if !user.HasPermission(requiredPermission) {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
