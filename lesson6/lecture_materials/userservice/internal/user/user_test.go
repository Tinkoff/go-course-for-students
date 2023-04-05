package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"userservice/internal/user"
)

func TestApp(t *testing.T) {
	expected := user.User{ID: 1, Name: "test"}

	t.Run("new_user", func(t *testing.T) {
		u := user.New(1, "test")

		assert.Equal(t, u.ID, expected.ID)
		assert.Equal(t, u.Name, expected.Name)
	})
}
