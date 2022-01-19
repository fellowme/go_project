package user_model

import (
	gin_model "github.com/fellowme/gin_common_library/model"
	"go_project/app/user/user_const"
)

type User struct {
	gin_model.BaseMysqlStructV2
	AccountId  int    `gorm:"comment:账号id" json:"account_id"`
	UserName   string `gorm:"type:varchar(50);comment:用户名称" json:"user_name,omitempty"`
	NickName   string `gorm:"type:varchar(50);comment:用户昵称" json:"nick_name,omitempty"`
	RealName   string `gorm:"type:varchar(50);comment:用户真实名字" json:"real_name,omitempty"`
	Gender     int    `gorm:"comment:用户性别 1 男 0 女 3其他" json:"gender,omitempty"`
	UserStatus int    `gorm:"comment:用户状态审核状态 1正常 0封禁" json:"user_status,omitempty"`
}

func (User) TableName() string {
	return user_const.UserTableName
}

func GetGenderMapCode(code int) string {
	genderMap := map[int]string{
		0: "未知",
		1: "女",
		2: "男",
	}
	return genderMap[code]
}

func GetUserStatusMapCode(userStatusCode int) string {
	userStatusMap := map[int]string{
		-1: "封禁",
		0:  "待审核",
		1:  "正常",
	}
	return userStatusMap[userStatusCode]
}
