package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	"go_project/app/shop/shop_rpc"
	service "go_project/rpc_service"
)

func main() {
	gRpcService := gin_app.InitRpcServer("/app/shop/shop_config/", "go_shop_rpc")
	shopService := shop_rpc.GetShopRpcService()
	service.RegisterShopServiceServer(gRpcService, &shopService)
	gin_app.CreateRpcServer(gRpcService)

}
