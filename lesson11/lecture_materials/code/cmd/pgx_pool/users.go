package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID       int
	Name     string
	Birthday time.Time
}

const getUserByIDQuery = `SELECT id, name, birthday FROM users WHERE id = $1`

func GetUserByID(ctx context.Context, pool *pgxpool.Pool, id int) (*User, error) {
	row := pool.QueryRow(ctx, getUserByIDQuery, id)

	user := &User{}
	if err := row.Scan(&user.ID, &user.Name, &user.Birthday); err != nil {
		return nil, fmt.Errorf("can't scan user: %w", err)
	}

	return user, nil
}
