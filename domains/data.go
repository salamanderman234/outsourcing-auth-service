package domain

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var (
	TokenExpiresAt = time.Now().Add(time.Duration(30) * time.Minute)
	TokenRefreshExpiresAt = time.Now().Add(time.Duration(72) * time.Hour)
	TokenCookieName = "token"
)

type Model interface {	
	IsModel() bool
	GetID() uint
	Search(ctx context.Context, client *gorm.DB) (any, error)
}

type Entity interface {
	IsEntity() bool
}

type JWTPayload struct {
	Username *string `json:"username"`
    Email    *string `json:"email"`
    Group    *string `json:"group"`
	Avatar *string `json:"avatar"`
}

type JWTClaims struct {
	jwt.RegisteredClaims
	JWTPayload
}

type AuthTokens struct {
	Refresh string
	Token string
}