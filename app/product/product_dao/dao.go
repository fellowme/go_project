package product_dao

import (
	gin_mysql "github.com/fellowme/gin_common_library/mysql"
	"go.uber.org/zap"
	"go_project/app/product/product_const"
	"go_project/app/product/product_param"
	"gorm.io/gorm"
)

type ProductDaoInterface interface {
	GetProductMainListDaoByParam(req product_param.GetProductMainRequestParam) (int64, []product_param.ProductMainResponse, error)
	QueryProductImageByProductMainIds(productMainIds []int) ([]product_param.ProductImageParam, error)
}
type ProductDao struct {
	dbMap map[string]*gorm.DB
}

func GetProductDao() ProductDao {
	return ProductDao{
		dbMap: gin_mysql.GetMysqlMap(),
	}
}

func (d ProductDao) GetProductMainListDaoByParam(req product_param.GetProductMainRequestParam) (int64, []product_param.ProductMainResponse, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductMainTableName)
	defer cancel()
	var total int64
	var data []product_param.ProductMainResponse
	if err := tx.Count(&total).Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&data).Error; err != nil {
		zap.L().Error("GetProductMainListDaoByParam find error", zap.Any("param", req), zap.Any("error", err))
		return total, data, err
	}
	return total, data, nil
}

func (d ProductDao) QueryProductImageByProductMainIds(productMainIds []int) ([]product_param.ProductImageParam, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductImageTableName)
	defer cancel()
	var data []product_param.ProductImageParam
	if err := tx.Select(" product_id AS product_id,GROUP_CONCAT( image_id ) AS image_ids").Where("product_id in (?) and is_delete = ? and product_image_type = ? ", productMainIds, false, 2).Group("product_id").Find(&data).Error; err != nil {
		zap.L().Error("QueryProductImageByProductMainIds find error", zap.Any("productMainIds", productMainIds), zap.Any("error", err))
		return data, err
	}
	return data, nil
}
