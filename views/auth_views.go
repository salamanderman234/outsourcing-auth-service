package view

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	helper "github.com/salamanderman234/outsourcing-auth-profile-service/helpers"
	validator "github.com/salamanderman234/outsourcing-auth-profile-service/validators"
)

type authViewset struct {
	group domain.AuthEntity
	authService domain.AuthService
	crudService domain.CrudService
}

func NewAuthViewset(group domain.AuthEntity, authService domain.AuthService, crudService domain.CrudService) domain.AuthViewSet {
	return &authViewset {
		group: group,
		authService: authService,
		crudService: crudService,
	}
}

func(a *authViewset) resetField() {
	a.group.ResetField()
}

func (a *authViewset) GetAuthEntity() domain.AuthEntity {
	return a.group
}
func (a *authViewset) SetAuthEntity(entity domain.AuthEntity) {
	a.group = entity
}
func (a *authViewset) GetAuthService() domain.AuthService {
	return a.authService
}
func (a *authViewset) Login(ctx echo.Context) error {
	a.resetField()
	creds := a.group
	requestContext := ctx.Request().Context()
	respData := []any{}
	respType := domain.ResponseSuccess
	respStatus := http.StatusOK
	respMessage := "login success"
	
	sendResponse := func () error {
		return helper.SendResponse(ctx, respStatus, respType, respMessage, respData...)
	}

	if err := ctx.Bind(&creds); err != nil {
		respStatus = http.StatusBadRequest
		respType = domain.ResponseBadRequest
		respMessage = "login credentials is required"
		return sendResponse()
	}
	if err := creds.LoginCredsValidate(requestContext); err != nil {
		errs := validator.GenerateFieldValidationError(err.(govalidator.Errors))
		respStatus = http.StatusBadRequest
		respType = domain.ResponseValidationErr
		respMessage = "request data does not comply with the rules"
		respData = []any{
			errs,
		}
		return sendResponse()
	}
	// calling service
	tokens, err := a.authService.Login(requestContext, creds)
	if errors.Is(err, domain.ErrInvalidCreds) {
		usernameField := a.group.GetUsernameFieldName()
		textErr := errors.New(fmt.Sprintf("%s or password is wrong", usernameField))
		errs := []error{
			govalidator.Error{Name: usernameField, Validator: "credentials", Err: textErr ,CustomErrorMessageExists: true,},
			govalidator.Error{Name: "password", Validator: "credentials", Err: textErr ,CustomErrorMessageExists: true,},
		}
		respData = []any {
			validator.GenerateFieldValidationError(errs),
		}
		respStatus = http.StatusUnauthorized
		respType = domain.ResponseUnauthorizeErr
		respMessage = textErr.Error()
		return sendResponse()
	} 
	if err != nil {
		respStatus = http.StatusInternalServerError
		respType = domain.ResponseServerErr
		respMessage = "something went wrong"
		return sendResponse()
	}
	// creating refresh cookie
	refreshCookie := http.Cookie {
		Name: domain.TokenRefreshName,
		Value: tokens.Refresh,
		HttpOnly: true,
		Expires: domain.TokenRefreshExpiresAt,
	}
	ctx.SetCookie(&refreshCookie)
	// creating response contain access token
	respData = []any{
		map[string]string{"token" : tokens.Access},
	}
	return sendResponse()
}
func (a *authViewset) Register(ctx echo.Context) error {
	a.resetField()
	creds := a.group
	requestContext := ctx.Request().Context()
	var respData any
	respType := domain.ResponseSuccess
	respStatus := http.StatusOK
	respMessage := "register success"
	
	sendResponse := func () error {
		return helper.SendResponse(ctx, respStatus, respType, respMessage, respData)
	}

	if err := ctx.Bind(&creds); err != nil {
		respStatus = http.StatusBadRequest
		respType = domain.ResponseBadRequest
		respMessage = "register data is required"
		return sendResponse()
	}
	// creds validate
	if err := creds.RegisterCredsValidate(requestContext); err != nil {
		errs := validator.GenerateFieldValidationError(err.(govalidator.Errors))
		respStatus = http.StatusBadRequest
		respType = domain.ResponseValidationErr
		respMessage = "request data does not comply with the rules"
		respData = []any{
			errs,
		}
		return sendResponse()
	}
	// calling service
	tokens, err := a.authService.Register(requestContext, creds)
	if errors.Is(err, domain.ErrDuplicateKey) {
		respStatus = http.StatusConflict
		respType = domain.ResponseDuplicateEntries
		respMessage = "user already exists"
		return sendResponse()
	}
	if err != nil {
		respStatus = http.StatusInternalServerError
		respType = domain.ResponseServerErr
		respMessage = "something went wrong"
		return sendResponse()
	}
	// creating refresh cookie
	refreshCookie := http.Cookie {
		Name: domain.TokenRefreshName,
		Value: tokens.Refresh,
		HttpOnly: true,
		Expires: domain.TokenRefreshExpiresAt,
	}
	ctx.SetCookie(&refreshCookie)
	// creating response contain access token
	respData = []any{
		map[string]string{"token" : tokens.Access},
	}
	return sendResponse()
}
func (a *authViewset) Verify(ctx echo.Context) error {
	a.resetField()
	respData := []any{}
	respType := domain.ResponseSuccess
	respStatus := http.StatusOK
	respMessage := "user verified"
	
	sendResponse := func () error {
		return helper.SendResponse(ctx, respStatus, respType, respMessage, respData...)
	}

	userClaims := ctx.Get("user")
	id := userClaims.(domain.JWTClaims).ID
	idint, _ := strconv.Atoi(id)
	user := struct {
		ID uint `json:"id"`
		Username string `json:"username"`
		Identity string `json:"name"`
		Group string `json:"group"`
		Avatar string `json:"avatar"`
	}{
		ID: uint(idint),
		Identity: *userClaims.(domain.JWTClaims).Name,
		Username: *userClaims.(domain.JWTClaims).Username,
		Group: strings.ToLower(*userClaims.(domain.JWTClaims).Group),
		Avatar: *userClaims.(domain.JWTClaims).Avatar,
	}
	respData = []any{
		user,
	}
	return sendResponse()
}
func (a *authViewset) Refresh(ctx echo.Context) error {
	a.resetField()
	group := a.group
	requestContext := ctx.Request().Context()
	respData := []any{}
	respType := domain.ResponseSuccess
	respStatus := http.StatusOK
	respMessage := "access token refreshed"
	
	sendResponse := func () error {
		return helper.SendResponse(ctx, respStatus, respType, respMessage, respData...)
	}

	cookie, err := ctx.Cookie(domain.TokenRefreshName)
	if err != nil {
		respStatus = http.StatusUnauthorized
		respType = domain.ResponseTokenErr
		respMessage = "refresh token is required"
		return sendResponse()
	}
	token := cookie.Value
	tokens, err := a.authService.RenewToken(requestContext, token, group)
	if errors.Is(err, domain.ErrTokenNotValid) {
		respType = domain.ResponseTokenErr
		respMessage = "refresh token is invalid"
		return sendResponse()
	}
	if errors.Is(err, domain.ErrTokenIsExpired) {
		respStatus = http.StatusUnauthorized
		respType = domain.ResponseTokenErr
		respMessage = "refresh token is expired"
		return sendResponse()
	}
	if err != nil {
		respStatus = http.StatusInternalServerError
		respType = domain.ResponseTokenErr
		respMessage = "something went wrong"
		return sendResponse()
	}
	// creating refresh cookie
	refreshCookie := http.Cookie {
		Name: domain.TokenRefreshName,
		Value: tokens.Refresh,
		HttpOnly: true,
		Expires: domain.TokenRefreshExpiresAt,
	}
	ctx.SetCookie(&refreshCookie)
	// creating response contain access token
	respData = []any{
		map[string]string{"token" : tokens.Access},
	}
	return sendResponse()
}