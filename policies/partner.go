package policy

import (
	"strconv"

	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	model "github.com/salamanderman234/outsourcing-auth-profile-service/models"
)

type PartnerPolicy struct {
}

func (p *PartnerPolicy) GeneralPartnerPolicy(user domain.JWTClaims) bool {
	groups := []domain.AuthModel{&model.Admin{}}
	if user.ID == "" {
		return false
	}
	return userGroupValidate(user, groups)
}

func (p *PartnerPolicy) CreatePolicy(user domain.JWTClaims) bool {
	return false
}
func (p *PartnerPolicy) ReadPolicy(user domain.JWTClaims) bool {
	return p.GeneralPartnerPolicy(user)
}
func (p *PartnerPolicy) FindPolicy(user domain.JWTClaims, resourceId uint) bool {
	userId,_ := strconv.Atoi(user.ID)
	return p.GeneralPartnerPolicy(user) || 
		(userGroupValidate(user, []domain.AuthModel{&model.Partner{}}) && userIdValidate(uint(userId), resourceId))
}
func (p *PartnerPolicy) UpdatePolicy(user domain.JWTClaims, resourceId uint) bool {
	userId,_ := strconv.Atoi(user.ID)
	return p.GeneralPartnerPolicy(user) || 
		(userGroupValidate(user, []domain.AuthModel{&model.Partner{}}) && userIdValidate(uint(userId), resourceId))
}
func (p *PartnerPolicy) DeletePolicy(user domain.JWTClaims, resourceId uint) bool {
	return p.GeneralPartnerPolicy(user)
}