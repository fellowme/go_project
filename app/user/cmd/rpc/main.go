package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	"go_project/app/user/user_rpc"
	service "go_project/rpc_service"
)

func main() {
	gRpcService := gin_app.InitRpcServer("/app/user/user_config/", "go_user_rpc")
	userService := user_rpc.GetUserRpcService()
	service.RegisterUserServiceServer(gRpcService, &userService)
	gin_app.CreateRpcServer(gRpcService)

}
