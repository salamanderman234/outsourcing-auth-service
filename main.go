package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-auth-profile-service/config"
	entity "github.com/salamanderman234/outsourcing-auth-profile-service/entities"
	custom_middleware "github.com/salamanderman234/outsourcing-auth-profile-service/middlewares"
	repository "github.com/salamanderman234/outsourcing-auth-profile-service/repositories"
	route "github.com/salamanderman234/outsourcing-auth-profile-service/routes"
	service "github.com/salamanderman234/outsourcing-auth-profile-service/services"
	view "github.com/salamanderman234/outsourcing-auth-profile-service/views"
)

func init() {
	config.SetConfig("./.env")
}
func main() {
	// router
	appVersion := config.GetAppVersion()
	router := echo.New()
	router.Use(custom_middleware.SetWithUser(false))
	// connect database
	client, err := config.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	// create repo
	repo := repository.NewRepository(client)
	// create service
	crudService := service.NewCrudService(repo)
	authService := service.NewAuthService(repo)
	// creating crud viewset
	partnerViewset := view.NewCrudViewSet(&entity.PartnerEntity{}, crudService)
	adminViewset := view.NewCrudViewSet(&entity.AdminEntity{}, crudService)
	// creating auth viewset
	partnerAuthViewset := view.NewAuthViewset(&entity.PartnerEntity{}, authService, crudService)
	adminAuthViewset := view.NewAuthViewset(&entity.AdminEntity{}, authService, crudService)
	// defining router group
	versionGroup := router.Group(fmt.Sprintf("/v%s", appVersion))
	partnerGroup := versionGroup.Group("/partners")
	partnerWithTokenGroup := partnerGroup.Group("", custom_middleware.WithToken(&entity.AdminEntity{}, &entity.PartnerEntity{}))
	adminGroup := versionGroup.Group("/admins")
	adminWithTokenGroup := adminGroup.Group("", custom_middleware.WithToken(&entity.AdminEntity{}))
	// defining resource routes
	route.RegisterResourceRoute(partnerWithTokenGroup, partnerViewset)
	route.RegisterResourceRoute(adminWithTokenGroup, adminViewset)
	// defining auth routes
	route.RegisterAuthRoute(partnerGroup, partnerAuthViewset, false)
	route.RegisterAuthRoute(adminGroup, adminAuthViewset, true)
	// start
	router.Logger.Fatal(router.Start(":1323"))
}