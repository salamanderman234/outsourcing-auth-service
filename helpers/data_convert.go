package helper

import (
	"encoding/json"

	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
)

func ConvertModelToEntity(data domain.Model, target domain.Entity) (error) {
	jsonEncode, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonEncode, target)
	if err != nil {
		return err
	}
	return nil
}

func ConvertEntityToModel(data domain.Entity, target domain.Model) (error) {
	jsonEncode, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonEncode, target)
	if err != nil {
		return err
	}
	return nil
}