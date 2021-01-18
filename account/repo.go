package account

// connect to database
import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/jackc/pgx/v4"
)

// ErrRepo ...
var ErrRepo = errors.New("unable to handle Repo Request")

type repo struct {
	db     *pgx.Conn
	logger log.Logger
}

// NewRepo realized db model actions
func NewRepo(db *pgx.Conn, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (r repo) CreateUser(ctx context.Context, user User) (string, error) {
	sql := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING uuid`
	if user.Email == "" || user.Password == "" {
		return "", ErrRepo
	}

	var uuid string
	err := r.db.QueryRow(ctx, sql, user.Email, user.Password).Scan(&uuid)
	if err != nil {
		return "", err
	}

	return uuid, nil
}

func (r repo) GetUser(ctx context.Context, uuid string) (string, error) {
	var email string

	err := r.db.QueryRow(ctx, "SELECT email FROM users WHERE uuid=$1", uuid).Scan(&email)
	if err != nil {
		return "", ErrRepo
	}

	return email, nil
}
