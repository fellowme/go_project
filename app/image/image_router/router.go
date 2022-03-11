package image_router

import (
	gin_util "github.com/fellowme/gin_common_library/util"
	"github.com/gin-gonic/gin"
	"go_project/app/image/image_control"
	"path"
)

func InitRouter(router *gin.RouterGroup) {
	imageApi := router.Group("/image")
	control := image_control.GetImageControl()
	{
		imageApi.GET("", control.GetImageList)
		imageApi.POST("", control.CreateImage)
		imageApi.GET("/:id", control.GetImage)
		//imageApi.PATCH("/:id", control.UpdateImage)
		imageApi.DELETE("/:id", control.DeleteImage)
	}
	staticApi := imageApi.Group("/upload")
	{
		staticApi.Static("/", path.Join(gin_util.GetPath(), "upload"))
	}

}
