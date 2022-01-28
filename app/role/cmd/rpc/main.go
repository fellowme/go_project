package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	"go_project/app/role/role_rpc"
	service "go_project/rpc_service"
)

func main() {
	gRpcService := gin_app.InitRpcServer("/app/role/role_config/", "go_role_rpc")
	roleService := role_rpc.GetRoleRpcService()
	service.RegisterUserRoleServiceServer(gRpcService, &roleService)
	gin_app.CreateRpcServer(gRpcService)
}
