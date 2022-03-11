package category_model

import (
	gin_model "github.com/fellowme/gin_common_library/model"
	"go_project/app/category/category_const"
)

type Category struct {
	gin_model.BaseMysqlStruct
	CategoryName     string `json:"category_name,omitempty" gorm:"comment:类别名称"`
	CategoryParentId int    `json:"category_parent_id,omitempty" gorm:"comment:父类别名称id"`
	CategorySort     int    `json:"category_sort,omitempty" gorm:"comment:类别排序"`
	CreateUserId     int    `json:"create_user_id" gorm:"comment:创建者id"`
}

func (receiver Category) TableName() string {
	return category_const.CategoryTableName
}
