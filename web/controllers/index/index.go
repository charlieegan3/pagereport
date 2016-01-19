package index

import (
	"bytes"
	"net/http"

	"github.com/charlieegan3/pagereport/web/controllers/common"
	"github.com/eknkc/amber"
)

var compiler *amber.Compiler

func init() {
	compiler = amber.New()
}

func View(w http.ResponseWriter, r *http.Request) {
	err := compiler.ParseFile("templates/index.amber")
	if err != nil {
		common.RenderError(w, err)
	}
	tpl, err := compiler.Compile()
	if err != nil {
		common.RenderError(w, err)
	}
	templateBuffer := new(bytes.Buffer)
	tpl.Execute(templateBuffer, nil)
	err = common.RenderPage(w, templateBuffer.String())
	if err != nil {
		common.RenderError(w, err)
	}
}
