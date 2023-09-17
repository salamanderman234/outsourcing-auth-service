package model_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	model "github.com/salamanderman234/outsourcing-auth-profile-service/models"
	repository "github.com/salamanderman234/outsourcing-auth-profile-service/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Models Suite")
}

var _ = Describe("Crud operation using partner", Label("partner"),func() {
	var repo domain.Repository
	var partner1 *model.Partner
	var partner2 *model.Partner

	BeforeEach(func () {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", 
			"gorm",
			"gorm",
			"localhost",
			"3306",
			"outsourcing-app",
		)
		client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		repo = repository.NewRepository(client)
		unique := time.Now().Unix()
		partner1 = &model.Partner{
			Email: fmt.Sprintf("%d@example.com", unique),
			Name: "example",
			Password: "secret",
		}
		partner2 = &model.Partner{
			Name: "example",
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

	When("duplicate email entries", func () {
		It("should not be created", func (ctx SpecContext) {
			err := repo.Create(ctx, partner1)
			Expect(err).To(Equal(gorm.ErrDuplicatedKey))
		})
	})

	
})
