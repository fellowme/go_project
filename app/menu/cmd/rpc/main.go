package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	"go_project/app/menu/menu_rpc"
	service "go_project/rpc_service"
)

func main() {
	gRpcService := gin_app.InitRpcServer("/app/menu/menu_config/", "go_menu_rpc")
	menuService := menu_rpc.GetMenuRpcService()
	service.RegisterMenuServiceServer(gRpcService, &menuService)
	gin_app.CreateRpcServer(gRpcService)
}
