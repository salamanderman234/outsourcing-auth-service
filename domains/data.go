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

type Form interface {
	GetCorrespodingEntity() Entity
}

type AuthForm interface {
	GetCorrespondingAuthEntity() AuthEntity
}

type Policy interface {
	CreatePolicy(user JWTClaims) bool
	ReadPolicy(user JWTClaims) bool
	FindPolicy(user JWTClaims, resourceId uint) bool
	UpdatePolicy(user JWTClaims, resourceId uint) bool
	DeletePolicy(user JWTClaims, resourceId uint) bool
}

type Model interface {
	IsModel() bool
	GetID() uint
	SetID(id uint)
	SearchQuery(ctx context.Context, client *gorm.DB) ([]Model, error)
	GetFillable() Model
}

type Entity interface {
	IsEntity() bool
	GetID() uint
	GetCorrespondingModel() Model
	GetViewable() Entity
	GetPolicy() Policy
	ResetField()
	NewObject() Entity
	FormGetAction() Form
	FormFindAction() Form
	FormCreateAction() Form
	FormUpdateAction() Form
	FormDeleteAction() Form
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
	FormRegisterAction() AuthForm
	FormLoginAction() AuthForm
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