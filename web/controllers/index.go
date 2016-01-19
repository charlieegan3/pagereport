package index

import (
	"bytes"
	"net/http"

	"github.com/charlieegan3/pagereport/web/controllers/common"
	"github.com/eknkc/amber"
)

func View(w http.ResponseWriter, r *http.Request, c *amber.Compiler) {
	err := c.ParseFile("templates/index.amber")
	if err != nil {
		common.RenderError(w, err)
	}
	tpl, err := c.Compile()
	if err != nil {
		common.RenderError(w, err)
	}
	templateBuffer := new(bytes.Buffer)
	tpl.Execute(templateBuffer, nil)
	err = common.RenderPage(w, templateBuffer.String())
	if err != nil {
		common.renderError(w, err)
	}
}
