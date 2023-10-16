package view

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	"github.com/salamanderman234/outsourcing-auth-profile-service/entity"
	helper "github.com/salamanderman234/outsourcing-auth-profile-service/helpers"
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
	data := c.entity
	requestContext := ctx.Request().Context()
	responseData := entity.BaseResponseDetail{}
	responseStatus := http.StatusCreated
	responseMessage := "created"

	sendResponse := func () error{
		return ctx.JSON(responseStatus, helper.CreateBaseResponse(responseStatus, "test",responseMessage, responseData))
	}

	if err := ctx.Bind(&data); err != nil {
		responseStatus = http.StatusBadRequest
		responseMessage = "request body is required"
		return sendResponse()
	}
	user := ctx.Get("user")
	if user == nil {
		responseStatus = http.StatusBadRequest
		responseMessage = "request body is required"
		return sendResponse()
	}
	userEntity := user.(domain.AuthEntity)
	new, err := c.crudService.Create(requestContext, data, userEntity)
	if errors.Is(err, domain.ErrDuplicateKey) {
		responseStatus = http.StatusBadRequest
		responseMessage = "duplicate entries"
		return sendResponse()
	}
	if err != nil {
		responseStatus = http.StatusInternalServerError
		responseMessage = "someting went wrong"
		return sendResponse()
	}
	responseData.Datas = new
	return sendResponse()
}
func (c *crudViewSet) Get(ctx echo.Context) error {
	return nil
}
func (c *crudViewSet) Update(ctx echo.Context) error {
	return nil
}
func (c *crudViewSet) Delete(ctx echo.Context) error {
	return nil
}