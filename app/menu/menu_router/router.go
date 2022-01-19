package menu_router

import (
	"github.com/gin-gonic/gin"
	"go_project/app/menu/menu_control"
)

func InitRouter(router *gin.RouterGroup) {
	menuApi := router.Group("/menu")
	control := menu_control.GetMenuControl()
	{
		menuApi.GET("", control.GetMenuList)
		menuApi.POST("", control.PostMenu)
		menuApi.PATCH("/:id", control.PatchMenu)
		menuApi.DELETE("/:id", control.DeleteMenu)
	}
}
