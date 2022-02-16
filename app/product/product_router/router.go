package product_router

import (
	"github.com/gin-gonic/gin"
	"go_project/app/product/product_control"
)

func InitRouter(api *gin.RouterGroup) {
	productMainApi := api.Group("/product_main")
	control := product_control.GetProductControl()
	{
		productMainApi.GET("", control.GetProductMainList)
		productMainApi.POST("", control.PostProductMain)
	}
}
