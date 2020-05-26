package users

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Service interface {
	CreateUser(ctx context.Context, name string, email string) (User, error)
	GetUser(ctx context.Context, email string) User
}

type service struct {
	Repo   Repository
	Logger log.Logger
}

func (s service) CreateUser(ctx context.Context, name string, email string) (User, error) {
	s.Logger.Log(name, email)
	user := User{
		Name:  name,
		Email: email,
	}
	fmt.Println(user.Name)
	return user, nil
}

func (s service) GetUser(ctx context.Context, email string) User {
	logger := log.With(s.Logger, "METHOD", "GetUser")
	s.Logger.Log(email)
	level.Error(logger).Log("err", nil)
	return User{
		Email: email,
	}

}

func NewService(R Repository, logger log.Logger) Service {
	return &service{
		Repo:   R,
		Logger: logger,
	}

}
