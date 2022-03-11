package category_param

import "github.com/fellowme/gin_common_library/param"

type (
	GetCategoryListRequestParam struct {
		param.PageRequestParam
		CategoryName string `json:"category_name,omitempty" form:"category_name"`
		Id           int    `json:"id,omitempty" form:"id"`
	}
	CategoryRequestParam struct {
		Id               int
		CategoryName     string `json:"category_name,omitempty" form:"category_name"`
		CategoryParentId int    `json:"category_parent_id,omitempty" form:"category_parent_id"`
		CategorySort     int    `json:"category_sort,omitempty" form:"category_sort"`
	}
)

type (
	CategoryListResponse struct {
		Total int64              `json:"total,omitempty"`
		List  []CategoryResponse `json:"list,omitempty"`
	}
	CategoryResponse struct {
		CategoryParam
	}
)

type (
	CategoryTreeParam struct {
		Id               int                  `json:"id,omitempty"`
		CategoryName     string               `json:"category_name,omitempty"`
		CategoryParentId int                  `json:"category_parent_id,omitempty"`
		CategorySort     int                  `json:"category_sort,omitempty"`
		CategoryChildren []*CategoryTreeParam `json:"category_children"`
	}

	CategoryParam struct {
		Id               int    `json:"id,omitempty"`
		CategoryName     string `json:"category_name,omitempty"`
		CategoryParentId int    `json:"category_parent_id,omitempty"`
		CategorySort     int    `json:"category_sort,omitempty"`
	}
)
