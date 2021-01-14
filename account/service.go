package account

import "context"

// Service ...
type Service interface {
	CreateUser(ctx context.Context, email, password string) (string, error)
	GetUser(ctx context.Context, uuid string) (string, error)
}
