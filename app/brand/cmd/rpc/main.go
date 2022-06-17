package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	grpc_consul "github.com/fellowme/gin_common_library/consul"
	"go_project/app/brand/brand_rpc"
	service "go_project/rpc_service"
	"time"
)

func main() {
	gRpcService := gin_app.InitRpcServer("/app/brand/brand_config/", "go_brand_rpc")
	brandService := brand_rpc.GetBrandRpcService()
	service.RegisterBrandServiceServer(gRpcService, brandService)
	gin_app.CreateRpcServer(gRpcService, grpc_consul.ServiceConsul{
		Id:       fmt.Sprintf("%s-version-%d", "go_brand_rpc", time.Now().Unix()),
		Name:     "go_brand_rpc",
		Port:     18090,
		Address:  "192.168.1.224",
		IsSecure: false,
	})
}
