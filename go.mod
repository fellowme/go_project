module go_project

go 1.15

require (
	github.com/apache/pulsar-client-go v0.8.0
	github.com/fellowme/gin_common_library v0.0.36
	github.com/gin-gonic/gin v1.7.7
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/olivere/elastic/v7 v7.0.29
	github.com/pkg/errors v0.9.1
	go.uber.org/zap v1.17.0
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
	gorm.io/gorm v1.22.4
)

replace github.com/spf13/viper v1.8.1 => github.com/spf13/viper v1.6.3
