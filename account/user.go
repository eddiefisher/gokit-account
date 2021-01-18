package account

import "context"

// User ...
type User struct {
	UUID     string `json:"uuid,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Repository its gateway for any database
type Repository interface {
	CreateUser(ctx context.Context, user User) (string, error)
	GetUser(ctx context.Context, uuid string) (string, error)
}
