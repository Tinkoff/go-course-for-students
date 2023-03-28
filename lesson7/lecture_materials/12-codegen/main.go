package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"

	_ "embed"
)

//go:embed template.tpl
var templateContent string

type Data struct {
	Package   string
	Fields    []string
	TableName string
	Struct    string
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fileName := filepath.Join(cwd, "model.go")
	fset := token.NewFileSet() // a set of sources
	node, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("error parse file: %s", err)
	}

	out, err := os.Create(filepath.Join(cwd, "model_out.go"))
	if err != nil {
		log.Fatalf("error create output file: %s", err)
	}
	defer out.Close()

	tpl, err := template.New("template.tpl").Funcs(template.FuncMap{
		"join": strings.Join,
	}).Parse(templateContent)
	if err != nil {
		log.Fatalf("parse template file error: %s", err)
	}

	data := &Data{
		Package: node.Name.Name,
	}

	// declaration
	decl := node.Decls[0].(*ast.GenDecl)

	// base struct
	model := decl.Specs[0].(*ast.TypeSpec)

	// resolve name of base strict
	data.Struct = model.Name.String()

	// resolve table name
	for _, comment := range decl.Doc.List {
		if strings.HasPrefix(comment.Text, "// db:table=") {
			data.TableName = strings.Replace(comment.Text, "// db:table=", "", 1)
			break
		}
	}

	// resolve fields
	structData := model.Type.(*ast.StructType)

	for _, field := range structData.Fields.List {
		if field.Tag == nil {
			continue
		}

		tag := reflect.StructTag(field.Tag.Value[1 : len(field.Tag.Value)-1])

		val, ok := tag.Lookup("db")
		if !ok {
			continue
		}

		data.Fields = append(data.Fields, val)
	}

	if err = tpl.ExecuteTemplate(out, "template.tpl", data); err != nil {
		log.Fatalf("execute template error: %s", err)
	}
}
