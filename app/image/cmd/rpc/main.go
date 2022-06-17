package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	grpc_consul "github.com/fellowme/gin_common_library/consul"
	"go_project/app/image/image_rpc"
	service "go_project/rpc_service"
	"time"
)

func main() {
	gRpcService := gin_app.InitRpcServer("/app/image/image_config/", "go_image_rpc")
	imageService := image_rpc.GetImageRpcService()
	service.RegisterImageServiceServer(gRpcService, imageService)
	gin_app.CreateRpcServer(gRpcService, grpc_consul.ServiceConsul{
		Id:       fmt.Sprintf("%s-version-%d", "go_image_rpc", time.Now().Unix()),
		Name:     "go_image_rpc",
		Port:     18088,
		Address:  "192.168.1.224",
		IsSecure: false,
	})
}
