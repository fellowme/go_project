package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	gin_mysql "github.com/fellowme/gin_common_library/mysql"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_project/app/role/role_model"
	"go_project/app/role/role_router"
	"syscall"
)

func initRouter(app *gin.Engine) {
	api := app.Group("/api/v1")
	role_router.InitRouter(api)
}

/*
	初始化mysql 表信息
*/
func initTable() {
	err := gin_mysql.UseMysql(nil).AutoMigrate(&role_model.Role{}, &role_model.RoleMenu{}, role_model.RoleUser{})
	if err != nil {
		zap.L().Error("role AutoMigrate error", zap.Any("error", err))
	}
}

/*
	主程序
*/
func main() {
	endPoint, app := gin_app.CreateServer("/app/role/role_config/", "go_role")
	initRouter(app)
	initTable()
	defer gin_app.DeferClose()
	server := endless.NewServer(endPoint, app)
	server.BeforeBegin = func(add string) {
		zap.L().Info(fmt.Sprintf("Actual pid is %d", syscall.Getpid()))
	}
	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Sprint("init server fail err = ", err))
	}
}
