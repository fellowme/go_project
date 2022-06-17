package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	grpc_consul "github.com/fellowme/gin_common_library/consul"
	"go_project/app/stock/stock_rpc"
	service "go_project/rpc_service"
	"time"
)

func main() {
	gRpcService := gin_app.InitRpcServer("/app/stock/stock_config/", "go_stock_rpc")
	stockService := stock_rpc.GetStockRpcService()
	service.RegisterStockServiceServer(gRpcService, stockService)
	gin_app.CreateRpcServer(gRpcService, grpc_consul.ServiceConsul{
		Id:       fmt.Sprintf("%s-version-%d", "go_stock_rpc", time.Now().Unix()),
		Name:     "go_stock_rpc",
		Port:     18094,
		Address:  "192.168.1.224",
		IsSecure: false,
	})
}
