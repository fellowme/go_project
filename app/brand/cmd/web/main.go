package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	"go_project/app/brand/brand_model"
	"go_project/app/brand/brand_router"
)

/*
	主程序
*/
func main() {
	modelList := []interface{}{
		&brand_model.Brand{},
	}
	gin_app.CreateAppServer("/app/brand/brand_config/", "go_brand", brand_router.InitRouter, modelList)
}
