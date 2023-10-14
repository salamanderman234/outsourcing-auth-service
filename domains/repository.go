package domain

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, data Model) (Model, error)
	Get(ctx context.Context, query SearchQueryFunc) ([]Model, error)
	FindById(ctx context.Context, id uint, target Model) (Model, error)
	Update(ctx context.Context, id uint, data Model) (Model,int, error)
	Delete(ctx context.Context, id uint, target Model) (int, error)
}

type SearchQueryFunc func(ctx context.Context, client *gorm.DB) ([]Model, error)
