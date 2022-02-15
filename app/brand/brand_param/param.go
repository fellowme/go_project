package brand_param

import gin_param "github.com/fellowme/gin_common_library/param"

type (
	GetBrandRequestParam struct {
		gin_param.PageRequestParam
		BrandName string `json:"brand_name" form:"brand_name"`
	}

	PostBrandRequestParam struct {
		Id           int
		BrandName    string `json:"brand_name,omitempty" form:"brand_name"`
		BrandImageId int    `json:"brand_image_id,omitempty" form:"brand_image_id"`
		BrandDetail  string `json:"brand_detail,omitempty" form:"brand_detail"`
		BrandWeight  int    `json:"brand_weight,omitempty" form:"brand_weight"`
		BrandStatus  int    `json:"brand_status,omitempty" form:"brand_status"`
	}
	GetBrandByIdsRequestParam struct {
		BrandIds    string `json:"brand_ids" form:"brand_ids"`
		BrandIdList []int
	}
	DeleteBrandByIdsRequestParam struct {
		GetBrandByIdsRequestParam
	}
)

type (
	BrandResponse struct {
		Id            int    `json:"id"`
		BrandName     string `json:"brand_name,omitempty"`
		BrandImageId  int    `json:"brand_image_id,omitempty"`
		BrandImageUrl string `json:"brand_image_url"`
		BrandDetail   string `json:"brand_detail,omitempty"`
		BrandWeight   int    `json:"brand_weight,omitempty"`
		BrandStatus   int    `json:"brand_status,omitempty"`
	}

	BrandListResponse struct {
		Total int64           `json:"total,omitempty"`
		List  []BrandResponse `json:"list,omitempty"`
	}
)
