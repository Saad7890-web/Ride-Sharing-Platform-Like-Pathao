package usecase

import (
	"context"
	"errors"
	"ride/internal/auth/domain"
	"ride/internal/auth/repository"
	"ride/internal/auth/service"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("Invalid credentials")
	ErrUserInactive = errors.New("user is inactive")
)

type AuthUseCase interface {
	Register(ctx context.Context, email, password string) error
	Login(ctx context.Context, email, password string) (string,error)
}

type authUseCase struct {
	useRepo repository.UserRepository
	jwtSvc service.JWTService
}

func NewAuthUseCase(
	userRepo repository.UserRepository,
	jwtSvc service.JWTService,
) AuthUseCase {
	return &authUseCase{
		useRepo: userRepo,
		jwtSvc: jwtSvc,
	}
}

func (a *authUseCase) Register(ctx context.Context, email, password string)error{
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &domain.User{
		Email: email,
		PasswordHash: string(hash),
		IsActive: true,
	}

	return a.useRepo.Create(user)
}

func (a *authUseCase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := a.useRepo.FindByEmail(email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if !user.IsActive {
		return "", ErrUserInactive
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	); err != nil {
		return "", ErrInvalidCredentials
	}

	return a.jwtSvc.GenerateToken(user.ID)
}