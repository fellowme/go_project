package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	gin_mysql "github.com/fellowme/gin_common_library/mysql"
	gin_router "github.com/fellowme/gin_common_library/router"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_project/app/shop/shop_model"
	"go_project/app/shop/shop_router"
	"syscall"
)

/*
	initRouter  初始化路由
*/

func initRouter(app *gin.Engine) {
	api := app.Group("/api/v1")
	shop_router.InitRouter(api)
}

/*
	initTable 初始化mysql 表信息
*/
func initTable() {
	err := gin_mysql.UseMysql(nil).AutoMigrate(&shop_model.Shop{})
	if err != nil {
		zap.L().Error("UseMysql error", zap.Any("error", err))
	}
}

/*
	主程序
*/
func main() {
	endPoint, app := gin_app.CreateServer("/app/shop/shop_config/", "go_shop")
	initRouter(app)
	initTable()
	gin_router.RegisterRouter(app.Routes())
	defer gin_app.DeferClose()
	server := endless.NewServer(endPoint, app)
	server.BeforeBegin = func(add string) {
		zap.L().Info(fmt.Sprintf("Actual pid is %d", syscall.Getpid()))
	}
	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Sprint("init server fail err=", err))
	}
}
