package entity

import (
	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	model "github.com/salamanderman234/outsourcing-auth-profile-service/models"
)

type PartnerEntity struct {
	ID 		 uint    `json:"id" query:"id"`
	Email    *string `json:"email" query:"email"`
	Password *string `json:"password" query:"password"`
	Name     *string `json:"name" query:"name"`
	Avatar   string  `json:"avatar" query:"avatar"`
	About    string  `json:"about" query:"about"`
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
func (p *PartnerEntity) CheckRequiredRegisterField() bool {
	if p.Email != nil && p.Name != nil && p.Password != nil {
		return true
	}
	return false
}
func (p *PartnerEntity) CheckRequiredLoginField() bool {
	if p.Email != nil && p.Password != nil {
		return true
	}
	return false
}