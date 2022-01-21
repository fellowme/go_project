package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
)

func main() {
	gin_app.CreateRpcServer("/app/menu/menu_config/", "go_menu_rpc")
}
