package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	grpc_consul "github.com/fellowme/gin_common_library/consul"
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
	gin_app.CreateAppServer("/app/role/role_config/", "go_role", role_router.InitRouter, modelList, grpc_consul.ServiceConsul{
		Name:     "go_role",
		Port:     8084,
		Address:  "192.168.1.224",
		IsSecure: false,
	})
}
