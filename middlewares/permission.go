package custom_middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	helper "github.com/salamanderman234/outsourcing-auth-profile-service/helpers"
)

func OnlyUsersDefined(authEntity ...domain.AuthEntity) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			respStatus := http.StatusUnauthorized
			respType := domain.ResponseTokenErr
			respMessage := ""
			respData := []any{}

			sendResp := func() error {
				return helper.SendResponse(ctx, respStatus, respType, respMessage, respData)
			}
			user := ctx.Get("user")
			if user == nil {
				respMessage = "missing authorize user"
			}
			valid := false
			for _, userEntity := range authEntity {
				if *user.(domain.JWTClaims).Group == userEntity.GetCorrespondingAuthModel().GetGroupName() {
					valid = true
					break
				}
			}
			if !valid {
				respStatus = http.StatusForbidden
				respType = domain.ResponseForbiddenErr
				respMessage = "user does not have permission to access this resource"
				return sendResp()
			}
			return next(ctx)
		}
	}
}