package domain

import (
	"context"
)

type CrudService interface {
	Create(ctx context.Context, data Entity) (any, error)
	Get(ctx context.Context, query Entity) (any, error)
	Find(ctx context.Context, id uint, entity Entity) (any, error)
	Update(ctx context.Context, id uint, updatedFields Entity) (any, error)
	Delete(ctx context.Context, id uint, entity Entity) (int, error)
}

type AuthService interface {
	Login(ctx context.Context, creds AuthEntity) (AuthTokens, error)
	Register(ctx context.Context, data AuthEntity) (AuthTokens, error)
	CheckTokenValid(token string) (JWTClaims, error)
	RenewToken(ctx context.Context, refreshToken string, group AuthEntity) (AuthTokens, error)
}