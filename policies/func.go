package policy

import domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"

func userGroupValidate(user domain.JWTClaims, userGroups []domain.AuthModel) bool {
	valid := false
	for _, userModel := range userGroups {
		if *user.Group == userModel.GetGroupName() {
			valid = true
		}
	}
	return valid
}

func userIdValidate(userId uint, resourceId uint) bool {
	return userId == resourceId
}