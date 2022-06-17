package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	grpc_consul "github.com/fellowme/gin_common_library/consul"
	"go_project/app/user/user_rpc"
	service "go_project/rpc_service"
	"time"
)

func main() {
	gRpcService := gin_app.InitRpcServer("/app/user/user_config/", "go_user_rpc")
	userService := user_rpc.GetUserRpcService()
	service.RegisterUserServiceServer(gRpcService, &userService)
	gin_app.CreateRpcServer(gRpcService, grpc_consul.ServiceConsul{
		Id:       fmt.Sprintf("%s-version-%d", "go_user_rpc", time.Now().Unix()),
		Name:     "go_user_rpc",
		Port:     18082,
		Address:  "192.168.1.224",
		IsSecure: false,
	})

}
