package role_model

import (
	gin_model "github.com/fellowme/gin_common_library/model"
	"go_project/app/role/role_const"
)

type (
	Role struct {
		gin_model.BaseMysqlStruct
		RoleName string `json:"role_name,omitempty" gorm:"type:varchar(50);comment:角色名称"`
		Remark   string `json:"remark,omitempty" gorm:"type:varchar(200);comment:说明"`
	}

	RoleUser struct {
		gin_model.BaseMysqlStruct
		RoleId int    `json:"role_id,omitempty" gorm:"comment:角色id"`
		UserId int    `json:"user_id,omitempty" gorm:"comment:用户Id"`
		Remark string `json:"remark,omitempty" gorm:"type:varchar(200);comment:说明"`
	}

	RoleMenu struct {
		gin_model.BaseMysqlStruct
		MenuId int    `json:"menu_id,omitempty" gorm:"comment:菜单id"`
		RoleId int    `json:"role_id,omitempty" gorm:"comment:角色id"`
		Remark string `json:"remark,omitempty" gorm:"type:varchar(200);comment:说明"`
	}
)

func (Role) TableName() string {
	return role_const.RoleTableName
}

func (RoleUser) TableName() string {
	return role_const.RoleUserTableName
}

func (RoleMenu) TableName() string {
	return role_const.RoleMenuTableName
}
