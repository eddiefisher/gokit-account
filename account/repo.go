package account

// connect to database
import (
	"context"
	"errors"
	"fmt"

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

func (r repo) CreateUser(ctx context.Context, user User) error {
	sql := `insert into users (email, password) values ($1, $2)`
	if user.Email == "" || user.Password == "" {
		return ErrRepo
	}

	ptag, err := r.db.Exec(ctx, sql, user.Email, user.Password)
	if err != nil {
		return err
	}

	fmt.Println(ptag.String())

	return nil
}

func (r repo) GetUser(ctx context.Context, uuid string) (string, error) {
	var email string

	err := r.db.QueryRow(ctx, "select email from users where uuid=$1", uuid).Scan(&email)
	if err != nil {
		return "", ErrRepo
	}

	return email, nil
}
