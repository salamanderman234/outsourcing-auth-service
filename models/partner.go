package model

import "gorm.io/gorm"

type Partner struct {
	gorm.Model
	Email string `json:"email" query:"email" gorm:"unique"`
	Password string `json:"password" query:"password"`
	Name string `json:"name" query:"name"`
	Avatar string `json:"avatar" query:"avatar" gorm:"default:''"`
	About string `json:"about" query:"about" gorm:"default''"`
}