package helper

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-auth-profile-service/entity"
)

func CreateBaseResponse(status int, messageType string, message string, datas ...any) entity.BaseResponse {
	var data any
	if len(datas) > 0 {
		data = datas[0]
	}
	return entity.BaseResponse{
		Status: status,
		Type: messageType,
		Message: message,
		Datas: data,
	}
}

func SendResponse (ctx echo.Context, responseStatus int, responseType string, responseMessage string, responseDatas ...any) error{
	
	return ctx.JSON(responseStatus, CreateBaseResponse(responseStatus, responseType, responseMessage, responseDatas, responseDatas))
}