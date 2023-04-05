package httpfiber

import (
	"errors"
	"userservice/internal/user"

	"github.com/gofiber/fiber/v2"
)

type userNameRequest struct {
	Name string `json:"name"`
}

func (u userNameRequest) validate() error {
	if u.Name == "" {
		return errors.New("field name is required")
	}

	return nil
}

type userResponse struct {
	ID       int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

func UserSuccessResponse(u *user.User) *fiber.Map {
	data := userResponse{
		ID:       u.ID,
		UserName: u.Name,
	}

	return &fiber.Map{
		"data":  data,
		"error": nil,
	}
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"data":  nil,
		"error": err.Error(),
	}

}
