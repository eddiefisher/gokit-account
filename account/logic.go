package account

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	repository Repository
	logger     log.Logger
}

// NewService realized application logic
func NewService(repository Repository, logger log.Logger) Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")

	user := User{
		Email:    email,
		Password: password,
	}

	uuid, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		_ = level.Error(logger).Log("err", err)
		return "", err
	}

	_ = logger.Log("create user")

	return uuid, nil
}

func (s service) GetUser(ctx context.Context, uuid string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")

	email, err := s.repository.GetUser(ctx, uuid)

	if err != nil {
		_ = level.Error(logger).Log("err", err)
		return "", err
	}

	_ = logger.Log("get user", uuid)

	return email, nil
}
