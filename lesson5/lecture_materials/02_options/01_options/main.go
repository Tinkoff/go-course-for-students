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

type ServerOptions struct {
	host          string
	port          int
	timeout       time.Duration
	anotherOption string
}

func NewServer(options *ServerOptions) *Server {
	if options.host == "" {
		options.host = "localhost"
	}
	if options.port == 0 {
		options.port = 8080
	}
	if options.timeout == 0 {
		options.timeout = time.Second
	}
	return &Server{
		host:          options.host,
		port:          options.port,
		timeout:       options.timeout,
		anotherOption: options.anotherOption,
	}
}

func main() {
	_ = NewServer(&ServerOptions{host: "0.0.0.0", port: 80})
}
