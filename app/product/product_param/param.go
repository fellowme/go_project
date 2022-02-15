package product_param

import (
	gin_model "github.com/fellowme/gin_common_library/model"
	gin_param "github.com/fellowme/gin_common_library/param"
)

type (
	GetProductMainRequestParam struct {
		gin_param.PageRequestParam
	}
)

type (
	ProductMainResponse struct {
		Id                  int                 `json:"id"`
		BrandId             int                 `json:"brand_id,omitempty" `
		BrandName           string              `json:"brand_name"`
		ShopId              int                 `json:"shop_id,omitempty" `
		ShopName            string              `json:"shop_name"`
		ShortTitle          string              `json:"short_title,omitempty" `
		Title               string              `json:"title,omitempty" `
		Weight              int                 `json:"weight,omitempty" `
		ProductMainStatus   int                 `json:"product_main_status,omitempty" `
		CategoryId          int                 `json:"category_id,omitempty" `
		CategoryName        string              `json:"category_name"`
		ProductCode         string              `json:"product_code,omitempty" `
		ProductMainType     int                 `json:"product_main_type,omitempty" `
		ProductMainTypeName string              `json:"product_main_type_name"`
		SaleTime            gin_model.LocalTime `json:"sale_time" `
		Images              []int               `json:"images"`
		ImageMapList        []ImageParam        `json:"image_map"`
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
)
