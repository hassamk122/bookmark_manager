package repos

import (
	"context"

	"github.com/hassamk122/bookmark_manager/internals/store"
)

type RefreshTokenRepo interface {
	CreateRefreshToken(ctx context.Context, args store.CreateRefreshTokenParams) (store.RefreshToken, error)
	GetTokenByID(ctx context.Context, tokenID int32) (store.RefreshToken, error)
	FindTokenByTokenString(ctx context.Context, token string) (store.RefreshToken, error)
}

type refreshTokenRepo struct {
	queries *store.Queries
}

func NewRefreshTokenRepo(q *store.Queries) *refreshTokenRepo {
	return &refreshTokenRepo{
		queries: q,
	}
}

func (r *refreshTokenRepo) CreateRefreshToken(ctx context.Context, args store.CreateRefreshTokenParams) (store.RefreshToken, error) {
	return r.queries.CreateRefreshToken(ctx, args)
}

func (r *refreshTokenRepo) GetTokenByID(ctx context.Context, tokenID int32) (store.RefreshToken, error) {
	return r.queries.GetTokenByID(ctx, tokenID)
}

func (r *refreshTokenRepo) FindTokenByTokenString(ctx context.Context, token string) (store.RefreshToken, error) {
	return r.queries.GetTokenByTokenString(ctx, token)
}
