package custom_middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	entity "github.com/salamanderman234/outsourcing-auth-profile-service/entities"
	helper "github.com/salamanderman234/outsourcing-auth-profile-service/helpers"
)

func tokenCheck(ctx echo.Context) (domain.JWTClaims, bool ,entity.BaseResponse) {
	token := ctx.Request().Header.Get("Authorization")
	respStatus := http.StatusUnauthorized
	respType := domain.ResponseTokenErr
	respMessage := ""
	respData := entity.BaseResponseDetail{}

	sendResp := func () entity.BaseResponse {
		return helper.CreateBaseResponse(respStatus, respType, respMessage, respData)
	}

	if token == "" {
		respMessage = "authorization token is required"
		return domain.JWTClaims{}, false ,sendResp()
	}
	claims, err := helper.VerifyToken(token)
	if errors.Is(err, domain.ErrTokenNotValid) {
		respMessage = "authorization token is not valid"
		return claims, false ,sendResp()
	}
	if errors.Is(err, domain.ErrTokenIsExpired) {
		respMessage = "authorization token is expired"
		return claims, false ,sendResp()
	}
	return claims, true, entity.BaseResponse{}
}

func WithToken(authEntity ...domain.AuthEntity) echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func (ctx echo.Context) error {
			respStatus := http.StatusUnauthorized
			respType := domain.ResponseTokenErr
			respMessage := ""
			respData := entity.BaseResponseDetail{}
		
			sendResp := func () error {
				return helper.SendResponse(ctx, respStatus, respType, respMessage, respData)
			}
			claims, ok ,errResponse := tokenCheck(ctx)
			if !ok {
				return ctx.JSON(errResponse.Status, errResponse)
			}
			if len(authEntity) >= 1 {
				group := claims.Group
				if claims.Group != nil {
					valid := false
					for _, ent := range authEntity {
						if *group == ent.GetCorrespondingAuthModel().GetGroupName() {
							valid = true
						}
					}
					if !valid {
						respStatus = http.StatusForbidden
						respType = domain.ResponseForbiddenErr
						respMessage = fmt.Sprintf("user dont have access to this resource")
						return sendResp()
					}
				}
			}
			ctx.Set("withUser", true)
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
			respData := entity.BaseResponseDetail{}
		
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

func SetWithUser(value bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("withUser", value)
			return next(c)
		}
	}
}