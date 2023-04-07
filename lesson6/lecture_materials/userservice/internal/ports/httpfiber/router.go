package httpfiber

import (
	"userservice/internal/app"

	"github.com/gofiber/fiber/v2"
)

func AppRouter(r fiber.Router, a app.App) {
	r.Get("/users/:user_id", getUser(a))
	r.Post("/users", createUser(a))
}
