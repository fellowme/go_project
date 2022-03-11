package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	"go_project/app/product/product_model"
	"go_project/app/product/product_router"
)

/*
	主程序
*/
func main() {
	modelList := []interface{}{
		&product_model.Product{}, &product_model.ProductMain{},
		&product_model.ProductMain{}, &product_model.ProductImage{}, &product_model.Stock{},
	}
	gin_app.CreateAppServer("/app/product/product_config/", "go_product", product_router.InitRouter, modelList)
}
