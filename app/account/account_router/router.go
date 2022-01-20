package account_router

import (
	gin_jwt "github.com/fellowme/gin_common_library/jwt"
	"github.com/gin-gonic/gin"
	"go_project/app/account/account_control"
)

func InitRouter(group *gin.RouterGroup) {
	control := account_control.GetAccountControl()
	accountRouter := group.Group("/account")
	{
		accountRouter.GET("", gin_jwt.JwtAuth(), control.GetAccountList)
		accountRouter.POST("/send_code", control.PostSendCode)
		accountRouter.POST("/verification_code", control.PostVerificationCode)
		accountRouter.POST("/login_out", control.PostLoginOut)
		accountRouter.POST("/login", control.PostLogin)
	}
}
