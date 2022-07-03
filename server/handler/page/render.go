package page

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	ginhelper "github.com/crackeer/gopkg/gin"
	"github.com/crackeer/gopkg/render"
	"github.com/crackeer/gopkg/util"
	"github.com/crackeer/goweb/container"
	"github.com/gin-gonic/gin"
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
	pageFilePath := mergePath(appConfig.ResourceDir, path+".html")
	framePath := mergePath(appConfig.ResourceDir, appConfig.DefaultFrameFile)
	return render.RenderHTML(framePath, pageFilePath, nil)
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

	framePath := ""
	if len(pageConfig.FrameFile) > 0 {
		framePath = mergePath(appConfig.ResourceDir, pageConfig.FrameFile)
	} else if len(appConfig.DefaultFrameFile) > 0 {
		framePath = mergePath(appConfig.ResourceDir, appConfig.DefaultFrameFile)
	}

	opt := render.DefaultOption()
	params := ginhelper.AllParams(ctx)
	jsData := map[string]interface{}{
		"query":     params,
		"extension": pageConfig.Extension,
	}
	bytes1, err := json.Marshal(pageConfig.DataAPIMesh)
	fmt.Println(string(bytes1))
	for k, v := range pageConfig.DefaultParams {
		params[k] = v
	}
	if len(pageConfig.DataAPIMesh) > 0 {
		response, _, err := container.APIRequestClient.Mesh(pageConfig.DataAPIMesh, params, map[string]string{}, "")
		if err != nil {
			return fmt.Sprintf("api mesh request error: %s", err.Error()), nil
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
		jsData["api_data"] = groupAPIData
	} else if len(pageConfig.DataAPI) > 0 {
		response, err := container.APIRequestClient.Request(pageConfig.DataAPI, params, map[string]string{}, "")
		if err != nil {
			return fmt.Sprintf("api request error: %s", err.Error()), nil
		}
		var apiData interface{}
		if err := util.Unmarshal(response.Data, &apiData); err != nil {
			apiData = response.Data
		}
		jsData["api_data"] = apiData
	}
	opt.InjectData = jsData
	opt.Title = pageConfig.Title
	bytes, err := json.Marshal(opt.InjectData)
	fmt.Println(string(bytes))

	return render.RenderHTML(framePath, mergePath(appConfig.ResourceDir, pageConfig.ContentFile), opt)
}

func mergePath(prefix string, addFile string) string {
	return strings.TrimRight(prefix, "/") + "/" + strings.Trim(addFile, "/")
}
