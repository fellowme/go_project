package menu_param

import "github.com/fellowme/gin_common_library/param"

type (
	PostMenuRequestParam struct {
		MenuName string `json:"menu_name,omitempty" form:"menu_name"`
		Path     string `json:"path" form:"path" binding:"required"`
		Method   string `json:"method" form:"method" binding:"required"`
		Handler  string `json:"handler" form:"handler" binding:"required"`
		Remark   string `json:"remark,omitempty" form:"remark"`
	}
	GetMenuRequestParam struct {
		param.PageRequestParam
		MenuName string `json:"menu_name,omitempty" form:"menu_name"`
		Path     string `json:"path" form:"path" `
		Method   string `json:"method" form:"method"`
	}
	PatchMenuRequestParam struct {
		Id       int
		MenuName string `json:"menu_name,omitempty" form:"menu_name"`
		Path     string `json:"path" form:"path" `
		Method   string `json:"method" form:"method" `
		Handler  string `json:"handler" form:"handler" `
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
		Path     string `json:"path" `
		Method   string `json:"method" `
		Handler  string `json:"handler" `
		Remark   string `json:"remark,omitempty" `
	}
)
