package repository

import (
	"context"
	"errors"

	"github.com/go-sql-driver/mysql"
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

func(r repository) Create(ctx context.Context, data domain.Model) (domain.Model,error)  {
	obj := data.GetFillable()
	result := r.client.WithContext(ctx).Create(obj)
	err := result.Error
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		err = domain.ErrDuplicateKey
	}
	return obj, err
}	

func(r repository) Get(ctx context.Context, query domain.SearchQueryFunc) ([]domain.Model, error) {
	return query(ctx, r.client)
}

func(r repository) FindById(ctx context.Context, id uint, target domain.Model) (domain.Model, error) {
	data := target
	result := r.client.WithContext(ctx).Where("id = ?", id).First(&data)
	err := result.Error
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = domain.ErrRecordNotFound
	}
	return data, err
}

func(r repository) Update(ctx context.Context, id uint, data domain.Model) (domain.Model, int, error) {
	obj := data.GetFillable()
	result := r.client.WithContext(ctx).
		Model(data).
		Where("id = ?", id).
		Updates(obj)
	err := result.Error
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = domain.ErrRecordNotFound
	}
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		err = domain.ErrDuplicateKey
	}
	return obj, int(result.RowsAffected),err
}

func(r repository) Delete(ctx context.Context, id uint, target domain.Model) (int, error) {
	result := r.client.WithContext(ctx).
		Delete(target, id)
	err := result.Error
	if result.RowsAffected == 0 {
		err = domain.ErrRecordNotFound
	}
	return int(result.RowsAffected), err
}