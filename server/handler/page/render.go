package page

import (
	"net/http"
	"strings"

	"github.com/crackeer/gopkg/config"
	"github.com/crackeer/gopkg/render"
	"github.com/crackeer/goweb/container"
	"github.com/gin-gonic/gin"
)

// Render
//  @param ctx
func Render(ctx *gin.Context) {
	pagePath := ctx.Request.URL.Path

	parts := strings.Split(strings.Trim(pagePath, "/"), "/")

	configPrefix := "config/page/"

	resourcePath := "resource"
	groupPageConfig, err := config.LoadYamlGroupPageConfig(configPrefix + parts[0])

	if err != nil {
		ctx.Data(http.StatusOK, "text/html", []byte(err.Error()))
		return
	}
	pageID := strings.Join(parts[1:], "/")
	pageConfig, exists := groupPageConfig.List[pageID]
	if !exists {
		ctx.Data(http.StatusOK, "text/html", []byte("page not found"))
		return
	}

	framePath := ""
	if len(pageConfig.FrameFile) > 0 {
		framePath = resourcePath + "/" + pageConfig.FrameFile
	}

	opt := render.DefaultOption()
	if len(pageConfig.DataAPI) > 0 {
		response, _ := container.APIRequestClient.Request(pageConfig.DataAPI, map[string]interface{}{}, map[string]string{}, "")
		opt.InjectData = map[string]interface{}{
			"api_data": response.Data,
		}
	}

	raws, _ := render.RenderHTML(framePath, resourcePath+"/"+pageConfig.ContentFile, opt)
	ctx.Data(http.StatusOK, "text/html", []byte(raws))
}
