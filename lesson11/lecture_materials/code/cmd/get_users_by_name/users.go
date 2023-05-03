package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID       int
	Name     string
	Birthday time.Time
}

const getUsersByNameQuery = `SELECT id, name, birthday FROM users WHERE name = $1`

func GetUsersByName(ctx context.Context, conn *pgx.Conn, name string) ([]User, error) {
	rows, err := conn.Query(ctx, getUsersByNameQuery, name)
	if err != nil {
		return nil, fmt.Errorf("can't select users by name: %w", err)
	}

	defer rows.Close()

	var result []User
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Birthday); err != nil {
			return nil, fmt.Errorf("can't scan user: %w", err)
		}

		result = append(result, user)
	}

	return result, nil
}
