package domain

import "context"

type Repository interface {
	Create(ctx context.Context, data any) (error) 
	Get(ctx context.Context, filter any) ([]any, error)
	FindById(ctx context.Context, id uint, model any) (any, error)
	Update(ctx context.Context, id uint, data any) (int, error)
	Delete(ctx context.Context, id uint, model any) (int, error)
}