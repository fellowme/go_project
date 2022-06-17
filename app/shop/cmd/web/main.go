package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	grpc_consul "github.com/fellowme/gin_common_library/consul"
	"go_project/app/shop/shop_model"
	"go_project/app/shop/shop_router"
	"time"
)

/*
	主程序
*/
func main() {
	modelList := []interface{}{
		&shop_model.Shop{},
	}
	gin_app.CreateAppServer("/app/shop/shop_config/", "go_shop", shop_router.InitRouter, modelList, grpc_consul.ServiceConsul{
		Id:       fmt.Sprintf("%s-version-%d", "go_shop", time.Now().Unix()),
		Name:     "go_shop",
		Port:     8091,
		Address:  "192.168.1.224",
		IsSecure: false,
	})
}
