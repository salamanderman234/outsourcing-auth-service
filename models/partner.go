package model

import "gorm.io/gorm"

type Partner struct {
	gorm.Model
	Email *string `json:"email" query:"email" gorm:"unique;not null;type:varchar(255)"`
	Password *string `json:"password" query:"password" gorm:"not null;type:varchar(32)"`
	Name *string `json:"name" query:"name" gorm:"not null;type:varchar(255)"`
	Avatar string `json:"avatar" query:"avatar" gorm:"default:''"`
	About string `json:"about" query:"about" gorm:"default:''"`
}