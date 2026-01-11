package repository

import "ride/internal/auth/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByID(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	AssignRole(userID string, roleID int) error
}
