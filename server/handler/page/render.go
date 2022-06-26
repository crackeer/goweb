package page

import (
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
	if len(pageConfig.DataAPI) > 0 {
		params := ginhelper.AllParams(ctx)
		response, _ := container.APIRequestClient.Request(pageConfig.DataAPI, params, map[string]string{}, "")
		var apiData interface{}
		if err := util.Unmarshal(response.Data, &apiData); err != nil {
			apiData = response.Data
		}
		opt.InjectData = map[string]interface{}{
			"api_data": apiData,
		}
	}

	return render.RenderHTML(framePath, mergePath(appConfig.Resource.PageDir, pageConfig.ContentFile), opt)
}

func mergePath(prefix string, addFile string) string {
	return strings.TrimRight(prefix, "/") + "/" + strings.Trim(addFile, "/")
}
