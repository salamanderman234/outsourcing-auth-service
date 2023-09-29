package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Email    *string `json:"email" query:"email" gorm:"unique;not null;type:varchar(255)"`
	Password *string `json:"password" query:"password" gorm:"not null;type:varchar(32)"`
}