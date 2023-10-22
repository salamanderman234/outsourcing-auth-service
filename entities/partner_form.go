package entity

import (
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
)

type PartnerFindForm struct {
	ID uint `query:"id" param:"id" valid:"required~id is required"`
}

func (p *PartnerFindForm) GetCorrespodingEntity() domain.Entity {
	return &PartnerEntity{
		ID: p.ID,
	}
}

type PartnerGetForm struct {
	Email string `query:"email"`
	Name  string `query:"name"`
}

func (p *PartnerGetForm) GetCorrespodingEntity() domain.Entity {
	return &PartnerEntity{
		Email: p.Email,
		Name: p.Name,
	}
}
type PartnerCreateForm struct {
	Email    string  `json:"email" valid:"required~email is required,email~must be a valid email,length(0|255)~email length must be less than 256 character"`
	Password string  `json:"password,omitempty"  valid:"required~password is required,length(8|32)~password length must be 8-32 character"`
	Name     string  `json:"name" valid:"required~name is required ,length(0|255)~name length must be less than 256 character"`
	Avatar   string  `json:"avatar" valid:"length(0|255)~avatar length must be less than 256 character"`
	About    string  `json:"about" valid:"length(0|255)~about length must be less than 256 character"`
}

func (p *PartnerCreateForm) GetCorrespodingEntity() domain.Entity {
	return &PartnerEntity{
		Email: p.Email,
		Name: p.Name,
		Password: p.Password,
		Avatar: p.Avatar,
		About: p.About,
	}
}
type PartnerUpdateForm struct {
	ID 		 uint    `json:"id" valid:"required~id is required"`
	Email    string  `json:"email" valid:"required~email is required,email~must be a valid email,length(0|255)~email length must be less than 256 character"`
	Name     string  `json:"name" valid:"required~name is required ,length(0|255)~name length must be less than 256 character"`
	Avatar   string  `json:"avatar" valid:"length(0|255)~avatar length must be less than 256 character"`
	About    string  `json:"about" valid:"length(0|255)~about length must be less than 256 character"`
}

func (p *PartnerUpdateForm) GetCorrespodingEntity() domain.Entity {
	return &PartnerEntity{
		ID: p.ID,
		Email: p.Email,
		Name: p.Name,
		Avatar: p.Avatar,
		About: p.About,
	}
}

type PartnerDeleteForm struct {
	ID uint `json:"id" valid:"required~id is required"`
}

func (p *PartnerDeleteForm) GetCorrespodingEntity() domain.Entity {
	return &PartnerEntity{
		ID: p.ID,
	}
}

type PartnerRegisterForm struct {
	Email    string  `json:"email" valid:"required~email is required,email~must be a valid email,length(0|255)~email length must be less than 256 character"`
	Password string  `json:"password,omitempty" valid:"required~password is required,length(8|32)~password length must be 8-32 character"`
	Name     string  `json:"name" valid:"required~name is required ,length(0|255)~name length must be less than 256 character"`
	Avatar   string  `json:"avatar"  valid:"length(0|255)~avatar length must be less than 256 character"`
	About    string  `json:"about" valid:"length(0|255)~about length must be less than 256 character"`
}

func (p *PartnerRegisterForm) GetCorrespondingAuthEntity() domain.AuthEntity {
	return &PartnerEntity{
		Email: p.Email,
		Password: p.Password,
		Name: p.Password,
		Avatar: p.Avatar,
		About: p.About,
	}
}

type PartnerLoginForm struct {
	Email    string  `json:"email" valid:"required~email is required,email~must be a valid email"`
	Password string  `json:"password,omitempty" valid:"required~password is required"`
}

func (p *PartnerLoginForm) GetCorrespondingAuthEntity() domain.AuthEntity {
	return &PartnerEntity{
		Email: p.Email,
		Password: p.Password,
	}
}