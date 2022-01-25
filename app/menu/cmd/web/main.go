package main

import (
	"fmt"
	gin_app "github.com/fellowme/gin_common_library/app"
	gin_mysql "github.com/fellowme/gin_common_library/mysql"
	gin_router "github.com/fellowme/gin_common_library/router"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_project/app/menu/menu_model"
	"go_project/app/menu/menu_router"
	"syscall"
)

func initRouter(app *gin.Engine) {
	api := app.Group("/api/v1")
	menu_router.InitRouter(api)
}

/*
	初始化mysql 表信息
*/
func initTable() {
	err := gin_mysql.UseMysql(nil).AutoMigrate(&menu_model.Menu{})
	if err != nil {
		zap.L().Error("menu AutoMigrate error", zap.Any("error", err))
	}
}

/*
	主程序
*/
func main() {
	endPoint, app := gin_app.CreateServer("/app/menu/menu_config/", "go_menu")
	defer gin_app.DeferClose()
	initTable()
	initRouter(app)
	gin_router.RegisterRouter(app.Routes())
	server := endless.NewServer(endPoint, app)
	server.BeforeBegin = func(add string) {
		zap.L().Info(fmt.Sprintf("Actual pid is %d", syscall.Getpid()))
	}
	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Sprint("init go_menu server fail err = ", err))
	}

}
