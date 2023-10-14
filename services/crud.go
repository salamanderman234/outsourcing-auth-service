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

func(c *crudService) Create(ctx context.Context, data domain.Entity) (domain.Entity, error) {
	// convert into model
	dataModel := data.GetCorrespondingModel()
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
	return data, nil
}

func(c *crudService) Get(ctx context.Context, query domain.Entity) ([]domain.Entity, error) {
	// convert to model
	dataModel := query.GetCorrespondingModel()
	err := helper.ConvertEntityToModel(query, dataModel)
	if err != nil {
		return nil, domain.ErrConversionDataType
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
		temp := query
		err = helper.ConvertModelToEntity(resultModel, temp)
		if err != nil {
			return nil, domain.ErrConversionDataType
		}
		entityResults[index] = temp
	}
	return entityResults, nil
}

func(c *crudService) Find(ctx context.Context, id uint, group domain.Entity) (domain.Entity, error) {
	// convert into model
	dataModel := group.GetCorrespondingModel()
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
	return group, nil
}

func(c *crudService) Update(ctx context.Context, id uint, updatedFields domain.Entity) (domain.Entity, error) {
	// convert to model
	dataModel := updatedFields.GetCorrespondingModel()
	err := helper.ConvertEntityToModel(updatedFields, dataModel)
	if err != nil {
		return nil, domain.ErrConversionDataType
	}
	// calling repo
	result, _ ,err := c.repo.Update(ctx, id, dataModel)
	if err != nil {
		return nil, err
	}
	// convert back to entity
	err = helper.ConvertModelToEntity(result.(domain.Model), updatedFields)
	if err != nil {
		return nil, domain.ErrConversionDataType
	}
	return updatedFields, nil
}

func(c *crudService) Delete(ctx context.Context, id uint, group domain.Entity) (int, error) {
	// convert to model
	dataModel := group.GetCorrespondingModel()
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