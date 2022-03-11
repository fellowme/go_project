package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	"go_project/app/image/image_model"
	"go_project/app/image/image_router"
)

/*
	主程序
*/
func main() {
	modelList := []interface{}{
		&image_model.Image{},
	}
	gin_app.CreateAppServer("/app/image/image_config/", "go_image", image_router.InitRouter, modelList)
}
