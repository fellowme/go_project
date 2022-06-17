package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	grpc_consul "github.com/fellowme/gin_common_library/consul"
	"go_project/app/category/category_rpc"
	service "go_project/rpc_service"
	"time"
)

func main() {
	gRpcService := gin_app.InitRpcServer("/app/category/category_config/", "go_category_rpc")
	categoryService := category_rpc.GetCategoryRpcService()
	service.RegisterCategoryServiceServer(gRpcService, categoryService)
	gin_app.CreateRpcServer(gRpcService, grpc_consul.ServiceConsul{
		Id:       fmt.Sprintf("%s-version-%d", "go_category_rpc", time.Now().Unix()),
		Name:     "go_category_rpc",
		Port:     18089,
		Address:  "192.168.1.224",
		IsSecure: false,
	})
}
