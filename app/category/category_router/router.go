package category_router

import (
	"github.com/gin-gonic/gin"
	"go_project/app/category/category_control"
)

func InitRouter(router *gin.RouterGroup) {
	categoryApi := router.Group("/category")
	control := category_control.GetCategoryControl()
	{
		categoryApi.GET("", control.GetCategoryList)
		categoryApi.POST("", control.CreateCategory)
		// categoryApi.GET("/rebuild", control.RebuildCategory)
		categoryApi.PATCH("/:id", control.UpdateCategory)
		categoryApi.DELETE("/:id", control.DeleteCategory)
	}
}
