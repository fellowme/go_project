package shop_param

import gin_param "github.com/fellowme/gin_common_library/param"

type (
	GetShopListRequestParam struct {
		gin_param.PageRequestParam
		ShopName string `json:"shop_name" form:"shop_name"`
	}
	PostShopRequestParam struct {
		ShopName    string `json:"shop_name,omitempty" form:"shop_name"`
		ShopImageId int    `json:"shop_image_id,omitempty" form:"shop_image_id"`
		ShopDetail  string `json:"shop_detail,omitempty" form:"shop_detail"`
		ShopWeight  *int   `json:"shop_weight,omitempty" form:"shop_weight"`
		ShopStatus  *int   `json:"shop_status,omitempty" form:"shop_status"`
	}
	PatchShopRequestParam struct {
		Id int
		PostShopRequestParam
	}
	GetShopByIdsRequestParam struct {
		ShopIds    string `json:"shop_ids,omitempty" form:"shop_ids" `
		ShopIdList []int
	}
	DeleteShopByIdsRequestParam struct {
		GetShopByIdsRequestParam
	}
)

type (
	ShopListResponse struct {
		Total int64          `json:"total,omitempty"`
		List  []ShopResponse `json:"list,omitempty"`
	}

	ShopResponse struct {
		Id           int    `json:"id"`
		ShopName     string `json:"shop_name,omitempty"`
		ShopImageId  int    `json:"shop_image_id,omitempty"`
		ShopImageUrl string `json:"shop_image_url"`
		ShopDetail   string `json:"shop_detail,omitempty"`
		ShopWeight   int    `json:"shop_weight,omitempty"`
		ShopStatus   int    `json:"shop_status,omitempty"`
	}
)
