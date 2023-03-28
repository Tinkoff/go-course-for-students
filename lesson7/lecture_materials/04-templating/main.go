package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("template.tpl")
	if err != nil {
		log.Fatalf("parse template file error: %s", err)
	}

	if err = tpl.Execute(os.Stdout, map[string]string{
		"name": "Alice",
	}); err != nil {
		log.Fatalf("execute template error: %s", err)
	}
}
