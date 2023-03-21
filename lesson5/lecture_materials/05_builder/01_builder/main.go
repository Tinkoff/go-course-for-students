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

func NewServer() Server {
	return Server{
		host:    "localhost",
		port:    8080,
		timeout: time.Second,
	}
}

func (s Server) WithHost(host string) Server {
	s.host = host
	return s
}

func (s Server) WithPort(port int) Server {
	s.port = port
	return s
}

func (s Server) WithTimeout(timeout time.Duration) Server {
	s.timeout = timeout
	return s
}

func (s Server) WithAnotherOption(anotherOption string) Server {
	s.anotherOption = anotherOption
	return s
}

func main() {
	_ = NewServer().WithPort(3030).WithAnotherOption("test").WithTimeout(10 * time.Second)
}
