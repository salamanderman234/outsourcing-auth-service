package entity

import (
	"context"

	"github.com/asaskevich/govalidator"
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	model "github.com/salamanderman234/outsourcing-auth-profile-service/models"
)

type PartnerEntity struct {
	ID 		 uint    `json:"id,omitempty" query:"id"`
	Email    string `json:"email" query:"email" valid:"required~email is required,email~must be a valid email,length(0|255)~email length must be less than 256 character"`
	Password string `json:"password,omitempty" query:"password" valid:"required~password is required,length(8|32)~password length must be 8-32 character"`
	Name     string `json:"name" query:"name" valid:"required~name is required ,length(0|255)~name length must be less than 256 character"`
	Avatar   string  `json:"avatar" query:"avatar"`
	About    string  `json:"about" query:"about" valid:"length(0|255)~about length must be less than 256 character"`
}

func (p *PartnerEntity) ResetField() {
	p.ID = 0
	p.Email = ""
	p.Password = ""
	p.Name = ""
	p.Avatar = ""
	p.About = ""
}

func (p *PartnerEntity) GetUsernameFieldName() string {
	return "email"
}

func (p *PartnerEntity) GetViewable() domain.Entity {
	p.Password = ""
	return p
}

func (p *PartnerEntity) GetCorrespondingModel() domain.Model {
	return &model.Partner{}
}
func (p *PartnerEntity) GetCorrespondingAuthModel() domain.AuthModel {
	return &model.Partner{}
}

func (p *PartnerEntity) IsEntity() bool {
	return true
}

func (p PartnerEntity) RegisterCredsValidate(ctx context.Context) error {
	_, err := govalidator.ValidateStruct(p)
	return err
}

func (p PartnerEntity) LoginCredsValidate(ctx context.Context) error {
	obj := struct{
		Email string `valid:"required~email is required,email~must be a valid email,length(0|255)~email length must be less than 256 character"`
		Password string `valid:"required~password is required"`
	}{
		Email: p.Email,
		Password: p.Password,
	}
	_, err := govalidator.ValidateStruct(obj)
	return err
}