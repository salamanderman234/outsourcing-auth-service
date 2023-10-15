package domain

import (
	"context"
)

type CrudService interface {
	Create(ctx context.Context, data Entity, user AuthEntity) (Entity, error)
	Get(ctx context.Context, query Entity, user AuthEntity) ([]Entity, error)
	Find(ctx context.Context, id uint, group Entity, user AuthEntity) (Entity, error)
	Update(ctx context.Context, id uint, updatedFields Entity, user AuthEntity) (Entity, error)
	Delete(ctx context.Context, id uint, group Entity, user AuthEntity) (int, error)
}

type AuthService interface {
	Login(ctx context.Context, creds AuthEntity) (AuthTokens, error)
	Register(ctx context.Context, data AuthEntity) (AuthTokens, error)
	CheckTokenValid(token string) (JWTClaims, error)
	RenewToken(ctx context.Context, refreshToken string, group AuthEntity) (AuthTokens, error)
}