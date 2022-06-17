package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	grpc_consul "github.com/fellowme/gin_common_library/consul"
	"go_project/app/brand/brand_router"
	"time"
)

/*
	主程序
*/
func main() {
	//modelList := []interface{}{
	//	&brand_model.Brand{},
	//}
	gin_app.CreateAppServer("/app/brand/brand_config/", "go_brand", brand_router.InitRouter, nil, grpc_consul.ServiceConsul{
		Id:       fmt.Sprintf("%s-version-%d", "go_brand", time.Now().Unix()),
		Name:     "go_brand",
		Port:     8090,
		Address:  "192.168.1.224",
		IsSecure: false,
	})
}
