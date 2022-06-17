package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	grpc_consul "github.com/fellowme/gin_common_library/consul"
	"go_project/app/image/image_model"
	"go_project/app/image/image_router"
	"time"
)

/*
	主程序
*/
func main() {
	modelList := []interface{}{
		&image_model.Image{},
	}
	gin_app.CreateAppServer("/app/image/image_config/", "go_image", image_router.InitRouter, modelList, grpc_consul.ServiceConsul{
		Id:       fmt.Sprintf("%s-version-%d", "go_image", time.Now().Unix()),
		Name:     "go_image",
		Port:     8088,
		Address:  "192.168.1.224",
		IsSecure: false,
	})
}
