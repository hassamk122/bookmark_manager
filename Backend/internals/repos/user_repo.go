package repos

import (
	"context"

	"github.com/hassamk122/bookmark_manager/internals/store"
)

type UserRepo interface {
	CreateUser(ctx context.Context, arg store.CreateUserParams) (store.CreateUserRow, error)
	GetUserByEmail(ctx context.Context, email string) (store.GetUserByEmailRow, error)
}

type userRepo struct {
	queries *store.Queries
}

func NewUserRepo(q *store.Queries) *userRepo {
	return &userRepo{
		queries: q,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, arg store.CreateUserParams) (store.CreateUserRow, error) {
	return r.queries.CreateUser(ctx, arg)
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (store.GetUserByEmailRow, error) {
	return r.queries.GetUserByEmail(ctx, email)
}
