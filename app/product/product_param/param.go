package product_param

import (
	gin_param "github.com/fellowme/gin_common_library/param"
)

type (
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
		ImageIdList       []int   `gorm:"-"`
	}
)

type (
	ProductMainResponse struct {
		Id                  int          `json:"id"`
		BrandId             int          `json:"brand_id,omitempty" `
		BrandName           string       `json:"brand_name"`
		ShopId              int          `json:"shop_id,omitempty" `
		ShopName            string       `json:"shop_name"`
		ShortTitle          string       `json:"short_title,omitempty" `
		Title               string       `json:"title,omitempty" `
		Weight              int          `json:"weight,omitempty" `
		ProductMainStatus   int          `json:"product_main_status,omitempty" `
		CategoryId          int          `json:"category_id,omitempty" `
		CategoryName        string       `json:"category_name"`
		ProductCode         string       `json:"product_code,omitempty" `
		ProductMainType     int          `json:"product_main_type,omitempty" `
		ProductMainTypeName string       `json:"product_main_type_name"`
		SaleTime            *string      `json:"sale_time" `
		Images              []int        `json:"images"`
		ImageMapList        []ImageParam `json:"image_map_list"`
	}
	ProductMainsResponse struct {
		ProductMains []ProductMainResponse
	}
	ProductMainListResponse struct {
		Total int64                 `json:"total,omitempty"`
		List  []ProductMainResponse `json:"list,omitempty"`
	}
)

type (
	ProductImageParam struct {
		ProductId int    `json:"product_id,omitempty"`
		ImageIds  string `json:"image_ids"`
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
