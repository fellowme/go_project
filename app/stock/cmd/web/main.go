package main

import (
	gin_app "github.com/fellowme/gin_common_library/app"
	"go_project/app/stock/stock_model"
	"go_project/app/stock/stock_router"
)

/*
	主程序
*/
func main() {
	modelList := []interface{}{&stock_model.Stock{}}
	gin_app.CreateAppServer("/app/stock/stock_config/", "go_stock", stock_router.InitRouter, modelList)
}
