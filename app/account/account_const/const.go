package account_const

const (
	AccountTableName                = "gin_account"
	AccountLoginTableName           = "gin_account_login"
	VerificationMobileCodeTableName = "gin_verification_mobile_code"
	VerificationEmailCodeTableName  = "gin_verification_email_code"
)

const (
	VerificationCodeExpireKeyTimeString = "-1m"
	VerificationCodeExpireKeyTimeSecond = 60 * 60
	SessionExpireKeyTimeSecond          = 60 * 60 * 24
)

const VerificationCodeLength = 4

const (
	RedisKeyVersion     = "1.0"
	PhoneFormatString   = "gin-%s-%s"
	SessionFormatString = "gin-%d-%s"
)

const (
	VerificationCodeExpireTimeOutTip = "验证码已过期"
	VerificationCodeErrorTip         = "验证码错误"
	AccountNotRegisterErrorTip       = "电话未注册"
	AccountIdListNotEmptyTip         = "账户列表id不能为空"
	UserIdListNotEmptyTip            = "用户列表id不能为空"
	AccountIdNotFindTip              = "用户信息不存在"
	PasswordErrorTip                 = "密码错误"
)
