package account_param

import gin_param "github.com/fellowme/gin_common_library/param"

type (
	GetAccountRequestParam struct {
		gin_param.PageRequestParam
	}
	PostAccountRequestParam struct {
		Mobile                 string `json:"mobile,omitempty" form:"mobile" binding:"required"`
		VerificationMobileCode string `json:"verification_mobile_code,omitempty" form:"verification_mobile_code"`
		RegisterPlatform       int    `json:"register_platform" form:"register_platform" binding:"required"`
	}
	PostPostVerificationCodeRequestParam struct {
		Mobile                 string `json:"mobile,omitempty" form:"mobile" binding:"required"`
		VerificationMobileCode string `json:"verification_mobile_code,omitempty" form:"verification_mobile_code" binding:"required"`
		LoginPlatform          int    `json:"login_platform" form:"login_platform" binding:"required"`
	}
	PostLoginOutRequestParam struct {
		UserId        int `json:"user_id,omitempty" form:"user_id" binding:"required"`
		LoginPlatform int `json:"login_platform" form:"login_platform" binding:"required"`
	}
	PostLoginRequestParam struct {
		Mobile        string `json:"mobile,omitempty" form:"mobile" binding:"required"`
		Password      string `json:"password" form:"password" binding:"required"`
		LoginPlatform int    `json:"login_platform" form:"login_platform" binding:"required"`
	}
)

type (
	GetAccountListResponse struct {
		Total int64                `json:"total"`
		List  []GetAccountResponse `json:"list"`
	}
	GetAccountResponse struct {
		Mobile           string `json:"mobile,omitempty" `
		Email            string `json:"email,omitempty" `
		RegisterPlatform int    `json:"register_platform,omitempty"`
	}
)

type (
	SessionUserParam struct {
		UserName   string             `json:"user_name,omitempty"`
		NickName   string             `json:"nick_name,omitempty"`
		RealName   string             `json:"real_name,omitempty"`
		Gender     string             `json:"gender,omitempty"`
		UserStatus string             `json:"user_status,omitempty"`
		AccountId  int32              `json:"account_id,omitempty"`
		Mobile     string             `json:"mobile"`
		Email      string             `json:"email"`
		Menu       []SessionMenuParam `json:"menu"`
	}
	SessionMenuParam struct {
		Id       int    `json:"id"`
		MenuName string `json:"menu_name,omitempty" `
		MenuPath string `json:"menu_path,omitempty" `
		MenuType int    `json:"menu_type"`
	}
)
