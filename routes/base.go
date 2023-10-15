package route

import (
	"github.com/labstack/echo/v4"
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	custom_middleware "github.com/salamanderman234/outsourcing-auth-profile-service/middlewares"
)

func RegisterResourceRoute(
	router *echo.Group, 
	viewset domain.CrudViewSet, 
) {
	router.POST("/", viewset.Create)
	router.GET("/", viewset.Get)
	router.PATCH("/", viewset.Update)
	router.DELETE("/", viewset.Delete)
}

func RegisterAuthRoute(
	router *echo.Group,
	authViewset domain.AuthViewSet,
) {
	router.POST("/login", authViewset.Login, custom_middleware.WithoutToken())
	router.POST("/register", authViewset.Register, custom_middleware.WithoutToken())
	router.GET("/verify", authViewset.Verify, custom_middleware.WithToken(authViewset.GetAuthEntity()))
	router.POST("/refresh", authViewset.Refresh)
}