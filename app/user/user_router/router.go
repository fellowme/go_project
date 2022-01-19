package user_router

import (
	"github.com/gin-gonic/gin"
	"go_project/app/user/user_control"
)

func InitRouter(router *gin.RouterGroup) {
	userApi := router.Group("/user")
	control := user_control.GetUserControl()
	{
		userApi.GET("/by_ids", control.GetUserByIds)
		userApi.GET("", control.GetUserList)
		userApi.POST("", control.CreateUser)
		userApi.GET("/:id", control.GetUser)
		userApi.PATCH("/:id", control.UpdateUser)
		userApi.DELETE("/:id", control.DeleteUser)
		userApi.POST("/by_ids", control.DeleteUserByIdList)
	}

}
