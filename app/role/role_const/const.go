package role_const

const (
	RoleTableName     = "gin_role"
	RoleUserTableName = "gin_role_user"
	RoleMenuTableName = "gin_role_menu"
)

const (
	UserIdListNotEmptyTip    = "用户列表id不能为空"
	UserIdNotFindTip         = "用户不存在"
	RoleIdNotFindTip         = "角色信息不存在"
	MenuIdNotFindTip         = "菜单信息不存在"
	DeleteRoleMenuNotFindTip = "删除的信息不存在"
	RoleMenuMapNotMatchTip   = "权限验证不通过"
)

const (
	RedisKeyVersion                     = "1.0"
	RoleMenuTireRedisKeyFormatString    = "gin-%s-%s"
	VerificationCodeExpireKeyTimeSecond = 60 * 60
)
