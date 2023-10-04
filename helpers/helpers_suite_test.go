package helper_test

import (
	"testing"

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
		// var examplePartnerEntity entity.PartnerEntity

		BeforeEach(func() {
			exampleName := "test name"
			examplePartnerModel = model.Partner{
				Name: &exampleName,
			}
			// examplePartnerEntity = entity.PartnerEntity{
			// 	Name: &exampleName,
			// }
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
	})
})
