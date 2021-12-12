package page

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/flosch/pongo2/v4"
)

// Page
type Page struct {
	Title string
	Data  interface{}

	skeletonFile string

	tplContent string
}

func NewPage(title string, data interface{}, skeletonFile string) *Page {
	return &Page{
		Title:        title,
		Data:         data,
		skeletonFile: skeletonFile,
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
	}

	if err := page.parse(tpl); err != nil {
		return err.Error()
	}
	tplParser, _ := pongo2.FromString(page.tplContent)
	html, _ := tplParser.Execute(pongo2.Context(context))

	return html

}
