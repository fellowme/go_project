package role_router

import (
	"github.com/gin-gonic/gin"
	"go_project/app/role/role_control"
)

func InitRouter(router *gin.RouterGroup) {
	roleApi := router.Group("/role")
	control := role_control.GetRoleControl()
	{
		roleApi.GET("", control.GetRoleList)
		roleApi.POST("", control.PostRole)
		roleApi.PATCH("/:id", control.PatchRole)
		roleApi.DELETE("/:id", control.DeleteRole)
		roleApi.GET("/user", control.GetRoleUser)
		roleApi.POST("/user", control.PostRoleUser)
		roleApi.DELETE("/user/:id", control.DeleteRoleUser)
		roleApi.GET("/menu", control.GetRoleMenu)
		roleApi.POST("/menu", control.PostRoleMenu)
		roleApi.DELETE("/menu/:id", control.DeleteRoleMenu)
	}
}
