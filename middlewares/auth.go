package custom_middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	"github.com/salamanderman234/outsourcing-auth-profile-service/entity"
	helper "github.com/salamanderman234/outsourcing-auth-profile-service/helpers"
)

func tokenCheck(ctx echo.Context) (domain.JWTClaims, entity.BaseResponse) {
	token := ctx.Request().Header.Get("Authorization")
	respStatus := http.StatusUnauthorized
	respType := domain.ResponseTokenErr
	respMessage := ""
	respData := []any{}

	sendResp := func () entity.BaseResponse {
		return helper.CreateBaseResponse(respStatus, respType, respMessage, respData)
	}

	if token == "" {
		respMessage = "authorization token is required"
		return domain.JWTClaims{}, sendResp()
	}
	claims, err := helper.VerifyToken(token)
	if errors.Is(err, domain.ErrTokenNotValid) {
		respMessage = "authorization token is not valid"
		return claims, sendResp()
	}
	if errors.Is(err, domain.ErrTokenIsExpired) {
		respMessage = "authorization token is expired"
		return claims, sendResp()
	}
	return claims, entity.BaseResponse{}
}

func WithToken(authEntity domain.AuthEntity) echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func (ctx echo.Context) error {
			respStatus := http.StatusUnauthorized
			respType := domain.ResponseTokenErr
			respMessage := ""
			respData := []any{}
		
			sendResp := func () error {
				return helper.SendResponse(ctx, respStatus, respType, respMessage, respData)
			}
			claims, errResponse := tokenCheck(ctx)
			if errResponse != (entity.BaseResponse{}) {
				return ctx.JSON(errResponse.Status, errResponse)
			}
			group := claims.Group
			if claims.Group != nil {
				if *group != authEntity.GetCorrespondingAuthModel().GetGroupName() {
					respStatus = http.StatusForbidden
					respType = domain.ResponseForbiddenErr
					respMessage = fmt.Sprintf("user is not a %s",authEntity.GetCorrespondingAuthModel().GetGroupName())
					return sendResp()
				}
			}
			ctx.Set("user", claims)
			return next(ctx)
		}
	}
}

func WithoutToken()echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func (ctx echo.Context) error {
			respStatus := http.StatusUnauthorized
			respType := domain.ResponseTokenErr
			respMessage := ""
			respData := []any{}
		
			sendResp := func () error {
				return helper.SendResponse(ctx, respStatus, respType, respMessage, respData)
			}
			if ctx.Request().Header.Get("Authorization") != "" {
				respStatus = http.StatusBadRequest
				respType = domain.ResponseBadRequest
				respMessage = "authorization token is not require"
				sendResp()
			}
			return next(ctx)
		}
	}
}