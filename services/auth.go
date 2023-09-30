package service

import (
	"context"
	"net/http"

	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
)

type authService struct {
	repo domain.Repository
}

func NewAuthService(repo domain.Repository) domain.AuthService {
	return &authService {
		repo: repo,
	}
}

func(a *authService) Login(ctx context.Context, creds domain.Entity, model domain.Entity) (*http.Cookie, error) {
	return nil, nil
}
func(a *authService) Register(ctx context.Context, data domain.Entity) (*http.Cookie, error) {
	return nil, nil
}
func(a *authService) CheckTokenValid(token *http.Cookie) (domain.Entity, bool, error) {
	return nil, false, nil
}