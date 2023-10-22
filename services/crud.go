package service

import (
	"context"

	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	helper "github.com/salamanderman234/outsourcing-auth-profile-service/helpers"
)

type crudService struct {
	repo domain.Repository
}

func NewCrudService(repo domain.Repository) domain.CrudService {
	return &crudService {
		repo: repo,
	}
}

func(c *crudService) Create(ctx context.Context, data domain.Entity, user domain.JWTClaims) (domain.Entity, error) {
	// convert into model
	dataModel := data.GetCorrespondingModel()
	dataModel.SetID(0)
	if !data.GetPolicy().CreatePolicy(user) {
		return nil, domain.ErrPolicies
	}
	err := helper.ConvertEntityToModel(data, dataModel)
	if err != nil {
		return nil, domain.ErrConversionDataType
	}
	// calling repository
	result, err := c.repo.Create(ctx, dataModel)
	if err != nil {
		return nil, err
	}
	// convert back to entity
	err = helper.ConvertModelToEntity(result.(domain.Model), data)
	if err != nil {
		return nil, domain.ErrConversionDataType
	}
	return data.GetViewable(), nil
}

func(c *crudService) Get(ctx context.Context, query domain.Entity, user domain.JWTClaims) ([]domain.Entity, error) {
	// convert to model
	dataModel := query.GetCorrespondingModel()
	err := helper.ConvertEntityToModel(query, dataModel)
	if err != nil {
		return nil, domain.ErrConversionDataType
	}
	if !query.GetPolicy().ReadPolicy(user) {
		return nil, domain.ErrPolicies
	}
	// calling repository
	result, err := c.repo.Get(ctx, dataModel.SearchQuery)
	if err != nil {
		return nil, err
	}
	// convert back to entity
	resultModels := result
	entityResults := make([]domain.Entity, len(resultModels))
	for index, resultModel := range resultModels {
		temp := query.NewObject()
		err = helper.ConvertModelToEntity(resultModel, temp)
		if err != nil {
			return nil, domain.ErrConversionDataType
		}
		entityResults[index] = temp.GetViewable()
	}
	return entityResults, nil
}

func(c *crudService) Find(ctx context.Context, id uint, group domain.Entity, user domain.JWTClaims) (domain.Entity, error) {
	// convert into model
	dataModel := group.GetCorrespondingModel()
	if !group.GetPolicy().FindPolicy(user, id) {
		return nil, domain.ErrPolicies
	}
	err := helper.ConvertEntityToModel(group, dataModel)
	if err != nil {
		return nil, domain.ErrConversionDataType
	}
	// calling repository
	result, err := c.repo.FindById(ctx, id, dataModel)
	if err != nil {
		return nil, err
	}
	// convert back to entity
	err = helper.ConvertModelToEntity(result.(domain.Model), group)
	if err != nil {
		return nil, domain.ErrConversionDataType
	}
	return group.GetViewable(), nil
}

func(c *crudService) Update(ctx context.Context, id uint, updatedFields domain.Entity, user domain.JWTClaims) (domain.Entity, error) {
	// convert to model
	dataModel := updatedFields.GetCorrespondingModel()
	dataModel.SetID(0)
	if !updatedFields.GetPolicy().UpdatePolicy(user, id) {
		return nil, domain.ErrPolicies
	}
	err := helper.ConvertEntityToModel(updatedFields, dataModel)
	if err != nil {
		return nil, domain.ErrConversionDataType
	}
	// calling repo
	result, rowsAffected ,err := c.repo.Update(ctx, id, dataModel)
	if rowsAffected == 0 {
		return nil, domain.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	// convert back to entity
	resultEntity := updatedFields.NewObject()
	err = helper.ConvertModelToEntity(result, resultEntity)
	if err != nil {
		return nil, domain.ErrConversionDataType
	}
	return resultEntity, nil
}

func(c *crudService) Delete(ctx context.Context, id uint, group domain.Entity, user domain.JWTClaims) (int, error) {
	// convert to model
	dataModel := group.GetCorrespondingModel()
	if !group.GetPolicy().DeletePolicy(user, id) {
		return 0, domain.ErrPolicies
	}
	err := helper.ConvertEntityToModel(group, dataModel)
	if err != nil {
		return 0, domain.ErrConversionDataType
	}
	// calling repo
	affected ,err := c.repo.Delete(ctx, id, dataModel)
	if err != nil {
		return 0, err
	}
	return affected, nil
}