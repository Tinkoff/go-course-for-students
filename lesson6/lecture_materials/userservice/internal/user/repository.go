package user

import "context"

type Repository interface {
	GetUserByID(ctx context.Context, user_id int64) (*User, error)
	AddUser(ctx context.Context, user User) (int64, error)
}
