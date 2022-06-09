package stock_cache

import (
	"errors"
	"fmt"
	"github.com/fellowme/gin_common_library/redis"
	"go_project/app/stock/stock_const"
)

func SetProductStockToRedis(productStockMap map[int]int64) (errorList []error) {
	if len(productStockMap) == 0 {
		return append(errorList, errors.New(stock_const.ProductStockEmptyTip))
	}
	for key, value := range productStockMap {
		newKey := fmt.Sprintf(stock_const.StockProductFmt, key)
		err := redis.SetKeyValue(stock_const.StockRedisName, newKey, value, stock_const.StockProductExpire)
		if err != nil {
			errorList = append(errorList, errors.New(stock_const.ProductStockEmptyTip))
		}
	}
	return
}
