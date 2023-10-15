package helper_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/salamanderman234/outsourcing-auth-profile-service/entity"
	helper "github.com/salamanderman234/outsourcing-auth-profile-service/helpers"
	model "github.com/salamanderman234/outsourcing-auth-profile-service/models"
)

func TestHelpers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Helpers Suite")
}

var _ = Describe("helper function functionality", func() {
	Describe("data conversion function", func() {
		// var emptyPartnerModel model.Partner
		// var emptyPartnerEntity entity.PartnerEntity
		var examplePartnerModel model.Partner
		var examplePartnerEntity entity.PartnerEntity

		BeforeEach(func() {
			exampleName := "test name"
			examplePartnerModel = model.Partner{
				Name: &exampleName,
			}
			examplePartnerEntity = entity.PartnerEntity{
				Name: exampleName,
			}
		})

		Describe("convertModelToEntity", func() {
			When("given correct data", func() {
				It("should not returning any error", func() {
					var result entity.PartnerEntity
					err := helper.ConvertModelToEntity(&examplePartnerModel, &result)
					Expect(err).To(BeNil())
					Expect(result.Name).To(Equal(examplePartnerModel.Name))
				})
			})
		})
		Describe("convertEntityToModel", func() {
			When("given correct data", func() {
				It("should not returning any error", func() {
					var result model.Partner
					err := helper.ConvertEntityToModel(&examplePartnerEntity, &result)
					Expect(err).To(BeNil())
					Expect(result.Name).To(Equal(examplePartnerEntity.Name))
				})
			})
		})
	})

	Describe("random function", func () {
		maxLength := 10
		When("given correct number", func () {
			It("should return n+2 random string", func () {
				result := helper.GenerateRandomString(maxLength)
				Expect(len(result)).To(Equal(maxLength))
			})
		})
	})

	Describe("jwt function", func() {
		var email string
		var username string
		var avatar string
		var role string
		var id uint

		BeforeEach(func() {
			email = "asiap@gmail.com"
			username = "asiap"
			avatar = "//"
			role = "admin"
			id = 1
		})
		When("given correct data", func() {
			It("should return valid token", func() {
				token, err := helper.CreateToken(&email, &username, &avatar, &role, id, time.Now().Add(time.Duration(30) * time.Minute))
				Expect(err).To(BeNil())
				payload, ok := helper.VerifyToken(token)
				Expect(ok).To(BeNil())
				Expect(payload.Username).ToNot(BeNil())
			})
		})
	})
})
