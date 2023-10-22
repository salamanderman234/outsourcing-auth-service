package model

import (
	"context"
	"fmt"

	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Email *string `json:"email" query:"email" gorm:"unique;not null;type:varchar(255)"`
	Password *string `json:"password" query:"password" gorm:"not null;type:varchar(255)"`
	Name *string `json:"name" query:"name" gorm:"not null;type:varchar(255)"`
	Avatar string `json:"avatar" query:"avatar" gorm:"default:''"`
}

func(p *Admin) IsModel() bool {
	return true
}

func (p *Admin) GetFillable() domain.Model {
	p.ID = 0
	return p
}

func (p *Admin) GetGroupName() string {
	return "admin"
}

func(p *Admin) GetID() uint {
	return p.ID
}
func(p *Admin) SetID(id uint) {
	p.ID = id
}

func(p *Admin) GetUsernameField() string {
	if p.Email == nil {
		return ""
	}
	return *p.Email
}

func(p *Admin) GetIdentityField() string {
	if p.Name == nil {
		return ""
	}
	return *p.Name
}

func(p *Admin) GetAvatarField() string {
	return p.Avatar
}

func(p *Admin) GetPasswordField() string {
	if p.Password == nil {
		return ""
	}
	return *p.Password
}

func (p *Admin) SetPasswordField(value *string)  {
	p.Password = value
}
func (p *Admin) SetUsernameField(value *string) {
	p.Email = value
}
func (p *Admin) SetEmptyID() {
	p.ID = 0
}

func(r *Admin) SearchQuery(ctx context.Context, client *gorm.DB) ([]domain.Model, error) {
	query := struct {
		Email string 
		Name string
	} {
		Email: *r.Email,
		Name: *r.Name,
	}
	admins := []Admin{}
	result := client.WithContext(ctx).
		Model(r).
		Where("email LIKE ?", fmt.Sprintf("%%%s%%", query.Email)).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", query.Name)).
		Find(&admins)
	if len(admins) <= 0 {
		return nil, domain.ErrRecordNotFound
	}
	resultModels := make([]domain.Model, len(admins))
	for index := range resultModels {
		resultModels[index] = &admins[index]
	}
	return resultModels, result.Error
}