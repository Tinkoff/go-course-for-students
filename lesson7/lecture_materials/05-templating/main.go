package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	FirstName, LastName string
	Age                 int
}

type Class struct {
	Grade    int
	Teacher  Person
	Students []Person
}

func main() {

	tpl, err := template.New("template.tpl").Funcs(template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}).ParseFiles("template.tpl")
	if err != nil {
		log.Fatalf("parse template file error: %s", err)
	}

	data := loadData()

	if err = tpl.ExecuteTemplate(os.Stdout, "template.tpl", data); err != nil {
		log.Fatalf("execute template error: %s", err)
	}
}
