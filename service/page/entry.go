package page

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/crackeer/goweb/container"
	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
)

// Page
type Page struct {
	GinCtx *gin.Context
	Title  string
	Data   interface{}

	skeletonFile string

	tplContent string
}

func NewPage(ctx *gin.Context, title string, data interface{}) *Page {
	return &Page{
		GinCtx:       ctx,
		Title:        title,
		Data:         data,
		skeletonFile: container.GetSkeletionTemplatePath(),
	}
}

func (page *Page) parse(tplFile string) error {
	bytes1, err := ioutil.ReadFile(tplFile)
	if err != nil {
		return err
	}
	page.tplContent = string(bytes1)
	if bytes, err := ioutil.ReadFile(page.skeletonFile); err == nil {
		page.tplContent = strings.ReplaceAll(string(bytes), "{{BODY}}", string(bytes1))
	}

	return nil

}

func (page *Page) Render(tpl string) string {

	bytes, _ := json.Marshal(page.Data)

	context := map[string]interface{}{
		"title": page.Title,
		"data":  page.Data,
		"raw":   string(bytes),
		"path":  page.GinCtx.Request.URL.String(),
	}

	if err := page.parse(tpl); err != nil {
		return err.Error()
	}
	tplParser, _ := pongo2.FromString(page.tplContent)
	html, _ := tplParser.Execute(pongo2.Context(context))

	return html

}
