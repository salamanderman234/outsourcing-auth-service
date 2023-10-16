package main

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-auth-profile-service/config"
	"github.com/salamanderman234/outsourcing-auth-profile-service/entity"
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
	router := echo.New()
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
	// creating auth viewset
	partnerAuthViewset := view.NewAuthViewset(&entity.PartnerEntity{}, authService, crudService)
	// defining router group
	partnerGroup := router.Group("/partners")
	// defining resource routes
	route.RegisterResourceRoute(partnerGroup, partnerViewset)
	// defining auth routes
	route.RegisterAuthRoute(partnerGroup, partnerAuthViewset)

	// start
	router.Logger.Fatal(router.Start(":1323"))
}