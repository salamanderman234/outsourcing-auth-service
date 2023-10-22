package view

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	entity "github.com/salamanderman234/outsourcing-auth-profile-service/entities"
	helper "github.com/salamanderman234/outsourcing-auth-profile-service/helpers"
)

func crudProcess(
		ctx echo.Context, 
		form domain.Form,
		mainProcessHandler domain.CrudMainProcessFunc,
		successHandler domain.ResponseSuccessProcessFunc,
) error {
	requestContext := ctx.Request().Context()
	respData := entity.BaseResponseDetail{}
	respStatus, respType, respMessage := successHandler()
	var response error
	defer func () {
		response = helper.SendResponse(ctx, respStatus, respType, respMessage, respData)
	}()
	
	if err := ctx.Bind(form); err != nil {
		fmt.Println(err)
		respStatus, respType, respMessage = handleError(domain.ErrBindAndValidation)
		return response
	}
	user, ok := getUserFromSession(ctx)
	if !ok {
		respStatus, respType, respMessage = handleError(domain.ErrUserSessionNotFound)
		return response
	}
	if err := validateForm(form); err != nil {
		errs := helper.GenerateFieldValidationError(err.(govalidator.Errors))
		respStatus, respType, respMessage = handleError(domain.ErrBindAndValidation)
		respData.Errors = errs
		return response
	}
	
	resultData, err := mainProcessHandler(requestContext, form.GetCorrespodingEntity(), user)
	
	if err != nil {
		respStatus, respType, respMessage = handleError(err)
		return response
	}

	respData.Datas = resultData
	return response
}

func getUserFromSession(ctx echo.Context) (domain.JWTClaims,bool){
	// get user session from session
	claims := domain.JWTClaims{}
	claims, okUser := ctx.Get("user").(domain.JWTClaims)
	withUser := ctx.Get("withUser").(bool)
	if withUser {
		return claims, withUser && okUser
	} else {
		return claims, true
	}
}

func handleError(err error) (int, string, string) {
	if errors.Is(err, domain.ErrDuplicateKey) {
		return http.StatusBadRequest, domain.ResponseDuplicateEntries, err.Error()
	}
	if errors.Is(err, domain.ErrInvalidCreds) {
		return http.StatusUnauthorized, domain.ResponseUnauthorizeErr, err.Error()
	}
	if errors.Is(err, domain.ErrPolicies) {
		return http.StatusForbidden, domain.ResponseForbiddenErr, err.Error()
	}
	if errors.Is(err, domain.ErrRecordNotFound) {
		return http.StatusNotFound, domain.ResponseNotFoundErr, err.Error()
	}
	if errors.Is(err, domain.ErrTokenNotValid) || errors.Is(err, domain.ErrTokenIsExpired) {
		return http.StatusUnauthorized, domain.ResponseTokenErr, err.Error()
	}
	if errors.Is(err, domain.ErrBindAndValidation) {
		return http.StatusBadRequest, domain.ResponseBadRequest, err.Error()
	}
	if errors.Is(err, domain.ErrUserSessionNotFound) {
		return http.StatusForbidden, domain.ResponseForbiddenErr, err.Error()
	}
	if errors.Is(err, domain.ErrRefreshCookieNotFound) {
		return http.StatusUnauthorized, domain.ResponseTokenErr, err.Error()
	}
	return http.StatusInternalServerError, domain.ResponseServerErr, "someting went wrong"
}

func validateForm(form any) error {
	_, result := govalidator.ValidateStruct(form)
	return result
}
