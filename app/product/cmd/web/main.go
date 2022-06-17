package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	grpc_consul "github.com/fellowme/gin_common_library/consul"
	"go_project/app/product/product_router"
	"time"
)

/*
	主程序
*/
func main() {
	//modelList := []interface{}{
	//	&product_model.Product{}, &product_model.ProductMain{},
	//	&product_model.ProductMain{}, &product_model.ProductImage{}, &product_model.Stock{},
	//}
	gin_app.CreateAppServer("/app/product/product_config/", "go_product", product_router.InitRouter, nil, grpc_consul.ServiceConsul{
		Id:       fmt.Sprintf("%s-version-%d", "go_product", time.Now().Unix()),
		Name:     "go_product",
		Port:     8092,
		Address:  "192.168.1.224",
		IsSecure: false,
	})
}
