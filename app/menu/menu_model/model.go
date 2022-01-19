package menu_model

import (
	gin_model "github.com/fellowme/gin_common_library/model"
	"go_project/app/menu/menu_const"
)

type Menu struct {
	gin_model.BaseMysqlStructV2
	MenuName string `json:"menu_name,omitempty" gorm:"type:varchar(50);comment:菜单名称"`
	MenuPath string `json:"menu_path,omitempty" gorm:"type:varchar(200);comment:菜单路径"`
	MenuType int    `json:"menu_type" gorm:"comment:菜单类型 0无限制 1 get 2 post 3 delete 4 patch 5 put 6 head"`
	Remark   string `json:"remark,omitempty" gorm:"type:varchar(200);comment:说明"`
}

func (Menu) TableName() string {
	return menu_const.MenuTableName
}
