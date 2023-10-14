package repository_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/salamanderman234/outsourcing-auth-profile-service/config"
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	helper "github.com/salamanderman234/outsourcing-auth-profile-service/helpers"
	model "github.com/salamanderman234/outsourcing-auth-profile-service/models"
	repository "github.com/salamanderman234/outsourcing-auth-profile-service/repositories"
)

func TestRepositories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repositories Suite")
}

var _ = Describe("Repository functionality", Label("Repository"),func() {
	var exampleName string
	var exampleEmail string
	var examplePassword string

	var dummy1 model.Partner
	var dummy2 model.Partner

	var repo domain.Repository

	BeforeEach(func () {
		// set database config
		config.SetConfig("../.env")
		client, err := config.ConnectDatabase()
		if err != nil {
			panic(err)
		}
		repo = repository.NewRepository(client)
		// set field variable
		// unique := time.Now().UnixNano()	-> ini membuat siapapun yang memakai ini pasti unique walaupun dipakai berulang-ulang
		exampleEmail = fmt.Sprintf("%s@example.com", helper.GenerateRandomString(5))
		exampleName = "salamanderman234"
		examplePassword = "salamanderpassword" 

		dummy1 = model.Partner{
			Email: &exampleEmail,
			Name: &exampleName,
			Password: &examplePassword,
			Avatar: "//",
			About: "adslfkjsalfkajdl",
		}
		dummy2 = model.Partner{
			Name: &exampleName,
			Password: &examplePassword,
			Avatar: "//",
			About: "adslfkjsalfkajdl",
		}
	})

	Describe("Repository.create()", func() {
		When("using correct model (with required field)", func() {
			It("should be created successfully", func(ctx SpecContext) {
				_, err := repo.Create(ctx, &dummy1)
				Expect(err).To(BeNil())
			})
		})
		When("using incorrect model (without required field)", func() {
			It("should not be created successfully", func(ctx SpecContext) {
				_, err := repo.Create(ctx, &dummy2)
				Expect(err).ToNot(BeNil())
			})
		})
		When("duplicate email entry", func() {
			It("should not be created successfully", func(ctx SpecContext) {
				new := model.Partner {
					Email: dummy1.Email,
					Password: dummy1.Password,
					Name: dummy1.Name,
				}
				_, err := repo.Create(ctx, &new)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("Repository.update()", func() {
		var created domain.Model
		var newName string
		var updated model.Partner
		BeforeEach(func(ctx SpecContext) {
			newName = "New name from update"
			updated = model.Partner{
				Name: &newName,
			}
			anotherMail := fmt.Sprintf("%s@example.com", helper.GenerateRandomString(5))
			new := model.Partner {
				Email: &anotherMail,
				Name: &exampleName,
				Password: &examplePassword,
			}
			obj, err := repo.Create(ctx, &new)
			if err != nil {
				panic(err)
			}
			created = obj.(*model.Partner)
		})

		When("using correct data (field and id)", func() {
			It("should be updated successfully", func(ctx SpecContext) {
				id := created.GetID()
				result,aff, err := repo.Update(ctx, id, &updated)
				Expect(err).To(BeNil())
				Expect(aff).To(Equal(1))
				Expect(*result.(*model.Partner).Name).To(Equal(newName))
			})
		})

		When("using incorrect data (wrong id)", func() {
			It("should be not updated successfully", func(ctx SpecContext) {
				id := uint(19833)
				_,aff, _ := repo.Update(ctx, id, &updated)
				Expect(aff).To(Equal(0))
			})
		})
	})

	Describe("Repository.delete()", func() {
		var created domain.Model
		BeforeEach(func(ctx SpecContext) {
			anotherMail := fmt.Sprintf("%s@example.com", helper.GenerateRandomString(5))
			new := model.Partner {
				Email: &anotherMail,
				Name: &exampleName,
				Password: &examplePassword,
			}
			obj, err := repo.Create(ctx, &new)
			if err != nil {
				panic(err)
			}
			created = obj.(*model.Partner)
		})
		
		When("using correct id (data is exist)", func() {
			It("should be deleted successfully", func(ctx SpecContext) {
				id := created.GetID()
				affected, err := repo.Delete(ctx, id, created)
				Expect(err).To(BeNil())
				Expect(affected).To(Equal(1))
			})
		})
		When("using incorrect id (data is not exist)", func() {
			It("should be not deleted successfully", func(ctx SpecContext) {
				id := uint(1872990)
				aff, _ := repo.Delete(ctx, id, created)
				Expect(aff).To(Equal(0))
			})
		})
	})
})
