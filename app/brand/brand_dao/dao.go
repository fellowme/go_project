package brand_dao

import (
	gin_mysql "github.com/fellowme/gin_common_library/mysql"
	"go.uber.org/zap"
	"go_project/app/brand/brand_const"
	"go_project/app/brand/brand_model"
	"go_project/app/brand/brand_param"
	"gorm.io/gorm"
)

type BrandDao struct {
	dbMap map[string]*gorm.DB
}

type BrandDaoInterface interface {
	GetBrandListDaoByParam(param brand_param.GetBrandRequestParam) (int64, []brand_param.BrandResponse, error)
	PostBrandDaoByParam(param brand_param.PostBrandRequestParam) error
	PatchBrandDaoByParam(param brand_param.PostBrandRequestParam) error
	DeleteBrandDaoById(id int) error
	GetBrandListByIdsDaoByParam(param brand_param.GetBrandByIdsRequestParam) ([]brand_param.BrandResponse, error)
	DeleteBrandListByIdsDaoByParam(param brand_param.DeleteBrandByIdsRequestParam) error
}

func GetBrandDao() *BrandDao {
	return &BrandDao{
		dbMap: gin_mysql.GetMysqlMap(),
	}
}

func (d BrandDao) GetBrandListDaoByParam(param brand_param.GetBrandRequestParam) (int64, []brand_param.BrandResponse, error) {
	var total int64
	var data []brand_param.BrandResponse
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, brand_const.BrandTableName)
	defer cancel()
	if param.BrandName != "" {
		tx = tx.Where("brand_name like ?", "%"+param.BrandName+"%")
	}
	if err := tx.Where("is_delete=false").Count(&total).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Or("brand_weight").Find(&data).Error; err != nil {
		zap.L().Error("GetBrandListDaoByParam error", zap.Any("param", param), zap.Any("error", err))
		return total, nil, err
	}
	return total, data, nil
}

func (d BrandDao) PostBrandDaoByParam(param brand_param.PostBrandRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, brand_const.BrandTableName)
	defer cancel()
	if err := tx.Where("is_delete=false and brand_name = ? ", param.BrandName).FirstOrCreate(&param).Error; err != nil {
		zap.L().Error("PostBrandDaoByParam error", zap.Any("param", param), zap.Any("error", err))
		return err
	}
	return nil
}

func (d BrandDao) PatchBrandDaoByParam(param brand_param.PostBrandRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, brand_const.BrandTableName)
	defer cancel()
	if err := tx.Where("is_delete=false and id = ? ", param.Id).Updates(&param).Error; err != nil {
		zap.L().Error("PatchBrandDaoByParam error", zap.Any("param", param), zap.Any("error", err))
		return err
	}
	return nil
}

func (d BrandDao) DeleteBrandDaoById(id int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, brand_const.BrandTableName)
	defer cancel()
	if err := tx.Where("is_delete=false and id = ? ", id).Delete(&brand_model.Brand{}).Error; err != nil {
		zap.L().Error("DeleteBrandDaoById error", zap.Any("id", id), zap.Any("error", err))
		return err
	}
	return nil
}

func (d BrandDao) GetBrandListByIdsDaoByParam(param brand_param.GetBrandByIdsRequestParam) ([]brand_param.BrandResponse, error) {
	var data []brand_param.BrandResponse
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, brand_const.BrandTableName)
	defer cancel()
	if err := tx.Where("is_delete=false and id in (?)", param.BrandIdList).Find(&data).Error; err != nil {
		zap.L().Error("GetBrandListDaoByParam error", zap.Any("param", param), zap.Any("error", err))
		return nil, err
	}
	return data, nil
}

func (d BrandDao) DeleteBrandListByIdsDaoByParam(param brand_param.DeleteBrandByIdsRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, brand_const.BrandTableName)
	defer cancel()
	if err := tx.Where("is_delete=false and id in (?) ", param.BrandIdList).UpdateColumn("is_delete", true).Error; err != nil {
		zap.L().Error("DeleteBrandListByIdsDaoByParam error", zap.Any("param", param), zap.Any("error", err))
		return err
	}
	return nil

}
