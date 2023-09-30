package entity

import domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"

type PartnerEntity struct {
	Email    *string `json:"email" query:"email"`
	Password *string `json:"password" query:"password"`
	Name     *string `json:"name" query:"name"`
	Avatar   string  `json:"avatar" query:"avatar"`
	About    string  `json:"about" query:"about"`
}

func (p *PartnerEntity) IsEntity() bool {
	return true
}

func (p PartnerEntity) GetObject() domain.Entity {
	return &p
}