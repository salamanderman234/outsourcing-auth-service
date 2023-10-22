package entity

import domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"

type AdminFindForm struct {
	ID uint `query:"id" param:"id" valid:"required~id is required"`
}

func (p *AdminFindForm) GetCorrespodingEntity() domain.Entity {
	return &AdminEntity{
		ID: p.ID,
	}
}
type AdminGetForm struct {
	ID    uint   `query:"id"`
	Email string `query:"email"`
	Name  string `query:"name"`
}

func (p *AdminGetForm) GetCorrespodingEntity() domain.Entity {
	return &AdminEntity{
		ID: p.ID,
		Email: p.Email,
		Name: p.Name,
	}
}
type AdminCreateForm struct {
	Email    string  `json:"email" query:"email" valid:"required~email is required,email~must be a valid email,length(0|255)~email length must be less than 256 character"`
	Password string  `json:"password,omitempty" query:"password" valid:"required~password is required,length(8|32)~password length must be 8-32 character"`
	Name     string  `json:"name" query:"name" valid:"required~name is required ,length(0|255)~name length must be less than 256 character"`
	Avatar   string  `json:"avatar" query:"avatar" valid:"length(0|255)~avatar length must be less than 256 character"`
}

func (p *AdminCreateForm) GetCorrespodingEntity() domain.Entity {
	return &AdminEntity{
		Email: p.Email,
		Name: p.Name,
		Password: p.Password,
		Avatar: p.Avatar,
	}
}
type AdminUpdateForm struct {
	ID 		 uint    `json:"id" valid:"required~id is required"`
	Email    string  `json:"email" valid:"required~email is required,email~must be a valid email,length(0|255)~email length must be less than 256 character"`
	Name     string  `json:"name" valid:"required~name is required ,length(0|255)~name length must be less than 256 character"`
	Avatar   string  `json:"avatar" valid:"length(0|255)~avatar length must be less than 256 character"`
}

func (p *AdminUpdateForm) GetCorrespodingEntity() domain.Entity {
	return &AdminEntity{
		ID: p.ID,
		Email: p.Email,
		Name: p.Name,
		Avatar: p.Avatar,
	}
}

type AdminDeleteForm struct {
	ID uint `json:"id" valid:"required~id is required"`
}

func (p *AdminDeleteForm) GetCorrespodingEntity() domain.Entity {
	return &AdminEntity{
		ID: p.ID,
	}
}

type AdminRegisterForm struct {
	Email    string  `json:"email" valid:"required~email is required,email~must be a valid email,length(0|255)~email length must be less than 256 character"`
	Password string  `json:"password,omitempty" valid:"required~password is required,length(8|32)~password length must be 8-32 character"`
	Name     string  `json:"name" valid:"required~name is required ,length(0|255)~name length must be less than 256 character"`
	Avatar   string  `json:"avatar"  valid:"length(0|255)~avatar length must be less than 256 character"`
	About    string  `json:"about" valid:"length(0|255)~about length must be less than 256 character"`
}

func (p *AdminRegisterForm) GetCorrespondingAuthEntity() domain.AuthEntity {
	return &AdminEntity{
		Email: p.Email,
		Password: p.Password,
		Name: p.Password,
		Avatar: p.Avatar,
	}
}

type AdminLoginForm struct {
	Email    string  `json:"email" valid:"required~email is required,email~must be a valid email"`
	Password string  `json:"password,omitempty" valid:"required~password is required"`
}


func (p *AdminLoginForm) GetCorrespondingAuthEntity() domain.AuthEntity {
	return &AdminEntity{
		Email: p.Email,
		Password: p.Password,
	}
}