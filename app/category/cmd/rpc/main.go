package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	"go_project/app/category/category_rpc"
	service "go_project/rpc_service"
)

func main() {
	gRpcService := gin_app.InitRpcServer("/app/category/category_config/", "go_category_rpc")
	categoryService := category_rpc.GetCategoryRpcService()
	service.RegisterCategoryServiceServer(gRpcService, categoryService)
	gin_app.CreateRpcServer(gRpcService)
}
