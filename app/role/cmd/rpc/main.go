package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
)

func main() {
	gin_app.CreateRpcServer("/app/role/role_config/", "go_role_rpc")
}
