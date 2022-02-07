package image_dao

import (
	gin_mysql "github.com/fellowme/gin_common_library/mysql"
	"go.uber.org/zap"
	"go_project/app/image/image_const"
	"go_project/app/image/image_model"
	"go_project/app/image/image_param"
	"gorm.io/gorm"
)

type ImageDaoInterface interface {
	GetImageListByParamDao(param image_param.GetImageRequestParam) (int64, []image_param.ImageResponse, error)
	CreateImageDao(image image_model.Image) error
	GetImageByIdDao(id int) (image_param.ImageResponse, error)
	DeleteImageByIdDao(id int) error
}

type ImageDao struct {
	dbMap map[string]*gorm.DB
}

func GetImageDao() *ImageDao {
	return &ImageDao{
		dbMap: gin_mysql.GetMysqlMap(),
	}
}

func (d *ImageDao) GetImageListByParamDao(param image_param.GetImageRequestParam) (int64, []image_param.ImageResponse, error) {
	var data []image_param.ImageResponse
	var total int64
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, image_const.ImageTableName)
	defer cancel()
	if param.ImageId != 0 {
		tx = tx.Where("id = ? ", param.ImageId)
	}
	if param.ImageType != 0 {
		tx = tx.Where("image_type = ? ", param.ImageType)
	}
	if param.ImageName != "" {
		tx = tx.Where("image_type like ? ", "%"+param.ImageName+"%")
	}
	if err := tx.Where("is_delete = ? ", false).Count(&total).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&data).Error; err != nil {
		zap.L().Error("image GetImageListByParam error", zap.Any("error", err), zap.Any("param", param))
		return total, data, err
	}
	return total, data, nil
}

func (d ImageDao) CreateImageDao(image image_model.Image) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, image_const.ImageTableName)
	defer cancel()
	if err := tx.Where("image_name = ? and is_delete = ?", image.ImageName, false).FirstOrCreate(&image).Error; err != nil {
		zap.L().Error("CreateImageDao error", zap.Any("error", err), zap.Any("image", image))
		return err
	}
	return nil
}

func (d ImageDao) GetImageByIdDao(id int) (image_param.ImageResponse, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, image_const.ImageTableName)
	defer cancel()
	var data image_param.ImageResponse
	if err := tx.Where("id = ? and is_delete = ?", id, false).First(&data).Error; err != nil && err != gorm.ErrRecordNotFound {
		zap.L().Error("GetImageByIdDao error", zap.Any("error", err), zap.Any("id", id))
		return data, err
	}
	return data, nil
}

func (d ImageDao) DeleteImageByIdDao(id int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, image_const.ImageTableName)
	defer cancel()
	if err := tx.Where("id = ? and is_delete = ?", id, false).Delete(&image_model.Image{}).Error; err != nil && err != gorm.ErrRecordNotFound {
		zap.L().Error("DeleteImageByIdDao error", zap.Any("error", err), zap.Any("id", id))
		return err
	}
	return nil

}
