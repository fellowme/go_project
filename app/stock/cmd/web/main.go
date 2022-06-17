package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	grpc_consul "github.com/fellowme/gin_common_library/consul"
	"go_project/app/stock/stock_model"
	"go_project/app/stock/stock_router"
)

/*
	主程序
*/
func main() {
	modelList := []interface{}{&stock_model.Stock{}}
	gin_app.CreateAppServer("/app/stock/stock_config/", "go_stock", stock_router.InitRouter, modelList, grpc_consul.ServiceConsul{
		Id:       "go_stock_v1",
		Name:     "go_stock",
		Port:     8094,
		Address:  "192.168.1.224",
		IsSecure: false,
	})
}
