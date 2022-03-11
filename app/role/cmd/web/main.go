package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	"go_project/app/role/role_model"
	"go_project/app/role/role_router"
)

/*
	主程序
*/
func main() {
	modelList := []interface{}{
		&role_model.Role{}, &role_model.RoleMenu{}, role_model.RoleUser{},
	}
	gin_app.CreateAppServer("/app/role/role_config/", "go_role", role_router.InitRouter, modelList)
}
