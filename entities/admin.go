package entity

import (
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	model "github.com/salamanderman234/outsourcing-auth-profile-service/models"
	policy "github.com/salamanderman234/outsourcing-auth-profile-service/policies"
)

type AdminEntity struct {
	ID 		 uint    `json:"id,omitempty" query:"id"`
	Email    string `json:"email" query:"email" valid:"required~email is required,email~must be a valid email,length(0|255)~email length must be less than 256 character"`
	Password string `json:"password,omitempty" query:"password" valid:"required~password is required,length(8|32)~password length must be 8-32 character"`
	Name     string `json:"name" query:"name" valid:"required~name is required ,length(0|255)~name length must be less than 256 character"`
	Avatar   string  `json:"avatar" query:"avatar" valid:"length(0|255)~avatar length must be less than 256 character"`
}

func (p *AdminEntity) ResetField() {
	p.ID = 0
	p.Email = ""
	p.Password = ""
	p.Name = ""
	p.Avatar = ""
}

func (p *AdminEntity) NewObject() domain.Entity {
	return &AdminEntity{}
}

func (p *AdminEntity) GetID() uint {
	return p.ID
}

func (p *AdminEntity) GetUsernameFieldName() string {
	return "email"
}

func (p *AdminEntity) GetViewable() domain.Entity {
	p.Password = ""
	return p
}

func (p *AdminEntity) GetCorrespondingModel() domain.Model {
	return &model.Admin{}
}
func (p *AdminEntity) GetCorrespondingAuthModel() domain.AuthModel {
	return &model.Admin{}
}

func (p *AdminEntity) IsEntity() bool {
	return true
}

func (p *AdminEntity) GetPolicy() domain.Policy {
	return &policy.AdminPolicy{}
}

func (p *AdminEntity) FormGetAction() domain.Form {
	return &AdminGetForm{}
}
func (p *AdminEntity) FormCreateAction() domain.Form {
	return &AdminCreateForm{}
}
func (p *AdminEntity) FormUpdateAction() domain.Form {
	return &AdminUpdateForm{}
}
func (p *AdminEntity) FormDeleteAction() domain.Form {
	return &AdminDeleteForm{}
}
func (p *AdminEntity) FormFindAction() domain.Form {
	return &AdminFindForm{}
}

func (p *AdminEntity) FormRegisterAction() domain.AuthForm {
	return &AdminRegisterForm{}
}
func (p *AdminEntity) FormLoginAction() domain.AuthForm {
	return &AdminLoginForm{}
}