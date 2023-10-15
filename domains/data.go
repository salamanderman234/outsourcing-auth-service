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
	TokenRefreshName = "refresh"
)

type Model interface {
	IsModel() bool
	GetID() uint
	SetID(id uint)
	SearchQuery(ctx context.Context, client *gorm.DB) ([]Model, error)
	GetFillable() Model
	GetCrudPolicies(action string, user AuthEntity) bool
}

type Entity interface {
	IsEntity() bool
	GetCorrespondingModel() Model
	GetViewable() Entity
	ResetField()
}

type AuthModel interface {	
	Model
	GetUsernameField() string
	GetIdentityField() string
	GetPasswordField() string
	GetAvatarField() string
	GetGroupName() string
	SetPasswordField(value *string) 
	SetUsernameField(value *string)
	SetEmptyID()
}

type AuthEntity interface {
	Entity
	GetCorrespondingAuthModel() AuthModel
	GetUsernameFieldName() string
	RegisterCredsValidate(ctx context.Context) error
	LoginCredsValidate(ctx context.Context) error
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