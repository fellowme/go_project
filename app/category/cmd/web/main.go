package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	grpc_consul "github.com/fellowme/gin_common_library/consul"
	"go_project/app/category/category_model"
	"go_project/app/category/category_router"
	"time"
)

/*
	主程序
*/
func main() {
	modelList := []interface{}{
		&category_model.Category{},
	}
	gin_app.CreateAppServer("/app/category/category_config/", "go_category", category_router.InitRouter, modelList, grpc_consul.ServiceConsul{
		Id:       fmt.Sprintf("%s-version-%d", "go_category", time.Now().Unix()),
		Name:     "go_category",
		Port:     8089,
		Address:  "192.168.1.224",
		IsSecure: false,
	})
}
