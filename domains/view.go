package domain

import (
	"context"

	echo "github.com/labstack/echo/v4"
)

type ViewSet interface {
	GetEntity() Entity
	SetEntity(entity Entity)
	GetCrudService() CrudService
	ResetField()
}

type CrudViewSet interface {
	ViewSet
	Create(ctx echo.Context) error
	Get(ctx echo.Context) error
	Find(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type AuthViewSet interface {
	GetAuthEntity() AuthEntity
	SetAuthEntity(entity AuthEntity)
	GetAuthService() AuthService
	Login(ctx echo.Context) error
	Register(ctx echo.Context) error
	Verify(ctx echo.Context) error
	Refresh(ctx echo.Context) error
	CreateResetPasswordToken(ctx echo.Context) error
	ResetPassword(ctx echo.Context) error
}

type ValidateProcessFunc func (data Entity) (bool, error)
type CrudMainProcessFunc func(ctx context.Context, data Entity, user JWTClaims) (any, error)
type ResponseSuccessProcessFunc func() (int, string, string)