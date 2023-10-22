package domain

import (
	"context"
)

type CrudService interface {
	Create(ctx context.Context, data Entity, user JWTClaims) (Entity, error)
	Get(ctx context.Context, query Entity, user JWTClaims) ([]Entity, error)
	Find(ctx context.Context, id uint, group Entity, user JWTClaims) (Entity, error)
	Update(ctx context.Context, id uint, updatedFields Entity, user JWTClaims) (Entity, error)
	Delete(ctx context.Context, id uint, group Entity, user JWTClaims) (int, error)
}

type AuthService interface {
	Login(ctx context.Context, creds AuthEntity) (AuthTokens, error)
	Register(ctx context.Context, data AuthEntity) (AuthTokens, error)
	CheckTokenValid(token string) (JWTClaims, error)
	RenewToken(ctx context.Context, refreshToken string, group AuthEntity) (AuthTokens, error)
	GenerateResetPasswordToken(ctx context.Context, email string, obj AuthEntity) (error)
	ResetPassword(ctx context.Context, token string, email string, newPassword string, obj AuthEntity) (error)
}