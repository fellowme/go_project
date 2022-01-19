package menu_param

import "github.com/fellowme/gin_common_library/param"

type (
	PostMenuRequestParam struct {
		MenuName string `json:"menu_name,omitempty" form:"menu_name" binding:"required"`
		MenuPath string `json:"menu_path,omitempty" form:"menu_path" binding:"required"`
		MenuType int    `json:"menu_type" form:"menu_type"`
		Remark   string `json:"remark,omitempty" form:"remark"`
	}
	GetMenuRequestParam struct {
		param.PageRequestParam
		MenuName string `json:"menu_name,omitempty" form:"menu_name"`
		MenuPath string `json:"menu_path,omitempty" form:"menu_path"`
		MenuType int    `json:"menu_type" form:"menu_type"`
	}
	PatchMenuRequestParam struct {
		Id       int
		MenuName string `json:"menu_name,omitempty" form:"menu_name"`
		MenuPath string `json:"menu_path,omitempty" form:"menu_path"`
		MenuType int    `json:"menu_type" form:"menu_type"`
		Remark   string `json:"remark,omitempty" form:"remark"`
	}
)

type (
	MenuListResponse struct {
		Total int64          `json:"total,omitempty"`
		List  []MenuResponse `json:"list,omitempty"`
	}
	MenuResponse struct {
		Id       int    `json:"id"`
		MenuName string `json:"menu_name,omitempty" `
		MenuPath string `json:"menu_path,omitempty" `
		MenuType int    `json:"menu_type"`
		Remark   string `json:"remark,omitempty"`
	}
)
