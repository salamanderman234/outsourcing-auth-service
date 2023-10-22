package helper

import (
	"strings"

	"github.com/asaskevich/govalidator"
	entity "github.com/salamanderman234/outsourcing-auth-profile-service/entities"
)

func GenerateFieldValidationError(errs []error) []entity.ErrorFieldDetail {
	result := []entity.ErrorFieldDetail{}
	for _, err := range errs {
		errDetail := err.(govalidator.Error)
		fieldError := entity.ErrorFieldDetail{
			Field: strings.ToLower(errDetail.Name),
			Type: errDetail.Validator,
			Detail: err.Error(),
		}
		result = append(result, fieldError)
	}
	return result
}