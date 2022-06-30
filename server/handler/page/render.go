package page

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/crackeer/gopkg/config"
	"github.com/crackeer/gopkg/ginhelper"
	"github.com/crackeer/gopkg/render"
	"github.com/crackeer/gopkg/util"
	"github.com/crackeer/goweb/container"
	"github.com/gin-gonic/gin"
)

func Render(ctx *gin.Context) {

	if html, err := renderByConfig(ctx); err == nil {
		ctx.Data(http.StatusOK, "text/html", []byte(html))
		return
	} else {
		fmt.Println(err.Error())
	}

	if html, err := renderByPath(ctx); err == nil {
		ctx.Data(http.StatusOK, "text/html", []byte(html))
		return
	}

	ctx.Data(http.StatusOK, "text/html", []byte("page not found"))

}

func renderByPath(ctx *gin.Context) (string, error) {
	appConfig := container.GetAppConfig()
	pageFilePath := mergePath(appConfig.Resource.PageDir, strings.TrimRight(ctx.Request.URL.Path, "/")+".html")
	return render.RenderHTML(appConfig.Resource.DefaultFrameFile, pageFilePath, nil)
}

// Render
//  @param ctx
func renderByConfig(ctx *gin.Context) (string, error) {
	pagePath := ctx.Request.URL.Path
	parts := strings.Split(strings.Trim(pagePath, "/"), "/")
	appConfig := container.GetAppConfig()

	pageConfigPath := mergePath(appConfig.Resource.PageConfDir, parts[0])
	fmt.Println(appConfig.Resource, pageConfigPath)
	groupPageConfig, err := config.LoadYamlGroupPageConfig(pageConfigPath)

	if err != nil {
		return "", err
	}
	pageID := strings.Join(parts[1:], "/")
	pageConfig, exists := groupPageConfig.List[pageID]
	if !exists {
		return "", errors.New("page not found")
	}

	framePath := ""
	if len(pageConfig.FrameFile) > 0 {
		framePath = mergePath(appConfig.Resource.PageDir, pageConfig.FrameFile)
	} else if len(appConfig.Resource.DefaultFrameFile) > 0 {
		framePath = appConfig.Resource.DefaultFrameFile
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

	return render.RenderHTML(framePath, mergePath(appConfig.Resource.PageDir, pageConfig.ContentFile), opt)
}

func mergePath(prefix string, addFile string) string {
	return strings.TrimRight(prefix, "/") + "/" + strings.Trim(addFile, "/")
}
