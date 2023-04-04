package httpfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"homework6/internal/app"
)

type Server struct {
	port string
	app  *fiber.App
}

func NewHTTPServer(port string, a app.App) Server {
	s := Server{port: port, app: fiber.New()}
	api := s.app.Group("/api/v1")
	AppRouter(api, a)
	return s
}

func (s *Server) Listen() error {
	return s.app.Listen(s.port)
}

func (s *Server) Test(req *http.Request, msTimeout ...int) (*http.Response, error) {
	return s.app.Test(req, msTimeout...)
}
