package brand_router

import (
	"github.com/gin-gonic/gin"
	"go_project/app/brand/brand_control"
)

func InitRouter(api *gin.RouterGroup) {
	brandApi := api.Group("/brand")
	control := brand_control.GetBrandControl()
	{
		brandApi.GET("", control.GetBrandList)
		brandApi.POST("", control.PostBrand)
		brandApi.PATCH("/:id", control.PatchBrand)
		brandApi.DELETE("/:id", control.DeleteBrand)
		brandApi.GET("/get_by_ids", control.GetBrandByIds)
		brandApi.POST("/delete_by_ids", control.DeleteBrandByIds)
	}
}
