package main

import (
	"fmt"
	gin_config "github.com/fellowme/gin_common_library/config"
	gin_jaeger "github.com/fellowme/gin_common_library/jaeger"
	gin_logger "github.com/fellowme/gin_common_library/logger"
	gin_mysql "github.com/fellowme/gin_common_library/mysql"
	gin_redis "github.com/fellowme/gin_common_library/redis"
	gin_translator "github.com/fellowme/gin_common_library/translator"
	gin_util "github.com/fellowme/gin_common_library/util"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"go.uber.org/zap"
	"go_project/app/role/role_rpc"
	service "go_project/rpc_service"
	"google.golang.org/grpc"
	"net"
)

func initExtend() {
	path := gin_util.GetPath()
	gin_config.InitConfig(path+"/app/role/role_config/", "go_role_rpc")
	gin_logger.InitServerLogger(path)
	gin_logger.InitRecoveryLogger(path)
	gin_redis.InitRedis()
	gin_mysql.InitMysqlMap()
	gin_jaeger.InitJaegerTracer()
	gin_translator.InitTranslator()
}

func createRpc() *grpc.Server {
	initExtend()
	roleService := role_rpc.GetRoleRpcService()
	gRpcService := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zap.L()),
		)))
	service.RegisterUserRoleMenuServiceServer(gRpcService, &roleService)
	return gRpcService
}

func main() {
	gRpcService := createRpc()
	defer gin_mysql.CloseMysqlConnect()
	defer gin_jaeger.IoCloser()
	defer gRpcService.GracefulStop()
	endPoint := fmt.Sprintf("%s:%d", gin_config.ServerConfigSettings.Server.ServerHost,
		gin_config.ServerConfigSettings.Server.ServerRpcPort)
	listener, err := net.Listen("tcp", endPoint)
	if err != nil {
		zap.L().Error("rpc listener error", zap.Any("error", err))
		return
	}
	err = gRpcService.Serve(listener)
	zap.L().Info("grpc server 启动")
	if err != nil {
		zap.L().Error("rpc Serve error", zap.Any("error", err))
		return
	}
}
