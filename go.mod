module go_project

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fellowme/gin_common_library v0.0.20
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.7.4
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gomodule/redigo v1.8.5
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.10.0
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.26.0
	gorm.io/gorm v1.22.4
)

replace github.com/spf13/viper v1.9.0 => github.com/spf13/viper v1.6.3
