package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	"go_project/app/user/user_model"
	"go_project/app/user/user_router"
)

/*
	主程序
*/
func main() {
	modelList := []interface{}{
		&user_model.User{},
	}
	gin_app.CreateAppServer("/app/user/user_config/", "go_user", user_router.InitRouter, modelList)

}
