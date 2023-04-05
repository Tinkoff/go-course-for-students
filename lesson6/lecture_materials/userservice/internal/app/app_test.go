package app_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"userservice/internal/app"
	"userservice/internal/user"
)

type FakeRepository struct{}

func (f *FakeRepository) AddUser(_ context.Context, user user.User) (int64, error) {
	return 100, nil
}

func (f *FakeRepository) GetUserByID(_ context.Context, user_id int64) (*user.User, error) {
	return &user.User{ID: 100, Name: "test_user"}, nil
}

func TestApp(t *testing.T) {
	ctx := context.TODO()

	app := app.NewApp(&FakeRepository{})

	t.Run("create_user", func(t *testing.T) {
		u, err := app.CreateUser(ctx, "test_user")

		assert.Equal(t, err, nil)
		assert.Equal(t, u.ID, 100)
		assert.Equal(t, u.Name, "test_user")
	})
}
