package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type JWTService interface {
	GenerateToken(userID string)(string, error)
	ValidateToken(token string)(*jwt.RegisteredClaims, error)
}

type jwtService struct {
	secretKey string
	ttl time.Duration 
}

func NewJwtService(secret string, ttl time.Duration) JWTService {
	return &jwtService{
		secretKey: secret,
		ttl: ttl,
	}
}

func (j *jwtService) GenerateToken(userID string) (string, error){
	claims := jwt.RegisteredClaims{
		Subject: userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.ttl)),
		IssuedAt: jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(tokenStr string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.secretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	return token.Claims.(*jwt.RegisteredClaims), nil
}