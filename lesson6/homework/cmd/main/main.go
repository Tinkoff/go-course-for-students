package main

import (
	"homework6/internal/adapters/adrepo"
	"homework6/internal/app"
	"homework6/internal/ports/httpfiber"
)

func main() {
	server := httpfiber.NewHTTPServer(":18080", app.NewApp(adrepo.New()))
	err := server.Listen()
	if err != nil {
		panic(err)
	}
}
