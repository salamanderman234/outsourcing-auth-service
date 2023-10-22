package policy

import (
	"strconv"

	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	model "github.com/salamanderman234/outsourcing-auth-profile-service/models"
)

type AdminPolicy struct {
}

func (p *AdminPolicy) GeneralAdminPolicy(user domain.JWTClaims) bool {
	groups := []domain.AuthModel{&model.Admin{}}
	if user.ID == "" {
		return false
	}
	return userGroupValidate(user, groups)
}

func (p *AdminPolicy) CreatePolicy(user domain.JWTClaims) bool {
	return false
}
func (p *AdminPolicy) ReadPolicy(user domain.JWTClaims) bool {
	return p.GeneralAdminPolicy(user)
}
func (p *AdminPolicy) FindPolicy(user domain.JWTClaims, resourceId uint) bool {
	return p.GeneralAdminPolicy(user) 
}
func (p *AdminPolicy) UpdatePolicy(user domain.JWTClaims, resourceId uint) bool {
	userId,_ := strconv.Atoi(user.ID)
	return (userGroupValidate(user, []domain.AuthModel{&model.Admin{}}) && userIdValidate(uint(userId), resourceId))
}
func (p *AdminPolicy) DeletePolicy(user domain.JWTClaims, resourceId uint) bool {
	return false
}