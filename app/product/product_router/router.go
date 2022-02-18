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
		productMainApi.GET("/:id", control.GetProductMain)
		productMainApi.PATCH("/:id", control.PatchProductMain)
		productMainApi.DELETE("/:id", control.DeleteProductMain)
	}
	productApi := api.Group("/product")
	{
		productApi.GET("", control.GetProductList)
		productApi.POST("", control.PostProduct)
		productApi.GET("/:id", control.GetProduct)
		productApi.PATCH("/:id", control.PatchProduct)
		productApi.DELETE("/:id", control.DeleteProduct)
		productApi.GET("/get_by_product_main_ids", control.GetProductByProductMainIds)
	}
	stockApi := api.Group("/stock")
	{
		stockApi.GET("", control.GetProductStockList)
		stockApi.POST("", control.PostProductStock)
		stockApi.PATCH("/:id", control.PatchProductStock)
		stockApi.DELETE("/:id", control.DeleteProductStock)
		stockApi.POST("/delete_by_id", control.DeleteProductStockById)
		stockApi.GET("/get_by_id", control.GetProductStockById)
	}

}
