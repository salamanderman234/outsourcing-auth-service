package domain

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, data Model) (any, error)
	Get(ctx context.Context, query SearchQueryFunc) (any, error)
	FindById(ctx context.Context, id uint, target Model) (any, error)
	Update(ctx context.Context, id uint, data Model) (any,int, error)
	Delete(ctx context.Context, id uint, target Model) (int, error)
}

type SearchQueryFunc func(ctx context.Context, client *gorm.DB) (any, error)
