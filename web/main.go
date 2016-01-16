package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/azer/crud"
	"github.com/eknkc/amber"
	_ "github.com/lib/pq"
)

var DB *crud.DB
var compiler *amber.Compiler

func init() {
	DB, _ := crud.Connect("postgres", os.Getenv("DATABASE_URL"))
	DB.Ping()
	compiler = amber.New()
}

func renderError(w http.ResponseWriter, err error) {
	io.WriteString(w, fmt.Sprintf("%v", err))
}

func renderPage(w http.ResponseWriter, renderedTemplate string) error {
	data := struct{ Body template.HTML }{template.HTML(renderedTemplate)}
	t, err := template.ParseFiles("templates/base.html")
	t.Execute(w, data)
	return err
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := compiler.ParseFile("templates/index.amber")
	if err != nil {
		renderError(w, err)
	}
	tpl, err := compiler.Compile()
	if err != nil {
		renderError(w, err)
	}
	templateBuffer := new(bytes.Buffer)
	tpl.Execute(templateBuffer, nil)
	err = renderPage(w, templateBuffer.String())
	if err != nil {
		renderError(w, err)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
