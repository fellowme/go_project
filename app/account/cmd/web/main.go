package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	grpc_consul "github.com/fellowme/gin_common_library/consul"
	"go_project/app/account/account_model"
	"go_project/app/account/account_router"
	"time"
)

/*
	主程序
*/
func main() {
	modelList := []interface{}{
		account_model.Account{}, account_model.VerificationEmailCode{}, account_model.VerificationMobileCode{}, account_model.LoginTime{},
	}
	gin_app.CreateAppServer("/app/account/account_config/", "go_account", account_router.InitRouter, modelList, grpc_consul.ServiceConsul{
		Id:       fmt.Sprintf("%s-version-%d", "go_account", time.Now().Unix()),
		Name:     "go_account",
		Port:     8083,
		Address:  "192.168.1.224",
		IsSecure: false,
	})

}
