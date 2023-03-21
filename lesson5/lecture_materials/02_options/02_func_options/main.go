package main

import (
	"time"
)

type Server struct {
	host          string
	port          int
	timeout       time.Duration
	anotherOption string
}

type ServerOption func(*Server)

func NewServer(options ...ServerOption) *Server {
	server := &Server{
		host:    "localhost",
		port:    8080,
		timeout: time.Second,
	}

	for _, option := range options {
		option(server)
	}

	return server
}

func WithHost(host string) ServerOption {
	return func(server *Server) {
		server.host = host
	}
}

func WithPort(port int) ServerOption {
	return func(server *Server) {
		server.port = port
	}
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(server *Server) {
		server.timeout = timeout
	}
}

func WithAnotherOptions(anotherOption string) ServerOption {
	return func(server *Server) {
		server.anotherOption = anotherOption
	}
}

func main() {
	_ = NewServer(
		WithPort(80),
		WithTimeout(10*time.Second),
	)
}
