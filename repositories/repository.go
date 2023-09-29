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
	return &repository{
		client: client,
	}
}

func(r repository) Create(ctx context.Context, data domain.Model) (any,error)  {
	obj := data.GetObject()
	result := r.client.WithContext(ctx).Create(obj)
	return obj, result.Error
}	

func(r repository) Get(ctx context.Context, query domain.SearchQueryFunc) (any, error) {
	return query(ctx, r.client)
}

func(r repository) FindById(ctx context.Context, id uint, target domain.Model) (any, error) {
	data := target.GetObject()
	result := r.client.WithContext(ctx).Where("id = ?", id).First(&data)
	return data,result.Error
}

func(r repository) Update(ctx context.Context, id uint, data domain.Model) (any, int, error) {
	obj := data.GetObject()
	result := r.client.WithContext(ctx).
		Model(data).
		Where("id = ?", id).
		Updates(obj)
	return obj, int(result.RowsAffected),result.Error
}

func(r repository) Delete(ctx context.Context, id uint, target domain.Model) (int, error) {
	result := r.client.WithContext(ctx).
		Where("id = ?", id).
		Delete(target.GetObject())
	return int(result.RowsAffected), result.Error
}