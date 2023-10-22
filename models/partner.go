package model

import (
	"context"
	"fmt"

	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	"gorm.io/gorm"
)

type Partner struct {
	gorm.Model
	Email *string `json:"email" query:"email" gorm:"unique;not null;type:varchar(255)"`
	Password *string `json:"password" query:"password" gorm:"not null;type:varchar(255)"`
	Name *string `json:"name" query:"name" gorm:"not null;type:varchar(255)"`
	Avatar string `json:"avatar" query:"avatar" gorm:"default:''"`
	About string `json:"about" query:"about" gorm:"default:''"`
}

func(p *Partner) IsModel() bool {
	return true
}

func (p *Partner) GetFillable() domain.Model {
	p.ID = 0
	return p
}

func (p *Partner) GetGroupName() string {
	return "partner"
}

func(p *Partner) GetID() uint {
	return p.ID
}
func(p *Partner) SetID(id uint) {
	p.ID = id
}

func(p *Partner) GetUsernameField() string {
	if p.Email == nil {
		return ""
	}
	return *p.Email
}

func(p *Partner) GetIdentityField() string {
	if p.Name == nil {
		return ""
	}
	return *p.Name
}

func(p *Partner) GetAvatarField() string {
	return p.Avatar
}

func(p *Partner) GetPasswordField() string {
	if p.Password == nil {
		return ""
	}
	return *p.Password
}

func (p *Partner) SetPasswordField(value *string)  {
	p.Password = value
}
func (p *Partner) SetUsernameField(value *string) {
	p.Email = value
}
func (p *Partner) SetEmptyID() {
	p.ID = 0
}

func(r *Partner) SearchQuery(ctx context.Context, client *gorm.DB) ([]domain.Model, error) {
	query := struct {
		Email string 
		Name string
	} {
		Email: *r.Email,
		Name: *r.Name,
	}
	partners := []Partner{}
	result := client.WithContext(ctx).
		Model(r).
		Where("email LIKE ?", fmt.Sprintf("%%%s%%", query.Email)).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", query.Name)).
		Find(&partners)
	if len(partners) <= 0 {
		return nil, domain.ErrRecordNotFound
	}
	resultModels := make([]domain.Model, len(partners))
	for index := range resultModels {
		resultModels[index] = &partners[index]
	}
	return resultModels, result.Error
}