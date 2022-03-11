package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	"go_project/app/category/category_model"
	"go_project/app/category/category_router"
)

/*
	主程序
*/
func main() {
	modelList := []interface{}{
		&category_model.Category{},
	}
	gin_app.CreateAppServer("/app/category/category_config/", "go_category", category_router.InitRouter, modelList)
}
