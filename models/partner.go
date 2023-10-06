package model

import (
	"context"
	"errors"

	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	"gorm.io/gorm"
)

type Partner struct {
	gorm.Model
	Email *string `json:"email" query:"email" gorm:"unique;not null;type:varchar(255)"`
	Password *string `json:"password" query:"password" gorm:"not null;type:varchar(32)"`
	Name *string `json:"name" query:"name" gorm:"not null;type:varchar(255)"`
	Avatar string `json:"avatar" query:"avatar" gorm:"default:''"`
	About string `json:"about" query:"about" gorm:"default:''"`
}

func(p *Partner) IsModel() bool {
	return true
}

func(p *Partner) GetID() uint {
	return p.ID
}

func(p Partner) GetObject() domain.Model {
	return &p
}

func(r *Partner) Search(ctx context.Context, client *gorm.DB) (any, error) {
	partners := []Partner{}
	result := client.WithContext(ctx).
		Model(r).
		Where(r).
		Find(&partners)
	err := result.Error
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = domain.ErrRecordNotFound
	}
	return partners, err
}