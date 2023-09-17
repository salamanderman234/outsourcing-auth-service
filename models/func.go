package model

import domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"

func GetAllModel() []domain.Model {
	return []domain.Model{
		&Partner{},
	}
}