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
	entity "github.com/salamanderman234/outsourcing-auth-profile-service/entities"
	helper "github.com/salamanderman234/outsourcing-auth-profile-service/helpers"
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

func(a *authViewset) ResetField() {
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
	requestContext := ctx.Request().Context()
	form := a.GetAuthEntity().FormLoginAction()
	respData := entity.BaseResponseDetail{}
	respStatus := http.StatusOK
	respType := domain.ResponseSuccess
	respMessage := "ok"
	var response error
	defer func () {
		response = helper.SendResponse(ctx, respStatus, respType, respMessage, respData)
	}()

	if err := ctx.Bind(form); err != nil {
		fmt.Println(err)
		respStatus, respType, respMessage = handleError(domain.ErrBindAndValidation)
		return response
	}
	if err := validateForm(form); err != nil {
		errs := helper.GenerateFieldValidationError(err.(govalidator.Errors))
		respStatus, respType, respMessage = handleError(domain.ErrBindAndValidation)
		respData.Errors = errs
		return response
	}
	// calling service
	tokens, err := a.authService.Login(requestContext, form.GetCorrespondingAuthEntity())
	if errors.Is(err, domain.ErrInvalidCreds) {
		usernameField := a.group.GetUsernameFieldName()
		textErr := errors.New(fmt.Sprintf("%s or password is wrong", usernameField))
		errs := []error{
			govalidator.Error{Name: usernameField, Validator: "credentials", Err: textErr ,CustomErrorMessageExists: true,},
			govalidator.Error{Name: "password", Validator: "credentials", Err: textErr ,CustomErrorMessageExists: true,},
		}

		respData.Errors = helper.GenerateFieldValidationError(errs)
		respStatus, respType, respMessage = handleError(domain.ErrInvalidCreds)
		return response
	} 
	if err != nil {
		respStatus, respType, respMessage = handleError(err)
		return response
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
	respData.Datas = []any{
		map[string]string{"token" : tokens.Access},
	}
	return response
}
func (a *authViewset) Register(ctx echo.Context) error {
	requestContext := ctx.Request().Context()
	form := a.GetAuthEntity().FormRegisterAction()
	respData := entity.BaseResponseDetail{}
	respStatus := http.StatusCreated
	respType := domain.ResponseSuccess
	respMessage := "created"

	var response error
	defer func () {
		response = helper.SendResponse(ctx, respStatus, respType, respMessage, respData)
	}()

	if err := ctx.Bind(form); err != nil {
		fmt.Println(err)
		respStatus, respType, respMessage = handleError(domain.ErrBindAndValidation)
		return response
	}
	if err := validateForm(form); err != nil {
		errs := helper.GenerateFieldValidationError(err.(govalidator.Errors))
		respStatus, respType, respMessage = handleError(domain.ErrBindAndValidation)
		respData.Errors = errs
		return response
	}
	// calling service
	tokens, err := a.authService.Register(requestContext, form.GetCorrespondingAuthEntity())
	if err != nil {
		respStatus, respType, respMessage = handleError(err)
		return response
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
	respData.Datas = []any{
		map[string]string{"token" : tokens.Access},
	}
	return response
}
func (a *authViewset) Verify(ctx echo.Context) error {
	a.ResetField()
	respData := entity.BaseResponseDetail{}
	respStatus := http.StatusOK
	respType := domain.ResponseSuccess
	respMessage := "user verified"
	
	sendResponse := func () error {
		return helper.SendResponse(ctx, respStatus, respType, respMessage, respData)
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
	respData.Datas = []any {
		user,
	}
	return sendResponse()
}
func (a *authViewset) Refresh(ctx echo.Context) error {
	a.ResetField()
	group := a.GetAuthEntity()
	requestContext := ctx.Request().Context()
	respData := entity.BaseResponseDetail{}
	respType := domain.ResponseSuccess
	respStatus := http.StatusOK
	respMessage := "access token refreshed"
	
	var response error
	defer func () {
		response = helper.SendResponse(ctx, respStatus, respType, respMessage, respData)
	}()

	cookie, err := ctx.Cookie(domain.TokenRefreshName)
	if err != nil {
		respStatus, respType, respMessage = handleError(err)
		return response
	}
	token := cookie.Value
	tokens, err := a.authService.RenewToken(requestContext, token, group)
	if err != nil {
		respStatus, respType, respMessage = handleError(err)
		return response
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
	respData.Datas = []any{
		map[string]string{"token" : tokens.Access},
	}
	return response
}

func (a *authViewset) CreateResetPasswordToken(ctx echo.Context) error {
	a.ResetField()
	requestContext := ctx.Request().Context()
	respData := entity.BaseResponseDetail{}
	respType := domain.ResponseSuccess
	respStatus := http.StatusOK
	respMessage := "reset password token has been sent to requested email"
	
	var response error
	defer func () {
		response = helper.SendResponse(ctx, respStatus, respType, respMessage, respData)
	}()
	body := struct {
		Email string `json:"email" valid:"required~email is required"`
	}{}
	if err := ctx.Bind(&body); err != nil {
		respStatus, respType, respMessage = handleError(err)
		return response
	}
	if err := validateForm(body); err != nil {
		errs := helper.GenerateFieldValidationError(err.(govalidator.Errors))
		respStatus, respType, respMessage = handleError(domain.ErrBindAndValidation)
		respData.Errors = errs
		return response
	}
	err := a.authService.GenerateResetPasswordToken(requestContext,body.Email, a.GetAuthEntity())
	if err != nil {
		respStatus, respType, respMessage = handleError(err)
		return response
	}
	return response
}
func (a *authViewset) ResetPassword(ctx echo.Context) error {
	a.ResetField()
	requestContext := ctx.Request().Context()
	respData := entity.BaseResponseDetail{}
	respType := domain.ResponseSuccess
	respStatus := http.StatusOK
	respMessage := "user password has been changed"
	
	var response error
	defer func () {
		response = helper.SendResponse(ctx, respStatus, respType, respMessage, respData)
	}()
	body := struct {
		Token       string `form:"token" json:"token" valid:"required~reset token is required"`
		NewPassword string `form:"password" json:"password" valid:"required~new password is required"`
		Email string `form:"email" json:"email" valid:"required~email is required"`
	}{}
	if err := ctx.Bind(&body); err != nil {
		respStatus, respType, respMessage = handleError(err)
		return response
	}
	if err := validateForm(body); err != nil {
		errs := helper.GenerateFieldValidationError(err.(govalidator.Errors))
		respStatus, respType, respMessage = handleError(domain.ErrBindAndValidation)
		respData.Errors = errs
		return response
	}
	err := a.authService.ResetPassword(requestContext, body.Token, body.Email ,body.NewPassword, a.GetAuthEntity())
	if err != nil {
		respStatus, respType, respMessage = handleError(err)
		return response
	}
	return response
}