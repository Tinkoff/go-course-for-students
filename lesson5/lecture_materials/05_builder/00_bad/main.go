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

func NewServer(host string, port int, timeout time.Duration, anotherOption string) *Server {
	return &Server{
		host:          host,
		port:          port,
		timeout:       timeout,
		anotherOption: anotherOption,
	}
}

func main() {
	_ = NewServer("", 0, 2*time.Second, "")
}
