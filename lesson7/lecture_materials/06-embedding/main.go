package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed assets/*
var assets embed.FS

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", handler())

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func handler() http.Handler {

	fsys := fs.FS(assets)
	html, _ := fs.Sub(fsys, "assets")

	return http.FileServer(http.FS(html))
}
