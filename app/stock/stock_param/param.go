package stock_param

import gin_param "github.com/fellowme/gin_common_library/param"

type (
	GetStockRequestParam struct {
		gin_param.PageRequestParam
		ProductMainId int `json:"product_main_id,omitempty" form:"product_main_id,omitempty" `
		ProductId     int `json:"product_id" form:"product_id,omitempty"`
	}
	PostStockRequestParam struct {
		Id            int
		ProductMainId int   `json:"product_main_id,omitempty" form:"product_main_id,omitempty" `
		ProductId     int   `json:"product_id" form:"product_id,omitempty"`
		StockNumber   int64 `json:"stock_number,omitempty" form:"stock_number"`
	}
	PostStockByIdsRequestParam struct {
		Ids               string `json:"ids" form:"ids"`
		IdList            []int
		ProductMainIds    string `json:"product_main_ids,omitempty" form:"product_main_ids,omitempty" `
		ProductMainIdList []int
		ProductIds        string `json:"product_ids" form:"product_ids,omitempty"`
		ProductIdList     []int
	}
	PostStockTorRedisByIdsRequestParam struct {
		ProductMainIds    string `json:"product_main_ids,omitempty" form:"product_main_ids,omitempty" `
		ProductMainIdList []int
		ProductIds        string `json:"product_ids" form:"product_ids,omitempty"`
		ProductIdList     []int
	}
)

type (
	StockResponse struct {
		Id              int    `json:"id,omitempty"`
		ProductMainId   int    `json:"product_main_id,omitempty"`
		ProductId       int    `json:"product_id"`
		ProductName     string `json:"product_name,omitempty"`
		ProductMainName string `json:"product_main_name,omitempty"`
		StockNumber     int64  `json:"stock_number,omitempty"`
	}
	StockListResponse struct {
		Total int64           `json:"total,omitempty"`
		List  []StockResponse `json:"list,omitempty"`
	}
)

type (
	ProductMainStockParam struct {
		ProductMainId int   `json:"product_main_id,omitempty"`
		ProductId     int   `json:"product_id"`
		StockTotal    int64 `json:"stock_total,omitempty"`
	}

	ProductParam struct {
		ProductMainId   int    `json:"product_main_id,omitempty"`
		ProductId       int    `json:"product_id"`
		ProductName     string `json:"product_name,omitempty"`
		ProductMainName string `json:"product_main_name,omitempty"`
	}
)
