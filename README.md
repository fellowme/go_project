# go_project

## 安装swagger

- go get -u github.com/swaggo/swag/cmd/swag@v1.1.0
- swag init
- go get -u github.com/swaggo/gin-swagger
- go get -u github.com/swaggo/files

## 执行vendor

- go mod tidy
- go mod vendor

## protobuf_file

    protobuf 文件 执行 sh grpc.sh  生成 pb.go文件