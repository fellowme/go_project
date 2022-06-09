package stock_model

import (
	gin_model "github.com/fellowme/gin_common_library/model"
	"go_project/app/stock/stock_const"
)

type Stock struct {
	gin_model.BaseMysqlStruct
	ProductMainId int `json:"product_main_id,omitempty" gorm:"type:int(11);index:product_main_id;comment:product_main_id"`
	ProductId     int `json:"product_id,omitempty" gorm:"type:int(11);index:product_id;comment:商品id"`
	Stock         int `json:"stock" gorm:"type:int(5);comment:库存"`
}

func (receiver Stock) TableName() string {
	return stock_const.StockTableName
}
