package main

import (
	"context"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Date time.Time

type User struct {
	ID       int
	Name     string
	Birthday *Date
}

func (d *Date) Scan(value any) error {
	timeValue := value.(time.Time)
	*d = Date(timeValue)

	return nil
}

func (d *Date) Value() (driver.Value, error) {
	return time.Time(*d), nil
}

func (d *Date) String() string {
	return time.Time(*d).Format(time.RFC822)
}

const getUserByIDQuery = `SELECT id, name, birthday FROM users WHERE id = $1`

func GetUserByID(ctx context.Context, conn *pgx.Conn, id int) (*User, error) {
	row := conn.QueryRow(ctx, getUserByIDQuery, id)

	user := &User{}
	if err := row.Scan(&user.ID, &user.Name, &user.Birthday); err != nil {
		return nil, fmt.Errorf("can't scan user: %w", err)
	}

	return user, nil
}
