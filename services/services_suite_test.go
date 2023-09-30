package service_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/salamanderman234/outsourcing-auth-profile-service/config"
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	"github.com/salamanderman234/outsourcing-auth-profile-service/entity"
	repository "github.com/salamanderman234/outsourcing-auth-profile-service/repositories"
	service "github.com/salamanderman234/outsourcing-auth-profile-service/services"
)

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Services Suite")
}


var _ = Describe("Auth service functionality", Label("Auth Service"), func() {
	var authService domain.AuthService
	var validCreds domain.Entity
	var invalidCreds domain.Entity
	var wrongCreds domain.Entity

	var emailCred string
	var passCred string

	BeforeEach(func() {
		// set database config
		config.SetConfig("../.env")
		client, err := config.ConnectDatabase()
		if err != nil {
			panic(err)
		}
		// create new repository
		repo := repository.NewRepository(client)
		// create auth service
		authService = service.NewAuthService(repo)

		unique := time.Now().UnixNano()
		emailCred = fmt.Sprintf("%d@example.com", unique)
		passCred = "examplepassword"

		validCreds = &entity.Credentials {
			Email: &emailCred,
			Password: &passCred,
		}
		invalidCreds = &entity.Credentials{
			Email: &emailCred,
		}
		invalidPass := "fdsfds"
		invalidEmail := "notfound@gmail.com"
		wrongCreds = &entity.Credentials{
			Email: &invalidEmail,
			Password: &invalidPass,
		}
	})

	Describe("AuthService.Register()", func() {
		When("using valid data (with required field)", func() {
			It("should be registered successfully", func(ctx SpecContext) {
				exampleName := "asiap"
				data := &entity.PartnerEntity{
					Email: &emailCred,
					Password: &passCred,
					Name: &exampleName,
				}
				cookie, err := authService.Register(ctx, data)
				Expect(err).To(BeNil())
				Expect(cookie).ToNot(BeNil())
			})
		})
		When("using invalid data (without required field)", func() {
			It("should be not registered successfully", func(ctx SpecContext) {
				exampleName := "asiap"
				data := &entity.PartnerEntity{
					Email: &emailCred,
					Name: &exampleName,
				}
				cookie, err := authService.Register(ctx, data)
				Expect(err).ToNot(BeNil())
				Expect(cookie).To(BeNil())
			})
		})
	})

	Describe("AuthService.Login()", func() {
		When("using valid credentials (with required field)", func() {
			It("should be logged in successfully", func(ctx SpecContext) {
				model := &entity.PartnerEntity{}
				cookie, err := authService.Login(ctx, validCreds, model)
				Expect(err).To(BeNil())
				Expect(cookie).ToNot(BeNil())
			})
		})
		When("using invalid credentials (missing required field)", func() {
			It("should be not logged in successfully", func(ctx SpecContext) {
				model := &entity.PartnerEntity{}
				cookie, err := authService.Login(ctx, invalidCreds, model)
				Expect(err).ToNot(BeNil())
				Expect(cookie).To(BeNil())
			})
		})
		When("using wrong credentials", func() {
			It("should be not logged in successfully", func(ctx SpecContext) {
				model := &entity.PartnerEntity{}
				cookie, err := authService.Login(ctx, wrongCreds, model)
				Expect(err).ToNot(BeNil())
				Expect(cookie).To(BeNil())
			})
		})
	})

	Describe("AuthService.CheckTokenValid()", func() {
		var cookie *http.Cookie
		BeforeEach(func(ctx SpecContext) {
			model := &entity.PartnerEntity{}
			cookieResult, err := authService.Login(ctx, validCreds, model)
			if err != nil {
				panic(err)
			}
			cookie = cookieResult
		})
		When("using valid jwt token", func() {
			It("should be returning user data", func(ctx SpecContext) {
				user, valid, err := authService.CheckTokenValid(cookie)
				Expect(user).ToNot(BeNil())
				Expect(valid).To(Equal(true))
				Expect(err).To(BeNil())
			})
		})
		// create invalid token
		When("using invalid jwt token", func() {
			It("should be not returning user data", func(ctx SpecContext) {
				user, valid, err := authService.CheckTokenValid(cookie)
				Expect(user).ToNot(BeNil())
				Expect(valid).To(Equal(true))
				Expect(err).To(BeNil())
			})
		})
	})
})