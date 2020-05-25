package users

import (
	"context"
	"github.com/go-kit/kit/log"
)

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, email string) User
}

type repo struct {
	Logger log.Logger
}

func (r repo) CreateUser(ctx context.Context, user User) error {
	log.With(r.Logger,"method","createUser")
	return nil
}

func (r repo) GetUser(ctx context.Context, email string) User {
	log.With(r.Logger,"method","GetUser")
	return User{
		Email: email,
	}
}

func NewRepo(logger log.Logger) Repository {
	return &repo{
		Logger: logger,
	}
}
