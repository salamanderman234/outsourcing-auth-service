package domain

import (
	"context"

	model "github.com/salamanderman234/outsourcing-auth-profile-service/models"
)

type PartnerRepository interface {
	Create(ctx context.Context, data model.Partner) (error)
	Get(ctx context.Context, filter model.Partner) ([]model.Partner, error)
	FindById(ctx context.Context, id uint) (model.Partner, error)
	Update(ctx context.Context, id uint, data model.Partner) (int, error)
	Delete(ctx context.Context, id uint) (int, error)
}

type AdminRepository interface {
	Create(ctx context.Context, data model.Admin) (error)
	Get(ctx context.Context, filter model.Admin) ([]model.Admin, error)
	FindById(ctx context.Context, id uint) (model.Admin, error)
	Update(ctx context.Context, id uint, data model.Admin) (int, error)
	Delete(ctx context.Context, id uint) (int, error)
}

