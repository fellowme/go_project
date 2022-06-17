package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	grpc_consul "github.com/fellowme/gin_common_library/consul"
	"go_project/app/menu/menu_rpc"
	service "go_project/rpc_service"
	"time"
)

func main() {
	gRpcService := gin_app.InitRpcServer("/app/menu/menu_config/", "go_menu_rpc")
	menuService := menu_rpc.GetMenuRpcService()
	service.RegisterMenuServiceServer(gRpcService, &menuService)
	gin_app.CreateRpcServer(gRpcService, grpc_consul.ServiceConsul{
		Id:       fmt.Sprintf("%s-version-%d", "go_menu_rpc", time.Now().Unix()),
		Name:     "go_menu_rpc",
		Port:     18086,
		Address:  "192.168.1.224",
		IsSecure: false,
	})
}
