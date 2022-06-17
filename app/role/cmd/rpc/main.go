package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	grpc_consul "github.com/fellowme/gin_common_library/consul"
	"go_project/app/role/role_rpc"
	service "go_project/rpc_service"
	"time"
)

func main() {
	gRpcService := gin_app.InitRpcServer("/app/role/role_config/", "go_role_rpc")
	roleService := role_rpc.GetRoleRpcService()
	service.RegisterUserRoleServiceServer(gRpcService, &roleService)
	gin_app.CreateRpcServer(gRpcService, grpc_consul.ServiceConsul{
		Id:       fmt.Sprintf("%s-version-%d", "go_role_rpc", time.Now().Unix()),
		Name:     "go_role_rpc",
		Port:     18084,
		Address:  "192.168.1.224",
		IsSecure: false,
	})
}
