package helper

import (
	"github.com/labstack/echo/v4"
	entity "github.com/salamanderman234/outsourcing-auth-profile-service/entities"
)

func CreateBaseResponse(status int, messageType string, message string, datas entity.BaseResponseDetail) entity.BaseResponse {
	return entity.BaseResponse{
		Status: status,
		Type: messageType,
		Message: message,
		Detail: datas,
	}
}

func SendResponse (ctx echo.Context, responseStatus int, responseType string, responseMessage string, responseDatas entity.BaseResponseDetail) error{
	return ctx.JSON(responseStatus, CreateBaseResponse(responseStatus, responseType, responseMessage, responseDatas))
}