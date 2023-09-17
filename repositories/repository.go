package repository

import (
	"context"

	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	"gorm.io/gorm"
)

type repository struct {
	client *gorm.DB
}

func NewRepository(client *gorm.DB) domain.Repository {
	return &repository {
		client: client,
	}
}

func(r repository) Create(ctx context.Context, data any) (error)  {
	result := r.client.WithContext(ctx).Create(data)
	if result != nil {
		return result.Error
	}
	return nil
}
func(r repository) Get(ctx context.Context, filter any) ([]any, error) {
	var data []any
	result := r.client.WithContext(ctx).Model(filter).Where(filter).Find(data)
	if result != nil {
		return nil, result.Error
	}
	return data,nil
}
func(r repository) FindById(ctx context.Context, id uint, model any) (any, error) {
	var data any
	result := r.client.WithContext(ctx).Model(model).Where("id = ?", id).First(data)
	if result != nil {
		return nil, result.Error
	}
	return data,nil
}
func(r repository) Update(ctx context.Context, id uint, data any) (int, error) {
	return 0, nil
}
func(r repository) Delete(ctx context.Context, id uint, model any) (int, error) {
	return 0, nil
}