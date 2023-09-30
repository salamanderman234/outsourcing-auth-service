package service

import (
	"context"

	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
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
	return nil, nil
}
func(c *crudService) Get(ctx context.Context, query domain.Entity) (any, error) {
	return nil, nil
}
func(c *crudService) Find(ctx context.Context, id uint, entity domain.Entity) (any, error) {
	return nil, nil
}
func(c *crudService) Update(ctx context.Context, id uint, entity domain.Entity, updatedFields domain.Entity) (any, error) {
	return nil, nil
}
func(c *crudService) Delete(ctx context.Context, id uint, entity domain.Entity) (int, error) {
	return 0, nil
}