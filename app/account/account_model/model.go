package account_model

import (
	gin_model "github.com/fellowme/gin_common_library/model"
	"go_project/app/account/account_const"
)

type Account struct {
	gin_model.BaseMysqlStruct
	Mobile           string `json:"mobile,omitempty" gorm:"type:varchar(11);index:mobile;comment:电话号码"`
	Email            string `json:"email,omitempty" gorm:"index:email;comment:邮箱"`
	Password         string `json:"password,omitempty" gorm:"comment:密码"`
	SaltKey          string `json:"salt_key,omitempty" gorm:"comment:盐值"`
	RegisterPlatform int    `json:"register_platform,omitempty" gorm:"comment:注册平台"`
}

func (Account) TableName() string {
	return account_const.AccountTableName
}

type LoginTime struct {
	gin_model.BaseMysqlStruct
	AccountId     int                 `gorm:"comment:账户id" json:"account_id,omitempty"`
	LoginType     int                 `json:"login_type" gorm:"comment:-1退出 1登录"`
	LoginPlatform int                 `json:"login_platform" gorm:"comment:上次登录平台"`
	LoginTime     gin_model.LocalTime `json:"login_time" gorm:"comment:上次登录时间"`
}

func (LoginTime) TableName() string {
	return account_const.AccountLoginTableName
}

type VerificationMobileCode struct {
	gin_model.BaseMysqlStruct
	AccountId                  int                 `gorm:"comment:账户id" json:"account_id,omitempty"`
	VerificationMobileCode     string              `gorm:"comment:验证码" json:"verification_mobile_code,omitempty"`
	VerificationMobileTime     gin_model.LocalTime `gorm:"comment:发送验证码时间" json:"verification_mobile_time"`
	IsVerificationMobile       bool                `gorm:"comment:是否验证" json:"is_verification_mobile,omitempty"`
	VerificationMobileSendTime gin_model.LocalTime `gorm:"comment:发送时间" json:"verification_mobile_send_time"`
}

func (VerificationMobileCode) TableName() string {
	return account_const.VerificationMobileCodeTableName
}

type VerificationEmailCode struct {
	gin_model.BaseMysqlStruct
	AccountId                 int                 `gorm:"comment:账户id" json:"account_id,omitempty"`
	VerificationEmailUrl      string              `gorm:"comment:验证地址" json:"verification_email_url"`
	VerificationEmailTime     gin_model.LocalTime `gorm:"comment:发送验证时间" json:"verification_email_time"`
	IsVerificationEmail       bool                `gorm:"comment:是否验证" json:"is_verification_email,omitempty"`
	VerificationEmailSendTime gin_model.LocalTime `gorm:"comment:发送时间" json:"verification_email_send_time"`
}

func (VerificationEmailCode) TableName() string {
	return account_const.VerificationEmailCodeTableName
}
