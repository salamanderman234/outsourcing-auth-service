package service_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/salamanderman234/outsourcing-auth-profile-service/config"
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	"github.com/salamanderman234/outsourcing-auth-profile-service/entity"
	helper "github.com/salamanderman234/outsourcing-auth-profile-service/helpers"
	repository "github.com/salamanderman234/outsourcing-auth-profile-service/repositories"
	service "github.com/salamanderman234/outsourcing-auth-profile-service/services"
)

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Services Suite")
}

var _ = Describe("Crud service functionality", Label("Crud Service"), func() {
	
})

var _ = Describe("Partner auth service functionality", Label("Partner Auth Service"), func() {
	var authService domain.PartnerAuthService
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
				token, err := authService.Register(ctx, data)
				Expect(err).To(BeNil())
				_, err = helper.VerifyToken(token.Token)
				Expect(err).To(BeNil())
				_, err = helper.VerifyToken(token.Refresh)
				Expect(err).To(BeNil())
			})
		})
		When("using invalid data (without required field)", func() {
			It("should be not registered successfully", func(ctx SpecContext) {
				exampleName := "asiap"
				data := &entity.PartnerEntity{
					Email: &emailCred,
					Name: &exampleName,
				}
				token, err := authService.Register(ctx, data)
				Expect(err).ToNot(BeNil())
				Expect(token.Token).To(Equal(""))
				Expect(token.Refresh).To(Equal(""))
			})
		})
	})

	Describe("AuthService.Login()", func() {
		When("using valid credentials (with required field)", func() {
			It("should be logged in successfully", func(ctx SpecContext) {
				token, err := authService.Login(ctx, validCreds)
				Expect(err).To(BeNil())
				Expect(token).ToNot(Equal(""))
				_, err = helper.VerifyToken(token.Token)
				Expect(err).To(BeNil())
				_, err = helper.VerifyToken(token.Refresh)
				Expect(err).To(BeNil())
			})
		})
		When("using invalid credentials (missing required field)", func() {
			It("should be not logged in successfully", func(ctx SpecContext) {
				token, err := authService.Login(ctx, invalidCreds)
				Expect(token.Refresh).To(Equal(""))
				Expect(token.Token).To(Equal(""))
				Expect(err).ToNot(BeNil())
			})
		})
		When("using wrong credentials", func() {
			It("should be not logged in successfully", func(ctx SpecContext) {
				token, err := authService.Login(ctx, wrongCreds)
				Expect(token.Refresh).To(Equal(""))
				Expect(token.Token).To(Equal(""))
				Expect(err).To(Equal(domain.ErrInvalidCreds))
			})
		})
	})

	Describe("AuthService.CheckTokenValid()", func() {
		var jwtToken string
		var expiredToken string
		var invalid string

		var email string
		var username string
		var avatar string
		var role string

		// preps
		BeforeEach(func(ctx SpecContext) {
			email = "asi@gmail.com"
			username = "asiap"
			avatar = "//"
			role = "partner"
			token, err := helper.CreateToken(&email, &username, &avatar,&role, uint(1), domain.TokenExpiresAt)
			if err != nil {
				panic(err)
			}
			jwtToken = token
			claims := domain.JWTClaims {
				RegisteredClaims: jwt.RegisteredClaims{
					ID: "1",
					ExpiresAt: jwt.NewNumericDate(time.Now()),
				},
				JWTPayload: domain.JWTPayload{
					Username: &username,
					Email: &email,
					Group: &role,
					Avatar: &avatar,
				},
			}
			invalidToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			signed, err := invalidToken.SignedString([]byte("asipa"))
			if err != nil {
				panic(err)
			}
			expiredToken = signed
			wrongsign, err := invalidToken.SignedString([]byte("wrong"))
			if err != nil {
				panic(err)
			}
			invalid = wrongsign
		})
		// when token is valid token
		When("using valid jwt token", func() {
			It("should be returning user data", func(ctx SpecContext) {
				claims, err := authService.CheckTokenValid(jwtToken)
				Expect(claims).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
		// when token is expires
		When("using invalid jwt token (expires token)", func() {
			It("should be not returning user data", func(ctx SpecContext) {
				_, err := authService.CheckTokenValid(expiredToken)
				Expect(err).To(Equal(domain.ErrTokenIsExpired))
			})
		})
		// when token is not valid
		When("using invalid jwt token (not a valid token)", func() {
			It("should be not returning user data", func(ctx SpecContext) {
				_, err := authService.CheckTokenValid(invalid)
				Expect(err).To(Equal(domain.ErrTokenNotValid))
			})
		})
	})

	Describe("AuthService.RenewToken()", func() {
		var refresh string
		var invalidRefresh string
		var expiresRefresh string
		var invalidIdRefresh string
		BeforeEach(func(ctx SpecContext) {
			unique := time.Now().UnixNano()
			email := fmt.Sprintf("%d@example.com", unique)
			password := "//"
			name := "asiap"
			avatar := "//"
			role := "partner"
			data := entity.PartnerEntity{
				Email: &email,
				Password: &password,
				Name: &name,
			}
			tokens, err := authService.Register(ctx, &data)
			if err != nil {
				panic(err)
			}
			refresh = tokens.Refresh

			claims := domain.JWTClaims {
				RegisteredClaims: jwt.RegisteredClaims{
					ID: "1",
					ExpiresAt: jwt.NewNumericDate(time.Now()),
				},
				JWTPayload: domain.JWTPayload{
					Username: &name,
					Email: &email,
				},
			}
			invalidToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			wrongSigned, err := invalidToken.SignedString([]byte("wrong"))
			if err != nil {
				panic(err)
			}
			invalidRefresh = wrongSigned
			expiresSigned, err := invalidToken.SignedString([]byte("asipa"))
			if err != nil {
				panic(err)
			}
			expiresRefresh = expiresSigned
			invalidIdToken, err := helper.CreateToken(&email, &name, &avatar, &role, 9989, domain.TokenRefreshExpiresAt)
			if err != nil {
				panic(err)
			}
			invalidIdRefresh = invalidIdToken
		})

		When("using valid refresh token", func() {
			It("should be returning a new pair of token", func(ctx SpecContext) {
				tokens, err := authService.RenewToken(ctx, refresh)
				Expect(err).To(BeNil())
				Expect(tokens.Refresh).ToNot(Equal(""))
				Expect(tokens.Token).ToNot(Equal(""))
			})
		})
		When("using invalid refresh token", func() {
			It("should be not returning a new pair of token", func(ctx SpecContext) {
				tokens, err := authService.RenewToken(ctx, invalidRefresh)
				Expect(err).To(Equal(domain.ErrTokenNotValid))
				Expect(tokens.Refresh).To(Equal(""))
				Expect(tokens.Token).To(Equal(""))
			})
		})
		When("using invalid id refresh token", func() {
			It("should be not returning a new pair of token", func(ctx SpecContext) {
				tokens, err := authService.RenewToken(ctx, expiresRefresh)
				Expect(err).To(Equal(domain.ErrTokenIsExpired))
				Expect(tokens.Refresh).To(Equal(""))
				Expect(tokens.Token).To(Equal(""))
			})
		})
		When("using invalid id refresh token", func() {
			It("should be not returning a new pair of token", func(ctx SpecContext) {
				tokens, err := authService.RenewToken(ctx, invalidIdRefresh)
				Expect(err).ToNot(BeNil())
				Expect(tokens.Refresh).To(Equal(""))
				Expect(tokens.Token).To(Equal(""))
			})
		})
	})
})