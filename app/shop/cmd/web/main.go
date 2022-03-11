package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	"go_project/app/shop/shop_model"
	"go_project/app/shop/shop_router"
)

/*
	主程序
*/
func main() {
	modelList := []interface{}{
		&shop_model.Shop{},
	}
	gin_app.CreateAppServer("/app/shop/shop_config/", "go_shop", shop_router.InitRouter, modelList)
}
