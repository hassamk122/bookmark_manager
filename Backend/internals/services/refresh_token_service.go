package services

import (
	"context"
	"database/sql"

	"github.com/hassamk122/bookmark_manager/internals/repos"
)

type RefreshTokenService interface {
}

type refreshTokenService struct {
	DB               *sql.DB
	RefreshTokenRepo repos.RefreshTokenRepo
}

func NewRefreshTokenService(db *sql.DB, refreshTokenRepo repos.RefreshTokenRepo) *refreshTokenService {
	return &refreshTokenService{
		DB:               db,
		RefreshTokenRepo: refreshTokenRepo,
	}
}

func (r *refreshTokenService) Generate(ctx context.Context, userID int) (string, error) {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	defer tx.Rollback()

	// will later
	return "", err
}
