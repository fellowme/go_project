package product_param

import (
	gin_model "github.com/fellowme/gin_common_library/model"
	gin_param "github.com/fellowme/gin_common_library/param"
)

type (
	// GetProductMainRequestParam  spu 请求参数
	GetProductMainRequestParam struct {
		gin_param.PageRequestParam
	}
	PostProductMainRequestParam struct {
		Id                int
		BrandId           int     `json:"brand_id,omitempty" form:"brand_id"`
		ShopId            int     `json:"shop_id,omitempty" form:"shop_id"`
		ShortTitle        string  `json:"short_title,omitempty" form:"short_title"`
		Title             string  `json:"title,omitempty" form:"title"`
		Weight            int     `json:"weight,omitempty" form:"weight"`
		ProductMainStatus int     `json:"product_main_status,omitempty" form:"product_main_status"`
		CategoryId        int     `json:"category_id,omitempty" form:"category_id"`
		ProductCode       string  `json:"product_code,omitempty" form:"product_code"`
		ProductMainType   int     `json:"product_main_type,omitempty" form:"product_main_type"`
		SaleTime          *string `json:"sale_time" form:"sale_time"`
		Images            string  `json:"images" form:"images" gorm:"-"`
	}
	PostProductMainExtRequestParam struct {
		PostProductMainRequestParam
		ImageIdList []int `gorm:"-"`
	}
	PostDeleteProductMainAllRequestParam struct {
		Ids    string `json:"ids" form:"ids"`
		IdList []int
	}

	// PostProductRequestParam sku 请求参数
	PostProductRequestParam struct {
		Id            int
		ProductMainId int    `json:"product_main_id,omitempty" form:"product_main_id,omitempty" `
		ShortTitle    string `json:"short_title,omitempty" form:"short_title,omitempty"`
		Title         string `json:"title,omitempty" form:"title,omitempty"`
		ProductStatus int    `json:"product_status,omitempty" form:"product_status,omitempty"`
		IsMainProduct bool   `json:"is_main_product,omitempty" form:"is_main_product,omitempty"`
		Price         string `json:"price,omitempty" form:"price,omitempty"`
		RealPrice     string `json:"real_price,omitempty" form:"real_price,omitempty"`
		Description   string `json:"description,omitempty" form:"description,omitempty"`
		Weight        int    `json:"weight,omitempty" form:"weight,omitempty"`
		Images        string `json:"images" form:"images" gorm:"-"`
		Stock         int    `json:"stock" form:"stock" gorm:"-" `
	}
	PostProductExtRequestParam struct {
		PostProductRequestParam
		ImageIdList []int `gorm:"-"`
	}
	GetProductRequestParam struct {
		gin_param.PageRequestParam
	}

	GetProductStockRequestParam struct {
		gin_param.PageRequestParam
		ProductMainId int `json:"product_main_id,omitempty" form:"product_main_id,omitempty" `
		ProductId     int `json:"product_id" form:"product_id,omitempty"`
	}
	PostProductStockRequestParam struct {
		Id            int
		ProductMainId int   `json:"product_main_id,omitempty" form:"product_main_id,omitempty" `
		ProductId     int   `json:"product_id" form:"product_id,omitempty"`
		StockNumber   int64 `json:"stock_number,omitempty" form:"stock_number"`
	}
	PostProductStockByIdsRequestParam struct {
		Ids               string `json:"ids" form:"ids"`
		IdList            []int
		ProductMainIds    string `json:"product_main_ids,omitempty" form:"product_main_ids,omitempty" `
		ProductMainIdList []int
		ProductIds        string `json:"product_ids" form:"product_ids,omitempty"`
		ProductIdList     []int
	}
	PostProductIdsRequestParam struct {
		Ids               string `json:"ids" form:"ids"`
		IdList            []int
		ProductMainIds    string `json:"product_main_ids,omitempty" form:"product_main_ids,omitempty" `
		ProductMainIdList []int
	}
	PostProductMainIdsRequestParam struct {
		Ids    string `json:"ids" form:"ids"`
		IdList []int
	}
	DeletePostProductIdsRequestParam struct {
		Ids    string `json:"ids" form:"ids"`
		IdList []int
	}
	PostProductIdsToMqRequestParam struct {
		Ids    string `json:"ids" form:"ids"`
		IdList []int
	}
)

