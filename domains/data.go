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
	SearchQuery(ctx context.Context, client *gorm.DB) ([]Model, error)

}

type Entity interface {
	IsEntity() bool
	GetCorrespondingModel() Model
}

type AuthModel interface {	
	Model
	GetUsernameField() string
	GetIdentityField() string
	GetPasswordField() string
	GetAvatarField() string
	SetPasswordField(value *string) 
	SetUsernameField(value *string)
	SetEmptyID()
}

type AuthEntity interface {
	Entity
	GetCorrespondingAuthModel() AuthModel
	CheckRequiredRegisterField() bool
	CheckRequiredLoginField() bool
}

type JWTPayload struct {
	Name 		*string `json:"name"`
    Username    *string `json:"username"`
    Group    	*string `json:"group"`
	Avatar 		*string `json:"avatar"`
}

type JWTClaims struct {
	jwt.RegisteredClaims
	JWTPayload
}

type AuthTokens struct {
	Refresh string
	Access string
}