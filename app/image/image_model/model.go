package image_model

import (
	gin_model "github.com/fellowme/gin_common_library/model"
	"go_project/app/image/image_const"
)

type Image struct {
	gin_model.BaseMysqlStruct
	CreateUserId    int    `json:"create_user_id" gorm:"comment:创建者id"`
	ImageUrl        string `json:"image_url,omitempty" gorm:"comment:图片地址"`
	ImageName       string `json:"image_name,omitempty" gorm:"comment:图片名字"`
	ImageUniqueName string `json:"image_unique_name,omitempty" gorm:"comment:图片唯一名字"`
	ImageSort       int    `json:"image_sort,omitempty" gorm:"comment:图片排序"`
	ImageType       int    `json:"image_type,omitempty" gorm:"type:tinyint(1);comment:图片类型 1图片 2 视频"`
	ImageHeight     string `json:"image_height,omitempty" gorm:"comment:图片高度"`
	ImageWidth      string `json:"image_width,omitempty" gorm:"comment:图片宽度"`
	ImageSize       int64  `json:"image_size,omitempty" gorm:"comment:图片大小 MB"`
}

func (i Image) TableName() string {
	return image_const.ImageTableName
}
