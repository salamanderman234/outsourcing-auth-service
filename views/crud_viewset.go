package view

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
)

type crudViewSet struct {
	entity domain.Entity
	crudService domain.CrudService
}

func NewCrudViewSet(entity domain.Entity, service domain.CrudService) domain.CrudViewSet {
	return &crudViewSet{
		entity: entity,
		crudService: service,
	}
} 

func (c *crudViewSet) ResetField() {
	c.entity.ResetField()
}

func (c *crudViewSet) GetCrudService() domain.CrudService{
	return c.crudService
}

func (c *crudViewSet) GetEntity() domain.Entity {
	return c.entity
}
func (c *crudViewSet) SetEntity(entity domain.Entity) {
	c.entity = entity
}
func (c *crudViewSet) Create(ctx echo.Context) error {
	form := c.GetEntity().FormCreateAction()
	successHandler := func() (int, string, string) {
		return http.StatusCreated, domain.ResponseSuccess, "created"
	}
	mainHandler := func(ctx context.Context, data domain.Entity, user domain.JWTClaims) (any, error) {
		return c.crudService.Create(ctx, data, user)
	}
	return crudProcess(ctx, form, mainHandler, successHandler)
}

func (c *crudViewSet) Get(ctx echo.Context) error {
	form := c.GetEntity().FormGetAction()
	successHandler := func() (int, string, string) {
		return http.StatusOK, domain.ResponseSuccess, "ok"
	}
	mainHandler := func(ctx context.Context, data domain.Entity, user domain.JWTClaims) (any, error) {
		return c.crudService.Get(ctx, data, user)
	}
	return crudProcess(ctx, form, mainHandler, successHandler)
}
func (c *crudViewSet) Update(ctx echo.Context) error {
	form := c.GetEntity().FormUpdateAction()
	successHandler := func() (int, string, string) {
		return http.StatusOK, domain.ResponseSuccess, "ok"
	}
	mainHandler := func(ctx context.Context, data domain.Entity, user domain.JWTClaims) (any, error) {
		return c.crudService.Update(ctx, data.GetID(), data, user)
	}
	return crudProcess(ctx, form, mainHandler, successHandler)

}
func (c *crudViewSet) Find(ctx echo.Context) error {
	form := c.GetEntity().FormFindAction()
	successHandler := func() (int, string, string) {
		return http.StatusOK, domain.ResponseSuccess, "ok"
	}
	mainHandler := func(ct context.Context, data domain.Entity, user domain.JWTClaims) (any, error) {
		return c.crudService.Find(ct, data.GetID(), data, user)
	}
	return crudProcess(ctx, form, mainHandler, successHandler)

}
func (c *crudViewSet) Delete(ctx echo.Context) error {
	form := c.GetEntity().FormDeleteAction()
	successHandler := func() (int, string, string) {
		return http.StatusOK, domain.ResponseSuccess, "ok"
	}
	mainHandler := func(ctx context.Context, data domain.Entity, user domain.JWTClaims) (any, error) {
		return c.crudService.Delete(ctx, data.GetID(), data, user)
	}
	return crudProcess(ctx, form, mainHandler, successHandler)
}