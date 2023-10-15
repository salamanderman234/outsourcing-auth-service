package domain

import echo "github.com/labstack/echo/v4"

type ViewSet interface {
	GetEntity() Entity
	SetEntity(entity Entity)
	GetCrudService() CrudService
}

type CrudViewSet interface {
	ViewSet
	Create(ctx echo.Context) error
	Get(ctx echo.Context) error
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
}