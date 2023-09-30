package entity

import domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"

type Credentials struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (c *Credentials) IsEntity() bool {
	return true
}

func (c Credentials) GetObject() domain.Entity {
	return &c
}