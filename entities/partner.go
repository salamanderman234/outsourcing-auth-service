package entity

import (
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	model "github.com/salamanderman234/outsourcing-auth-profile-service/models"
	policy "github.com/salamanderman234/outsourcing-auth-profile-service/policies"
)

type PartnerEntity struct {
	ID 		 uint    `json:"id"`
	Email    string  `json:"email"`
	Password string  `json:"password,omitempty"`
	Name     string  `json:"name"`
	Avatar   string  `json:"avatar"`
	About    string  `json:"about"`
}

func (p *PartnerEntity) ResetField() {
	p.ID = 0
	p.Email = ""
	p.Password = ""
	p.Name = ""
	p.Avatar = ""
	p.About = ""
}

func (p *PartnerEntity) NewObject() domain.Entity {
	return &PartnerEntity{}
}
func (p *PartnerEntity) GetID() uint {
	return p.ID
}

func (p *PartnerEntity) GetPolicy() domain.Policy {
	return &policy.PartnerPolicy{}
}

func (p *PartnerEntity) FormGetAction() domain.Form {
	return &PartnerGetForm{}
}

func (p *PartnerEntity) FormFindAction() domain.Form {
	return &PartnerFindForm{}
}

func (p *PartnerEntity) FormCreateAction() domain.Form {
	return &PartnerCreateForm{}
}

func (p *PartnerEntity) FormUpdateAction() domain.Form {
	return &PartnerUpdateForm{}
}
func (p *PartnerEntity) FormDeleteAction() domain.Form {
	return &PartnerDeleteForm{}
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

func (p *PartnerEntity) FormRegisterAction() domain.AuthForm {
	return &PartnerRegisterForm{}
}
func (p *PartnerEntity) FormLoginAction() domain.AuthForm {
	return &PartnerLoginForm{}
}