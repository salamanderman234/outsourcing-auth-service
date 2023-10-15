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
	"gorm.io/gorm"
)

func init() {
	config.SetConfig("../.env")
}

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Services Suite")
}

var _ = Describe("Service test", Label("Service"), func() {
	var client *gorm.DB
	var repo domain.Repository
	BeforeEach(func() {
		client, _ = config.ConnectDatabase()
		// create new repository
		repo = repository.NewRepository(client)
	})
	Describe("Crud service functionality", Label("Crud Service"), func() {
		var crudService domain.CrudService
		var partner entity.PartnerEntity
		var invalidPartner entity.PartnerEntity
		var exampleName string
		var exampleEmail string
		var examplePassword string
	
		BeforeEach(func() {
			// create crud service
			crudService = service.NewCrudService(repo)
			exampleEmail = fmt.Sprintf("%s@email.com",helper.GenerateRandomString(12))
			exampleName = "test"
			examplePassword = "password"
			partner = entity.PartnerEntity{
				Name: exampleName,
				Email: exampleEmail,
				Password: examplePassword,
			}
			invalidPartner = entity.PartnerEntity{
				Name: exampleName,
				Password: examplePassword,
			}
		})
	
		Describe("CrudService.Create()", func () {
			When("given correct entity ", func () {
				It("should return result", func (ctx SpecContext) {
					result, err := crudService.Create(ctx, &partner, &partner)
					obj := result.(*entity.PartnerEntity)
					Expect(obj.Email).To(Equal(exampleEmail))
					Expect(err).To(BeNil())
				})
			})
			When("given incorrect entity (missing required field)", func () {
				It("should return result", func (ctx SpecContext) {
					_, err := crudService.Create(ctx, &invalidPartner, &partner)
					Expect(err).ToNot(BeNil())
				})
			})
		})
		Describe("CrudService.Get()", func () {
			When("given query with existed data", func () {
				It("should return result", func (ctx SpecContext) {
					exampleQuery := entity.PartnerEntity {
						Name: exampleName,
					}
					result, err := crudService.Get(ctx, &exampleQuery , &partner)
					Expect(len(result)).ToNot(Equal(0))
					Expect(err).To(BeNil())
				})
			})
			When("given query with not existed data", func () {
				It("should not return any result", func (ctx SpecContext) {
					notExistedName := "fjasldfjaofhioldsfhkhf"
					exampleQuery := entity.PartnerEntity {
						Name: notExistedName,
					}
					result, err := crudService.Get(ctx, &exampleQuery, &partner)
					fmt.Print(result)
					Expect(len(result)).To(Equal(0))
					Expect(err).To(Equal(domain.ErrRecordNotFound))
				})
			})
		})
		// TODO
		Describe("CrudService.Find()", func () {
			When("given query with existed id", func () {
				It("should return result", func (ctx SpecContext) {
					group := entity.PartnerEntity {}
					_, err := crudService.Find(ctx, 1, &group, &partner)
					Expect(err).To(BeNil())
				})
			})
			When("given query with not existed id", func () {
				It("should not return any result", func (ctx SpecContext) {
					group := entity.PartnerEntity {}
					_, err := crudService.Find(ctx, 8878, &group, &partner)
					Expect(err).To(Equal(domain.ErrRecordNotFound))
				})
			})
		})
		Describe("CrudService.Update()", func () {
			When("updating existed data", func () {
				It("should return result", func (ctx SpecContext) {
					testName := "Change from crud service"
					updateField := entity.PartnerEntity {
						Name: testName,
					}
					result, err := crudService.Update(ctx, 1, &updateField, &partner)
					Expect(result.(*entity.PartnerEntity).Name).To(Equal(testName))
					Expect(err).To(BeNil())
				})
			})
			When("updating non existed data", func () {
				It("should not return any result", func (ctx SpecContext) {
					updatedFields := entity.PartnerEntity {}
					_, err := crudService.Find(ctx, 8878, &updatedFields, &partner)
					Expect(err).To(Equal(domain.ErrRecordNotFound))
				})
			})
		})
		// males bray
		Describe("CrudService.Delete()", func () {
			When("deleting existed data", func () {
				It("should successfully deleting data", func (ctx SpecContext) {
					updatedFields := entity.PartnerEntity {}
					exampleNewMail := helper.GenerateRandomString(11)
					newData := entity.PartnerEntity{
						Email: exampleNewMail,
						Name: exampleName,
						Password: examplePassword,
					}
					data, _ := crudService.Create(ctx, &newData, &partner)
					_, err := crudService.Delete(ctx, data.(*entity.PartnerEntity).ID, &updatedFields, &partner)
					Expect(err).To(BeNil())
				})
			})
			When("deleting non existed data", func () {
				It("should not success deleting any data", func (ctx SpecContext) {
					updatedFields := entity.PartnerEntity {}
					_, err := crudService.Delete(ctx, 887899, &updatedFields, &partner)
					Expect(err).ToNot(BeNil())
				})
			})
		})
	})
	
	// testing pertama gagal karean masalah unique email
	var _ = Describe("Partner auth service functionality", Label("Partner Auth Service"), func() {
		var authService domain.AuthService
		// var validCreds domain.AuthEntity
		var invalidCreds domain.AuthEntity
		var wrongCreds domain.AuthEntity
	
		var emailCred string
		var passCred string
	
		BeforeEach(func() {
			// create auth service
			authService = service.NewAuthService(repo)
	
			emailCred = fmt.Sprintf("%s@example.com", helper.GenerateRandomString(5))
			passCred = "examplepassword"
	
			// validCreds = &entity.PartnerEntity {
			// 	Email: &emailCred,
			// 	Password: &passCred,
			// }
			invalidCreds = &entity.PartnerEntity{
				Email: emailCred,
			}
			invalidPass := "fdsfds"
			invalidEmail := "notfound@gmail.com"
			wrongCreds = &entity.PartnerEntity{
				Email: invalidEmail,
				Password:  invalidPass,
			}
		})
	
		Describe("AuthService.Register()", func() {
			When("using valid data (with required field)", func() {
				It("should be registered successfully", func(ctx SpecContext) {
					exampleName := "asiap"
					data :=  entity.PartnerEntity{
						Email: emailCred,
						Password: passCred,
						Name: exampleName,
					}
					token, err := authService.Register(ctx, &data)
					Expect(err).To(BeNil())
					_, err = helper.VerifyToken(token.Access)
					Expect(err).To(BeNil())
					_, err = helper.VerifyToken(token.Refresh)
					Expect(err).To(BeNil())
				})
			})
			When("using invalid data (without required field)", func() {
				It("should be not registered successfully", func(ctx SpecContext) {
					exampleName := "asiap"
					data := &entity.PartnerEntity{
						Email: emailCred,
						Name: exampleName,
					}
					token, err := authService.Register(ctx, data)
					Expect(err).ToNot(BeNil())
					Expect(token.Access).To(Equal(""))
					Expect(token.Refresh).To(Equal(""))
				})
			})
		})
	
		Describe("AuthService.Login()", func() {
			When("using valid credentials (with required field)", func() {
				It("should be logged in successfully", func(ctx SpecContext) {
					exampleNewMail := helper.GenerateRandomString(11)
					exampleName := "new user from login"
					examplePassword := "password"
					newData := entity.PartnerEntity{
						Email: exampleNewMail,
						Name: exampleName,
						Password: examplePassword,
					}
					tokens, _ := authService.Register(ctx, &newData)
					claims, _ := helper.VerifyToken(tokens.Access)
					creds := entity.PartnerEntity {
						Email: *claims.JWTPayload.Username,
						Password: examplePassword,
					}
					token, err := authService.Login(ctx, &creds)
					Expect(err).To(BeNil())
					Expect(token).ToNot(Equal(""))
					_, err = helper.VerifyToken(token.Access)
					Expect(err).To(BeNil())
					_, err = helper.VerifyToken(token.Refresh)
					Expect(err).To(BeNil())
				})
			})
			When("using invalid credentials (missing required field)", func() {
				It("should be not logged in successfully", func(ctx SpecContext) {
					token, err := authService.Login(ctx, invalidCreds)
					Expect(token.Refresh).To(Equal(""))
					Expect(token.Access).To(Equal(""))
					Expect(err).ToNot(BeNil())
				})
			})
			When("using wrong credentials", func() {
				It("should be not logged in successfully", func(ctx SpecContext) {
					token, err := authService.Login(ctx, wrongCreds)
					Expect(token.Refresh).To(Equal(""))
					Expect(token.Access).To(Equal(""))
					Expect(err).To(Equal(domain.ErrRecordNotFound))
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
						Username: &email,
						Name: &username,
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
				email := fmt.Sprintf("%s@example.com", helper.GenerateRandomString(5))
				password := "//"
				name := "asiap"
				avatar := "//"
				role := "partner"
				data := entity.PartnerEntity{
					Email: email,
					Password: password,
					Name: name,
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
						Name: &name,
						Username: &email,
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
					tokens, err := authService.RenewToken(ctx, refresh, &entity.PartnerEntity{})
					Expect(err).To(BeNil())
					Expect(tokens.Refresh).ToNot(Equal(""))
					Expect(tokens.Access).ToNot(Equal(""))
				})
			})
			When("using invalid refresh token", func() {
				It("should be not returning a new pair of token", func(ctx SpecContext) {
					tokens, err := authService.RenewToken(ctx, invalidRefresh, &entity.PartnerEntity{} )
					Expect(err).To(Equal(domain.ErrTokenNotValid))
					Expect(tokens.Refresh).To(Equal(""))
					Expect(tokens.Access).To(Equal(""))
				})
			})
			When("using invalid id refresh token", func() {
				It("should be not returning a new pair of token", func(ctx SpecContext) {
					tokens, err := authService.RenewToken(ctx, expiresRefresh, &entity.PartnerEntity{})
					Expect(err).To(Equal(domain.ErrTokenIsExpired))
					Expect(tokens.Refresh).To(Equal(""))
					Expect(tokens.Access).To(Equal(""))
				})
			})
			When("using invalid id refresh token", func() {
				It("should be not returning a new pair of token", func(ctx SpecContext) {
					tokens, err := authService.RenewToken(ctx, invalidIdRefresh, &entity.PartnerEntity{})
					Expect(err).ToNot(BeNil())
					Expect(tokens.Refresh).To(Equal(""))
					Expect(tokens.Access).To(Equal(""))
				})
			})
		})
	})
})
