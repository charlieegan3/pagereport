package common

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func RenderError(w http.ResponseWriter, err error) {
	io.WriteString(w, fmt.Sprintf("%v", err))
}

func RenderPage(w http.ResponseWriter, renderedTemplate string) error {
	data := struct{ Body template.HTML }{template.HTML(renderedTemplate)}
	t, err := template.ParseFiles("templates/base.html")
	t.Execute(w, data)
	return err
}
