package domain

import (
	"context"

	"gorm.io/gorm"
)

type Model interface {	
	IsModel() bool
	GetObject() Model
	GetID() uint
	Search(ctx context.Context, client *gorm.DB) (any, error)
}