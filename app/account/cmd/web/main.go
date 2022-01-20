package main

import (
	"fmt"
	"go_project/app/account/account_model"
	"go_project/app/account/account_router"
	"syscall"
	"time"

	gin_config "github.com/fellowme/gin_common_library/config"
	gin_jaeger "github.com/fellowme/gin_common_library/jaeger"
	gin_logger "github.com/fellowme/gin_common_library/logger"
	gin_mysql "github.com/fellowme/gin_common_library/mysql"
	gin_redis "github.com/fellowme/gin_common_library/redis"
	gin_translator "github.com/fellowme/gin_common_library/translator"
	gin_util "github.com/fellowme/gin_common_library/util"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/*
	初始化配置文件
*/
func initExtend() {
	path := gin_util.GetPath()
	gin_config.InitConfig(path+"/app/account/account_config/", "go_account")
	gin_logger.InitServerLogger(path)
	gin_logger.InitRecoveryLogger(path)
	gin_redis.InitRedis()
	gin_mysql.InitMysqlMap()
	gin_jaeger.InitJaegerTracer()
	gin_translator.InitTranslator()
}

func initRouter(app *gin.Engine) {
	api := app.Group("/api/v1")
	account_router.InitRouter(api)
}

/*
	初始化app
*/
func creatApp() *gin.Engine {
	initExtend()
	app := gin.New()
	app.Use(gin_logger.RecoveryWithZap(gin_logger.RecoveryLogger,
		gin_config.ServerConfigSettings.Server.IsDebug), gin_jaeger.JaegerMiddleWare())
	initRouter(app)
	return app
}

/*
	初始化mysql 表信息
*/
func initTable() {
	err := gin_mysql.UseMysql(nil).AutoMigrate(account_model.Account{}, account_model.VerificationEmailCode{}, account_model.VerificationMobileCode{}, account_model.LoginTime{})
	if err != nil {
		zap.L().Error("account AutoMigrate error", zap.Any("error", err))
	}
}

/*
	主程序
*/
func main() {
	if !gin_config.ServerConfigSettings.Server.IsDebug {
		gin.SetMode(gin.ReleaseMode)
	}
	app := creatApp()
	defer gin_mysql.CloseMysqlConnect()
	defer gin_jaeger.IoCloser()
	initTable()
	endless.DefaultReadTimeOut = time.Duration(gin_config.ServerConfigSettings.Server.ReadTimeout) * time.Second
	endless.DefaultWriteTimeOut = time.Duration(gin_config.ServerConfigSettings.Server.WriteTimeout) * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf("%s:%d", gin_config.ServerConfigSettings.Server.ServerHost,
		gin_config.ServerConfigSettings.Server.ServerPort)
	server := endless.NewServer(endPoint, app)
	server.BeforeBegin = func(add string) {
		zap.L().Info(fmt.Sprintf("Actual pid is %d", syscall.Getpid()))
	}
	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Sprint("init server fail err = ", err))
	}
}
