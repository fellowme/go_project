package shop_dao

import (
	gin_mysql "github.com/fellowme/gin_common_library/mysql"
	"go.uber.org/zap"
	"go_project/app/shop/shop_const"
	"go_project/app/shop/shop_model"
	"go_project/app/shop/shop_param"
	"gorm.io/gorm"
)

type shopDao struct {
	dbMap map[string]*gorm.DB
}

type ShopDaoInterface interface {
	GetShopListDaoByParam(param shop_param.GetShopListRequestParam) (int64, []shop_param.ShopResponse, error)
	PostShopDaoByParam(param shop_param.PostShopRequestParam) error
	PatchShopDaoByParam(param shop_param.PatchShopRequestParam) error
	DeleteShopDaoById(id int) error
	QueryShopByIdsDaoByParam(param shop_param.GetShopByIdsRequestParam) ([]shop_param.ShopResponse, error)
	DeleteShopByIdsDaoByParam(param shop_param.DeleteShopByIdsRequestParam) error
}

func GetShopDao() *shopDao {
	return &shopDao{
		dbMap: gin_mysql.GetMysqlMap(),
	}
}

func (d *shopDao) GetShopListDaoByParam(param shop_param.GetShopListRequestParam) (int64, []shop_param.ShopResponse, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, shop_const.ShopTableName)
	defer cancel()
	if param.ShopName != "" {
		tx = tx.Where("shop_name = ?", "%"+param.ShopName+"%")
	}
	var total int64
	var data []shop_param.ShopResponse
	err := tx.Where("is_delete = false").Count(&total).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&data).Error
	if err != nil {
		zap.L().Error("shop GetShopListDaoByParam error", zap.Any("error", err), zap.Any("param", param))
	}
	return total, data, err
}

func (d *shopDao) PostShopDaoByParam(param shop_param.PostShopRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, shop_const.ShopTableName)
	defer cancel()
	err := tx.Where("shop_name = ? and is_delete = ?", param.ShopName, false).FirstOrCreate(&param).Error
	if err != nil {
		zap.L().Error("shop PostShopDaoByParam error", zap.Any("error", err), zap.Any("param", param))
	}
	return err
}

func (d shopDao) PatchShopDaoByParam(param shop_param.PatchShopRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, shop_const.ShopTableName)
	defer cancel()
	err := tx.Where("id = ? and is_delete = ?", param.Id, false).Updates(&param).Error
	if err != nil {
		zap.L().Error("shop PatchShopDaoByParam error", zap.Any("error", err), zap.Any("param", param))
	}
	return err
}

func (d shopDao) DeleteShopDaoById(id int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, shop_const.ShopTableName)
	defer cancel()
	err := tx.Where("id = ? and is_delete = ?", id, false).Delete(&shop_model.Shop{}).Error
	if err != nil {
		zap.L().Error("shop DeleteShopDaoById error", zap.Any("error", err), zap.Any("id", id))
	}
	return err
}

func (d shopDao) QueryShopByIdsDaoByParam(param shop_param.GetShopByIdsRequestParam) ([]shop_param.ShopResponse, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, shop_const.ShopTableName)
	defer cancel()
	var data []shop_param.ShopResponse
	err := tx.Where("is_delete = ? and id in (?) ", false, param.ShopIdList).Find(&data).Error
	if err != nil {
		zap.L().Error("shop QueryShopByIdsDaoByParam error", zap.Any("error", err), zap.Any("param", param))
	}
	return data, err
}

func (d shopDao) DeleteShopByIdsDaoByParam(param shop_param.DeleteShopByIdsRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, shop_const.ShopTableName)
	defer cancel()
	err := tx.Where("is_delete = ? and id in (?) ", false, param.ShopIdList).Delete(&shop_model.Shop{}).Error
	if err != nil {
		zap.L().Error("shop DeleteShopByIdsDaoByParam error", zap.Any("error", err), zap.Any("param", param))
	}
	return err
}
