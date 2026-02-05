package services

import (
	"crypto/sha256"
	"errors"
	"time"

	"github.com/perdana/sociomile/config"
	"github.com/perdana/sociomile/repositories"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	repo *repositories.UserRepo
}

func NewAuthService(repo *repositories.UserRepo) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil || user.ID == 0 {
		return "", errors.New("invalid credentials")
	}
	hashedPass := sha256.Sum256([]byte(password))
	if string(hashedPass[:]) != user.Password {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"id":        user.ID,
		"role":      user.Role,
		"tenant_id": user.TenantID,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Load().JWTSecret))
}
