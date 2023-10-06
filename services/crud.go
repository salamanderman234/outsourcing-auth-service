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

func(c *crudService) Create(ctx context.Context, data domain.Entity) (any, error) {
	var dataModel domain.Model
	err := helper.ConvertEntityToModel(data, dataModel)
	if err != nil {
		return nil, err
	}
	result, err := c.repo.Create(ctx, dataModel)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func(c *crudService) Get(ctx context.Context, query domain.Entity) (any, error) {
	var dataModel domain.Model
	err := helper.ConvertEntityToModel(query, dataModel)
	if err != nil {
		return nil, err
	}
	result, err := c.repo.Get(ctx, dataModel.SearchQuery)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func(c *crudService) Find(ctx context.Context, id uint, entity domain.Entity) (any, error) {
	var dataModel domain.Model
	err := helper.ConvertEntityToModel(entity, dataModel)
	if err != nil {
		return nil, err
	}
	result, err := c.repo.FindById(ctx, id, dataModel)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func(c *crudService) Update(ctx context.Context, id uint, updatedFields domain.Entity) (any, error) {
	var dataModel domain.Model
	err := helper.ConvertEntityToModel(updatedFields, dataModel)
	if err != nil {
		return nil, err
	}
	result, _ ,err := c.repo.Update(ctx, id, dataModel)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func(c *crudService) Delete(ctx context.Context, id uint, entity domain.Entity) (int, error) {
	var dataModel domain.Model
	err := helper.ConvertEntityToModel(entity, dataModel)
	if err != nil {
		return 0, err
	}
	affected ,err := c.repo.Delete(ctx, id, dataModel)
	if err != nil {
		return 0, err
	}
	return affected, nil
}