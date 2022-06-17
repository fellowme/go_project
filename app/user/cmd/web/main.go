package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	grpc_consul "github.com/fellowme/gin_common_library/consul"
	"go_project/app/user/user_model"
	"go_project/app/user/user_router"
	"time"
)

/*
	主程序
*/
func main() {
	modelList := []interface{}{
		&user_model.User{},
	}
	gin_app.CreateAppServer("/app/user/user_config/", "go_user", user_router.InitRouter, modelList, grpc_consul.ServiceConsul{
		Id:       fmt.Sprintf("%s-version-%d", "go_user", time.Now().Unix()),
		Name:     "go_user",
		Port:     8082,
		Address:  "192.168.1.224",
		IsSecure: false,
	})

}
