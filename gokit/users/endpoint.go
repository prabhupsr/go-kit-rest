package users

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func CreateEndPoint(s Service) *Endpoints {

	return &Endpoints{
		CreateUser: makeCreateUserEndPoint(s),
		GetUser:    makeGetUserEndPoint(s),
	}
}

func makeGetUserEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateGetUserRequest)
		user := s.GetUser(ctx, req.Email)
		return CreateGetUserResponse{Email: user.Email, Name: user.Name}, nil
	}
}

func makeCreateUserEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateUserRequest)
		user, err := s.CreateUser(ctx, req.Name, req.Email)
		return CreateUserResponse{Email: user.Email, Name: user.Name}, err
	}
}
