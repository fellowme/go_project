package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	"go_project/app/account/account_model"
	"go_project/app/account/account_router"
)

/*
	主程序
*/
func main() {
	modelList := []interface{}{
		account_model.Account{}, account_model.VerificationEmailCode{}, account_model.VerificationMobileCode{}, account_model.LoginTime{},
	}
	gin_app.CreateAppServer("/app/account/account_config/", "go_account", account_router.InitRouter, modelList)

}
