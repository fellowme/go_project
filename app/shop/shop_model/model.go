package shop_model

import (
	gin_model "github.com/fellowme/gin_common_library/model"
	"go_project/app/shop/shop_const"
)

type Shop struct {
	gin_model.BaseMysqlStruct
	ShopName     string `json:"shop_name,omitempty" gorm:"type:varchar(50);NOT NULL; index:shop_name; comment:专柜名称"`
	ShopImageId  int    `json:"shop_image_id,omitempty" gorm:"type:int(11);comment:专柜图片id"`
	ShopDetail   string `json:"shop_detail,omitempty" gorm:"type:varchar(200);comment:专柜详情"`
	ShopWeight   int    `json:"shop_weight,omitempty" gorm:"type:int(4);default:100;comment:专柜权重"`
	ShopStatus   int    `json:"shop_status,omitempty" gorm:"type:int(2);default:-1;comment:专柜状态"`
	CreateUserId int    `json:"create_user_id" gorm:"type:int(11);comment:创建者id"`
}

func (receiver Shop) TableName() string {
	return shop_const.ShopTableName
}
