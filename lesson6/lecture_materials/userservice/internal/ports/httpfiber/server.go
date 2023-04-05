package httpfiber

import (
	"userservice/internal/app"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	port     string
	fiberApp *fiber.App
}

func NewHTTPServer(port string, a app.App) Server {
	s := Server{port: port, fiberApp: fiber.New()}

	api := s.fiberApp.Group("/api/v1")
	AppRouter(api, a)

	return s
}

func (s *Server) Listen() error {
	return s.fiberApp.Listen(s.port)
}
