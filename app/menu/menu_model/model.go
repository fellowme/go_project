package menu_model

import (
	gin_model "github.com/fellowme/gin_common_library/model"
	"go_project/app/menu/menu_const"
)

type Menu struct {
	gin_model.BaseMysqlStruct
	MenuName string `json:"menu_name,omitempty" gorm:"type:varchar(50);comment:菜单名称"`
	Remark   string `json:"remark,omitempty" gorm:"type:varchar(200);comment:说明"`
	Method   string `json:"method" gorm:"type:varchar(50);comment:菜单类型"`
	Path     string `json:"path" gorm:"type:varchar(200);comment:菜单路径"`
	Handler  string `json:"handler" gorm:"type:varchar(200);comment:处理方式"`
}

func (Menu) TableName() string {
	return menu_const.MenuTableName
}
