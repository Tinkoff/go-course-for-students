package queries

import (
	"context"
	"fmt"
	"pg-course/internal/domain"
)

const getUserByIDQuery = `SELECT id, name, birthday FROM users WHERE id = $1`

func (q *Queries) FindUserByID(ctx context.Context, id int) (*domain.User, error) {
	row := q.pool.QueryRow(ctx, getUserByIDQuery, id)

	user := &domain.User{}
	if err := row.Scan(&user.ID, &user.Name, &user.Birthday); err != nil {
		return nil, fmt.Errorf("can't scan user: %w", err)
	}

	return user, nil
}

const getUsersByNameQuery = `SELECT id, name, birthday FROM users WHERE name = $1`

func (q *Queries) GetUsersByName(ctx context.Context, name string) ([]domain.User, error) {
	rows, err := q.pool.Query(ctx, getUsersByNameQuery, name)
	if err != nil {
		return nil, fmt.Errorf("can't select users by name: %w", err)
	}

	defer rows.Close()

	var result []domain.User
	for rows.Next() {
		user := domain.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Birthday); err != nil {
			return nil, fmt.Errorf("can't scan user: %w", err)
		}

		result = append(result, user)
	}

	return result, nil
}

