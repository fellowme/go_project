package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	grpc_consul "github.com/fellowme/gin_common_library/consul"
	"go_project/app/menu/menu_model"
	"go_project/app/menu/menu_router"
	"time"
)

/*
	主程序
*/
func main() {
	modelList := []interface{}{
		&menu_model.Menu{},
	}
	gin_app.CreateAppServer("/app/menu/menu_config/", "go_menu", menu_router.InitRouter, modelList, grpc_consul.ServiceConsul{
		Id:       fmt.Sprintf("%s-version-%d", "go_menu", time.Now().Unix()),
		Name:     "go_menu",
		Port:     8086,
		Address:  "192.168.1.224",
		IsSecure: false,
	})

}
