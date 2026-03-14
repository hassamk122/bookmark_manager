package services

import (
	"context"
	"database/sql"
	"log"

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

	log.Println("Starting database transaction")

	defer tx.Rollback()

	qtx := store.New(tx)
	repo := repos.NewUserRepo(qtx)

	log.Println("Trying to verify if mail exists")

	_, err = repo.GetUserByEmail(ctx, email)
	if err == nil {
		return customerr.ErrEmailTaken
	}

	log.Println("Trying to hash password")
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	log.Println("Trying to create user")
	_, err = repo.CreateUser(ctx, store.CreateUserParams{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	})

	if err != nil {
		log.Printf("Error creating user. Reason: %v", err)
		return err
	}

	return tx.Commit()
}
