package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	"go_project/app/stock/stock_rpc"
	service "go_project/rpc_service"
)

func main() {
	gRpcService := gin_app.InitRpcServer("/app/stock/stock_config/", "go_stock_rpc")
	stockService := stock_rpc.GetStockRpcService()
	service.RegisterStockServiceServer(gRpcService, stockService)
	gin_app.CreateRpcServer(gRpcService)
}
