package role_param

import "github.com/fellowme/gin_common_library/param"

type (
	GetRoleListRequestParam struct {
		param.PageRequestParam
		RoleName string `json:"role_name" form:"role_name"`
	}
	PostRoleRequestParam struct {
		RoleName string `json:"role_name,omitempty" form:"role_name" binding:"required"`
		Remark   string `json:"remark,omitempty" form:"remark"`
	}
	PatchRoleRequestParam struct {
		Id       int
		RoleName string `json:"role_name,omitempty" form:"role_name" `
		Remark   string `json:"remark,omitempty" form:"remark"`
	}
	PostRoleUserRequestParam struct {
		RoleId int    `json:"role_id,omitempty" form:"role_id" binding:"required"`
		UserId int    `json:"user_id,omitempty" form:"user_id" binding:"required"`
		Remark string `json:"remark,omitempty" form:"remark"`
	}
	GetRoleUserRequestParam struct {
		param.PageRequestParam
		RoleId int `json:"role_id,omitempty" form:"role_id"`
		UserId int `json:"user_id,omitempty" form:"user_id"`
	}
	GetRoleMenuRequestParam struct {
		param.PageRequestParam
		RoleId int `json:"role_id,omitempty" form:"role_id"`
		MenuId int `json:"menu_id" form:"menu_id"`
	}
	PostRoleMenuRequestParam struct {
		RoleId int    `json:"role_id,omitempty" form:"role_id" binding:"required"`
		MenuId int    `json:"menu_id" form:"menu_id" binding:"required"`
		Remark string `json:"remark,omitempty" form:"remark"`
	}
)

type (
	RoleListResponse struct {
		Total int64       `json:"total,omitempty"`
		List  []RoleParam `json:"list,omitempty"`
	}
	RoleUserListResponse struct {
		Total int64           `json:"total,omitempty"`
		List  []RoleUserParam `json:"list,omitempty"`
	}
	RoleMenuListResponse struct {
		Total int64           `json:"total,omitempty"`
		List  []RoleMenuParam `json:"list,omitempty"`
	}
)

type (
	RoleParam struct {
		Id       int    `json:"id,omitempty"`
		RoleName string `json:"role_name,omitempty"`
		Remark   string `json:"remark,omitempty"`
	}
	RoleUserParam struct {
		Id       int    `json:"id,omitempty"`
		RoleId   int    `json:"role_id"`
		RoleName string `json:"role_name,omitempty"`
		UserId   int    `json:"user_id"`
		UserName string `json:"user_name"`
		Remark   string `json:"remark,omitempty"`
	}
	RoleMenuParam struct {
		Id       int    `json:"id,omitempty"`
		RoleId   int    `json:"role_id"`
		RoleName string `json:"role_name,omitempty"`
		MenuId   int    `json:"menu_id"`
		MenuName string `json:"menu_name,omitempty" `
		Path     string `json:"path" `
		Method   string `json:"method" `
		Handler  string `json:"handler" `
		Remark   string `json:"remark,omitempty"`
	}
)
