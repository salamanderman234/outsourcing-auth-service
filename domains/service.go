package domain

import (
	"context"
	"net/http"
)

type CrudService interface {
	Create(ctx context.Context, data Entity) (any, error)
	Get(ctx context.Context, query Entity) (any, error)
	Find(ctx context.Context, id uint, entity Entity) (any, error)
	Update(ctx context.Context, id uint, entity Entity, updatedFields Entity) (any, error)
	Delete(ctx context.Context, id uint, entity Entity) (int, error)
}

type AuthService interface {
	Login(ctx context.Context, creds Entity, model Entity) (*http.Cookie, error)
	Register(ctx context.Context, data Entity) (*http.Cookie, error)
	CheckTokenValid(token *http.Cookie) (Entity, bool, error)
}