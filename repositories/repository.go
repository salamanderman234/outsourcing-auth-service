package repository

import (
	"context"
	"errors"

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
	obj := data
	result := r.client.WithContext(ctx).Create(obj)
	return obj, result.Error
}	

func(r repository) Get(ctx context.Context, query domain.SearchQueryFunc) (any, error) {
	return query(ctx, r.client)
}

func(r repository) FindById(ctx context.Context, id uint, target domain.Model) (any, error) {
	data := target
	result := r.client.WithContext(ctx).Where("id = ?", id).First(&data)
	err := result.Error
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = domain.ErrRecordNotFound
	}
	return data, err
}

func(r repository) Update(ctx context.Context, id uint, data domain.Model) (any, int, error) {
	obj := data
	result := r.client.WithContext(ctx).
		Model(data).
		Where("id = ?", id).
		Updates(obj)
	err := result.Error
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = domain.ErrRecordNotFound
	}
	return obj, int(result.RowsAffected),err
}

func(r repository) Delete(ctx context.Context, id uint, target domain.Model) (int, error) {
	result := r.client.WithContext(ctx).
		Where("id = ?", id).
		Delete(target)
	err := result.Error
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = domain.ErrRecordNotFound
	}
	return int(result.RowsAffected), err
}