type (
	ProductMainResponse struct {
		Id                int                 `json:"id"`
		BrandId           int                 `json:"brand_id,omitempty" `
		ShopId            int                 `json:"shop_id,omitempty" `
		ShortTitle        string              `json:"short_title,omitempty" `
		Title             string              `json:"title,omitempty" `
		Weight            int                 `json:"weight,omitempty" `
		ProductMainStatus int                 `json:"product_main_status,omitempty" `
		CategoryId        int                 `json:"category_id,omitempty" `
		ProductCode       string              `json:"product_code,omitempty" `
		ProductMainType   int                 `json:"product_main_type,omitempty" `
		SaleTime          gin_model.LocalTime `json:"sale_time" `
	}
	ProductMainExtResponse struct {
		ProductMainResponse
		Stock                   int64        `json:"stock"`
		ProductMainStatusString string       `json:"product_main_status_string"`
		SaleTimeString          string       `json:"sale_time_string"`
		ProductMainTypeName     string       `json:"product_main_type_name"`
		BrandName               string       `json:"brand_name"`
		CategoryName            string       `json:"category_name"`
		ShopName                string       `json:"shop_name"`
		Images                  []int        `json:"images"`
		ImageMapList            []ImageParam `json:"image_map_list"`
	}

	ProductMainListResponse struct {
		Total int64                    `json:"total,omitempty"`
		List  []ProductMainExtResponse `json:"list,omitempty"`
	}

	ProductMainExtEsResponse struct {
		ProductMainResponse
		Stock                   int64           `json:"stock"`
		ProductMainStatusString string          `json:"product_main_status_string"`
		SaleTimeString          string          `json:"sale_time_string"`
		ProductMainTypeName     string          `json:"product_main_type_name"`
		BrandName               string          `json:"brand_name"`
		CategoryName            string          `json:"category_name"`
		ShopName                string          `json:"shop_name"`
		Images                  []int           `json:"images"`
		ImageMapList            []ImageParam    `json:"image_map_list"`
		Product                 ProductResponse `json:"product"`
	}

	ProductResponse struct {
		Id            int    `json:"id,omitempty"`
		ProductMainId int    `json:"product_main_id,omitempty"`
		ShortTitle    string `json:"short_title,omitempty" `
		Title         string `json:"title,omitempty"`
		ProductStatus int    `json:"product_status,omitempty"`
		IsMainProduct bool   `json:"is_main_product,omitempty"`
		Price         string `json:"price,omitempty"`
		RealPrice     string `json:"real_price,omitempty"`
		Description   string `json:"description,omitempty"`
		Weight        int    `json:"weight,omitempty"`
	}
	ProductExtResponse struct {
		ProductResponse
		ProductStatusString string       `json:"product_status_string"`
		Stock               int64        `json:"stock"`
		Images              []int        `json:"images"`
		ImageMapList        []ImageParam `json:"image_map_list"`
	}
	ProductListResponse struct {
		Total int64                `json:"total,omitempty"`
		List  []ProductExtResponse `json:"list,omitempty"`
	}

	ProductStockResponse struct {
		Id              int    `json:"id,omitempty"`
		ProductMainId   int    `json:"product_main_id,omitempty"`
		ProductId       int    `json:"product_id"`
		ProductName     string `json:"product_name,omitempty"`
		ProductMainName string `json:"product_main_name,omitempty"`
		StockNumber     int64  `json:"stock_number,omitempty"`
	}
	ProductStockListResponse struct {
		Total int64                  `json:"total,omitempty"`
		List  []ProductStockResponse `json:"list,omitempty"`
	}
)

type (
	ProductImageParam struct {
		ProductId int    `json:"product_id,omitempty"`
		ImageIds  string `json:"image_ids"`
	}
	ProductMainStockParam struct {
		ProductMainId int   `json:"product_main_id,omitempty"`
		ProductId     int   `json:"product_id"`
		StockTotal    int64 `json:"stock_total,omitempty"`
	}
	ImageParam struct {
		Id        int    `json:"id"`
		ImageUrl  string `json:"image_url,omitempty" `
		ImageName string `json:"image_name,omitempty" `
	}
	BrandParam struct {
		Id        int
		BrandName string
	}
	ShopParam struct {
		Id       int
		ShopName string
	}
)
