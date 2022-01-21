package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
)

func main() {
	gin_app.CreateRpcServer("/app/user/user_config/", "go_user_rpc")
}
