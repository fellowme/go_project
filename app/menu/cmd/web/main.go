package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	"go_project/app/menu/menu_model"
	"go_project/app/menu/menu_router"
)

/*
	主程序
*/
func main() {
	modelList := []interface{}{
		&menu_model.Menu{},
	}
	gin_app.CreateAppServer("/app/menu/menu_config/", "go_menu", menu_router.InitRouter, modelList)

}
