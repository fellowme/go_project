package stock_router

import (
	"github.com/gin-gonic/gin"
	"go_project/app/stock/stock_control"
)

func InitRouter(api *gin.RouterGroup) {
	control := stock_control.GetStockControl()
	stockApi := api.Group("/stock")
	{
		stockApi.GET("", control.GetStockList)
		stockApi.POST("", control.PostStock)
		stockApi.PATCH("/:id", control.PatchStock)
		stockApi.DELETE("/:id", control.DeleteStock)
		stockApi.POST("/delete_by_id", control.DeleteStockById)
		stockApi.GET("/get_by_id", control.GetStockById)
		stockApi.POST("/to_redis", control.PostToRedis)
	}
}
