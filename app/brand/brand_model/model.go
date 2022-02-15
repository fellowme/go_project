package brand_model

import (
	gin_model "github.com/fellowme/gin_common_library/model"
	"go_project/app/brand/brand_const"
)

type Brand struct {
	gin_model.BaseMysqlStruct
	BrandName    string `json:"brand_name,omitempty" gorm:"type:varchar(50); index:brand_name; comment:品牌名称"`
	BrandImageId int    `json:"brand_image_id,omitempty" gorm:"type:int(11);comment:品牌图片id"`
	BrandDetail  string `json:"brand_detail,omitempty" gorm:"type:varchar(200);comment:品牌详情"`
	BrandWeight  int    `json:"brand_weight,omitempty" gorm:"type:int(3);comment:品牌权重"`
	BrandStatus  int    `json:"brand_status,omitempty" gorm:"type:int(1);comment:品牌状态"`
	CreateUserId int    `json:"create_user_id" gorm:"type:int(11);comment:创建者id"`
}

func (receiver Brand) TableName() string {
	return brand_const.BrandTableName
}
