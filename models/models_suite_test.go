package model_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/salamanderman234/outsourcing-auth-profile-service/config"
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	model "github.com/salamanderman234/outsourcing-auth-profile-service/models"
	repository "github.com/salamanderman234/outsourcing-auth-profile-service/repositories"
)

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Models Suite")
}

var _ = Describe("Crud operation using partner", Label("partner"),func() {
	var repo domain.PartnerRepository
	var partner1 model.Partner
	var partner2 model.Partner

	BeforeEach(func () {
		config.SetConfig("../.env")
		client, err := config.ConnectDatabase()
		if err != nil {
			panic(err)
		}
		repo = repository.NewPartnerRepository(client)
		unique := time.Now().UnixNano()
		exampleEmail := fmt.Sprintf("%d@example.com", unique)
		exampleName := "example"
		examplePassword := "secret"
		partner1 = model.Partner{
			Email: &exampleEmail,
			Name: &exampleName,
			Password: &examplePassword,
		}
		partner2 = model.Partner{
			Name: &exampleName,
		}

	})

	When("correct partner model given", func () {
		It("should be created", func (ctx SpecContext) {
			err := repo.Create(ctx, partner1)
			Expect(err).To(BeNil())
		})
	})

	When("incorrect partner model given", func () {
		It("should not be created", func (ctx SpecContext) {
			err := repo.Create(ctx, partner2)
			Expect(err).ToNot(BeNil())
		})
	})

	When("get user by filter", func() {
		It("return array of partner only contain one user", func (ctx SpecContext) {
			id := uint(1)
			data, _ := repo.FindById(ctx, id)
			Expect(data).To(Not(Equal(model.Partner{})))
		})
		It("return array of partners", func (ctx SpecContext) {
			exampleName := "example"
			filter := model.Partner {
				Name: &exampleName,
			}
			data, _ := repo.Get(ctx, filter)
			Expect(len(data) >=2).To(BeTrue())
		})
	})

	When("updating partner", func () {
		It("successfully update user with correct data", func(ctx SpecContext) {
			// create user
			exampleMail := fmt.Sprintf("%d@example.com", time.Now().UnixNano())
			exampleName := "changed from test"
			partner := partner1
			partner.Email = &exampleMail
			err := repo.Create(ctx, partner)
			Expect(err).To(BeNil())
			// search user
			partners, _ := repo.Get(ctx, model.Partner{
				Email: &exampleMail,
			})
			Expect(len(partners) == 1).To(Equal(true))
			// update user
			partner = partners[0]
			affected, err := repo.Update(ctx, partner.ID, model.Partner{
				Name: &exampleName,
			})
			Expect(affected).To(Equal(1))
			result, _ := repo.FindById(ctx, partner.ID)
			Expect(*result.Name).To(Equal(exampleName))
		})
	})

	When("deleting user", func () {
		It("successfully delete user by id", func (ctx SpecContext) {
			// create user
			exampleMail := fmt.Sprintf("%d@example.com", time.Now().UnixNano())
			partner := partner1
			partner.Email = &exampleMail
			err := repo.Create(ctx, partner)
			Expect(err).To(BeNil())
			// search user
			partners, _ := repo.Get(ctx, model.Partner{
				Email: &exampleMail,
			})
			Expect(len(partners) == 1).To(Equal(true))
			// deleting user
			partner = partners[0]
			affected, err := repo.Delete(ctx, partner.ID)
			Expect(affected).To(Equal(1))
			result, _ := repo.FindById(ctx, partner.ID)
			Expect(result).To(Equal(model.Partner{}))
		})
	})
	
})
