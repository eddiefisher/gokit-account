package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints exists for service
type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

// MakeEndpoints return endpoints realization
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// typecast request to CreateUserRequest
		req := request.(CreateUserRequest)
		ok, err := s.CreateUser(ctx, req.Email, req.Password)
		return CreateUserResponse{Ok: ok}, err
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUserRequest)
		email, err := s.GetUser(ctx, req.UUID)
		return GetUserResponse{Email: email}, err
	}
}
