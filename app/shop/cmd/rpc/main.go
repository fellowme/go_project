package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	grpc_consul "github.com/fellowme/gin_common_library/consul"
	"go_project/app/shop/shop_rpc"
	service "go_project/rpc_service"
	"time"
)

func main() {
	gRpcService := gin_app.InitRpcServer("/app/shop/shop_config/", "go_shop_rpc")
	shopService := shop_rpc.GetShopRpcService()
	service.RegisterShopServiceServer(gRpcService, &shopService)
	gin_app.CreateRpcServer(gRpcService, grpc_consul.ServiceConsul{
		Id:       fmt.Sprintf("%s-version-%d", "go_shop_rpc", time.Now().Unix()),
		Name:     "go_shop_rpc",
		Port:     18091,
		Address:  "192.168.1.224",
		IsSecure: false,
	})

}
