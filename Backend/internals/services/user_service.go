package services

import (
	"context"
	"database/sql"

	customerr "github.com/hassamk122/bookmark_manager/internals/custom_err"
	"github.com/hassamk122/bookmark_manager/internals/repos"
	"github.com/hassamk122/bookmark_manager/internals/store"
	"github.com/hassamk122/bookmark_manager/internals/utils"
)

type UserService interface {
	Register(ctx context.Context, username, email, password string) error
}

type userService struct {
	DB       *sql.DB
	UserRepo repos.UserRepo
}

func NewUserService(db *sql.DB, userRepo repos.UserRepo) *userService {
	return &userService{
		DB:       db,
		UserRepo: userRepo,
	}
}

func (s *userService) Register(ctx context.Context, username, email, password string) error {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	qtx := store.New(tx)
	repo := repos.NewUserRepo(qtx)

	_, err = repo.GetUserByEmail(ctx, email)
	if err != nil {
		return customerr.ErrEmailTaken
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	_, err = repo.CreateUser(ctx, store.CreateUserParams{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	})

	return tx.Commit()
}
