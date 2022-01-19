package user_param

import gin_param "github.com/fellowme/gin_common_library/param"

type (
	UserByIdsRequestParam struct {
		UserIds string `json:"user_ids" binding:"required" form:"user_ids"`
	}

	UserRequestParam struct {
		Id        int    `json:"id,omitempty"`
		AccountId int    `json:"account_id" form:"account_id" binding:"required"`
		UserName  string `json:"user_name,omitempty" form:"user_name" binding:"required"`
		NickName  string `json:"nick_name,omitempty" form:"nick_name"`
		RealName  string `json:"real_name,omitempty" form:"real_name"`
		Gender    int    `json:"gender,omitempty" form:"gender"`
	}

	UserPatchRequestParam struct {
		Id         int    `json:"id,omitempty"`
		UserName   string `json:"user_name,omitempty" form:"user_name"`
		NickName   string `json:"nick_name,omitempty" form:"nick_name"`
		RealName   string `json:"real_name,omitempty" form:"real_name"`
		Gender     int    `json:"gender,omitempty" form:"gender"`
		UserStatus int    `json:"user_status,omitempty" form:"user_status"`
	}

	UserListRequestParam struct {
		gin_param.PageRequestParam
		UserName string `json:"user_name,omitempty" form:"user_name"`
	}
	DeleteUserListRequestParam struct {
		Ids string `json:"ids" form:"ids"`
	}
)

type (
	UserInfoResponse struct {
		Id         int    `json:"id,omitempty"`
		UserName   string `json:"user_name,omitempty"`
		NickName   string `json:"nick_name,omitempty"`
		RealName   string `json:"real_name,omitempty"`
		Gender     string `json:"sex,omitempty"`
		UserStatus string `json:"user_status,omitempty"`
	}
	UserListResponse struct {
		Total int64              `json:"total,omitempty"`
		List  []UserInfoResponse `json:"list,omitempty"`
	}
)

type (
	UserParam struct {
		Id         int    `json:"id,omitempty"`
		AccountId  int    `json:"account_id"`
		UserName   string `json:"user_name,omitempty"`
		NickName   string `json:"nick_name,omitempty"`
		RealName   string `json:"real_name,omitempty"`
		Gender     int    `json:"gender,omitempty"`
		UserStatus int    `json:"user_status,omitempty"`
	}
)
