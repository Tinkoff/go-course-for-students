package app

import (
	"context"

	"userservice/internal/user"
)

type App struct {
	repository user.Repository
}

func NewApp(repo user.Repository) App {
	return App{repository: repo}
}

func (a *App) CreateUser(ctx context.Context, name string) (*user.User, error) {
	u := user.User{Name: name}

	id, err := a.repository.AddUser(ctx, u)
	if err != nil {
		return nil, err
	}

	u.ID = id

	return &u, nil
}

func (a *App) GetUser(ctx context.Context, id int64) (*user.User, error) {
	u, err := a.repository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}
