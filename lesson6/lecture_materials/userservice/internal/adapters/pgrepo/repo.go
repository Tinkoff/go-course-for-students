package pgrepo

import (
	"context"

	"github.com/jackc/pgx/v5"

	"userservice/internal/user"
)

type RepositoryPG struct {
	conn *pgx.Conn
}

func NewRepositoryPG(conn *pgx.Conn) *RepositoryPG {
	return &RepositoryPG{conn: conn}
}

func (r *RepositoryPG) GetUserByID(ctx context.Context, user_id int64) (*user.User, error) {
	q := `select u.id, u.name from users u where u.id = $1`

	u := &user.User{}

	if err := r.conn.QueryRow(ctx, q, user_id).Scan(&u.ID, &u.Name); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *RepositoryPG) AddUser(ctx context.Context, user user.User) (int64, error) {
	q := `insert into users(name) values($1) returning id`

	var user_id int64

	if err := r.conn.QueryRow(ctx, q, user.Name).Scan(&user_id); err != nil {
		return user_id, err
	}

	return user_id, nil

}
