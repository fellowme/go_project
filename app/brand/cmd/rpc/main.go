package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	"go_project/app/brand/brand_rpc"
	service "go_project/rpc_service"
)

func main() {
	gRpcService := gin_app.InitRpcServer("/app/brand/brand_config/", "go_brand_rpc")
	brandService := brand_rpc.GetBrandRpcService()
	service.RegisterBrandServiceServer(gRpcService, brandService)
	gin_app.CreateRpcServer(gRpcService)
}
