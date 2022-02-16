package shop_router

import (
	"github.com/gin-gonic/gin"
	"go_project/app/shop/shop_control"
)

func InitRouter(api *gin.RouterGroup) {
	shopApi := api.Group("/shop")
	control := shop_control.GetShopControl()
	{
		shopApi.GET("", control.GetShopList)
		shopApi.POST("", control.PostShop)
		shopApi.PATCH("/:id", control.PatchShop)
		shopApi.DELETE("/:id", control.DeleteShop)
		shopApi.GET("/get_by_ids", control.GetShopByIds)
		shopApi.POST("/delete_by_ids", control.DeleteShopByIds)
	}
}
