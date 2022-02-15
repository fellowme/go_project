package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	"go_project/app/image/image_rpc"
	service "go_project/rpc_service"
)

func main() {
	gRpcService := gin_app.InitRpcServer("/app/image/image_config/", "go_image_rpc")
	imageService := image_rpc.GetImageRpcService()
	service.RegisterImageServiceServer(gRpcService, imageService)
	gin_app.CreateRpcServer(gRpcService)
}
