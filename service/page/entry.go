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

	context map[string]interface{}

	tplContent string
	tpl        string
}

// SetPageType
//  @receiver page
//  @param pageType
func (page *Page) SetPageType(pageType string) {
	page.context["page_type"] = pageType
}

// SetData
//  @receiver page
//  @param data
func (page *Page) SetData(data interface{}) {
	bytes, _ := json.Marshal(page.Data)
	page.context["data"] = data
	page.context["raws"] = string(bytes)
}

// SetTitle
//  @receiver page
//  @param title
func (page *Page) SetTPLFile(fileName string) {
	page.tpl = fileName
}

// SetTitle
//  @receiver page
//  @param title
func (page *Page) SetTitle(title string) {
	page.context["title"] = title
}

// DefaultPage
//  @param ctx
//  @return *Page
func DefaultPage(ctx *gin.Context) *Page {
	return &Page{
		GinCtx:       ctx,
		context:      map[string]interface{}{},
		skeletonFile: container.GetSkeletionTemplatePath(),
	}
}

func NewPage(ctx *gin.Context, title string, data interface{}) *Page {
	return &Page{
		GinCtx:       ctx,
		Title:        title,
		Data:         data,
		context:      map[string]interface{}{},
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

func (page *Page) Render() string {

	if err := page.parse(page.tpl); err != nil {
		return err.Error()
	}
	tplParser, _ := pongo2.FromString(page.tplContent)
	html, _ := tplParser.Execute(pongo2.Context(page.context))

	return html

}
