package model

import "gorm.io/gorm"

type Partner struct {
	Email string `json:"email" query:"email" gorm:"unique"`
	Password string `json:"password" query:"password"`
	Name string `json:"name" query:"name"`
	Avatar string `json:"avatar" query:"avatar" gorm:"default:''"`
	About string `json:"about" query:"about" gorm:"default''"`
	gorm.Model
}

func (p *Partner) IsModel() bool {
	return true
}