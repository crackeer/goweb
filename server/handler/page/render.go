package page

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	ginhelper "github.com/crackeer/gopkg/gin"
	"github.com/crackeer/gopkg/util"
	"github.com/crackeer/goweb/container"
	"github.com/crackeer/goweb/define"
	"github.com/gin-gonic/gin"
	rollRender "github.com/unrolled/render"
)

// Render
//  @param ctx
func Render(ctx *gin.Context) {
	appConfig := container.GetAppConfig()
	fmt.Println(ctx.Request.URL.Path, appConfig.PublicFileExtension)
	isStaic := false
	contentType := "text/plain"
	for ext, cType := range appConfig.PublicFileExtension {
		if strings.HasSuffix(ctx.Request.URL.Path, "."+ext) {
			isStaic = true
			contentType = cType
		}
	}

	path := strings.TrimRight(ctx.Request.URL.Path, "/")

	if isStaic {
		filePath := mergePath(appConfig.ResourceDir, path)
		bytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			ctx.Data(http.StatusOK, "text/html", []byte(err.Error()))
		} else {
			ctx.Data(http.StatusOK, contentType, bytes)
		}
		return
	}

	if html, err := renderByConfig(ctx, path); err == nil {
		ctx.Data(http.StatusOK, "text/html", []byte(html))
		return
	}

	if html, err := renderByPath(ctx, path); err == nil {
		ctx.Data(http.StatusOK, "text/html", []byte(html))
		return
	}
	ctx.Data(http.StatusOK, "text/html", []byte("page not found"))

}

func renderByPath(ctx *gin.Context, path string) (string, error) {
	appConfig := container.GetAppConfig()
	return renderCake(appConfig.ResourceDir, appConfig.DefaultFrameFile, path, nil)
}

// Render
//  @param ctx
func renderByConfig(ctx *gin.Context, path string) (string, error) {

	appConfig := container.GetAppConfig()

	pageConfigPath := mergePath(appConfig.PageConfDir, path)
	pageConfig, err := loadPageConfig(pageConfigPath)

	if err != nil {
		return "", err
	}

	layout := appConfig.DefaultFrameFile
	if len(pageConfig.FrameFile) > 0 {
		layout = pageConfig.FrameFile
	}

	params := ginhelper.AllParams(ctx)
	jsData := map[string]interface{}{
		"query":     params,
		"extension": pageConfig.Extension,
	}
	bytes1, err := json.Marshal(pageConfig.DataAPIMesh)
	fmt.Println(string(bytes1))
	apiData, err := requestDataAPI(pageConfig, params)
	if err != nil {
		return err.Error(), nil
	}
	jsData["api_data"] = apiData

	//dataString, _ := util.MarshalEscapeHtml(jsData)

	pageData := map[string]interface{}{
		"title": pageConfig.Title,
		"data":  jsData,
	}

	value, err := renderCake(appConfig.ResourceDir, layout, pageConfig.ContentFile, pageData)
	if err != nil {
		return err.Error(), nil
	}
	return value, nil

	//return render.RenderHTML(framePath, mergePath(appConfig.ResourceDir, pageConfig.ContentFile), opt)
}

func mergePath(prefix string, addFile string) string {
	return strings.TrimRight(prefix, "/") + "/" + strings.Trim(addFile, "/")
}

func renderCake(dir, layout, name string, data interface{}) (string, error) {
	fmt.Println(dir, layout, name)
	object := rollRender.New(rollRender.Options{
		Directory:  dir,                           // Specify what path to load the templates from.
		FileSystem: &rollRender.LocalFileSystem{}, // Specify filesystem from where files are loaded.
		Layout:     layout,                        // Specify a layout template. Layouts can call {{ yield }} to render the current template or {{ partial "css" }} to render a partial from the current template.
		Extensions: []string{".html"},             // Specify extensions to load for templates.
		Delims: rollRender.Delims{
			Left:  "{[{",
			Right: "}]}",
		},
		Asset:      nil,
		AssetNames: nil,
	})

	buffer := bytes.NewBufferString("")
	err := object.HTML(buffer, 200, name, data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func requestDataAPI(pageConfig *define.PageConfig, params map[string]interface{}) (interface{}, error) {
	for k, v := range pageConfig.DefaultParams {
		params[k] = v
	}

	if len(pageConfig.DataAPIMesh) > 0 {
		response, _, err := container.APIRequestClient.Mesh(pageConfig.DataAPIMesh, params, map[string]string{}, "")
		if err != nil {
			return nil, fmt.Errorf("api mesh request error: %s", err.Error())
		}
		groupAPIData := map[string]interface{}{}
		for name, item := range response {
			var apiData interface{}
			if err := util.Unmarshal(item.Data, &apiData); err != nil {
				groupAPIData[name] = item.Data
			} else {
				groupAPIData[name] = apiData
			}
		}
		return groupAPIData, nil
	}

	if len(pageConfig.DataAPI) > 0 {
		response, err := container.APIRequestClient.Request(pageConfig.DataAPI, params, map[string]string{}, "")
		if err != nil {
			return nil, fmt.Errorf("api request error: %s", err.Error())
		}
		var apiData interface{}
		if err := util.Unmarshal(response.Data, &apiData); err != nil {
			apiData = response.Data
		}
		return apiData, nil
	}
	return nil, nil
}
