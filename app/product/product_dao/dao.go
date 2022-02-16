package product_dao

import (
	gin_mysql "github.com/fellowme/gin_common_library/mysql"
	"go.uber.org/zap"
	"go_project/app/product/product_const"
	"go_project/app/product/product_model"
	"go_project/app/product/product_param"
	"gorm.io/gorm"
)

type ProductDaoInterface interface {
	GetProductMainListDaoByParam(req product_param.GetProductMainRequestParam) (int64, []product_param.ProductMainResponse, error)
	QueryProductImageByProductMainIds(productMainIds []int) ([]product_param.ProductImageParam, error)
	PostProductMainDaoByParam(param product_param.PostProductMainRequestParam) error
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
	if err := tx.Where("is_delete = ?", false).Count(&total).Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&data).Error; err != nil {
		zap.L().Error("GetProductMainListDaoByParam find error", zap.Any("param", req), zap.Any("error", err))
		return total, data, err
	}
	return total, data, nil
}

func (d ProductDao) QueryProductImageByProductMainIds(productMainIds []int) ([]product_param.ProductImageParam, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductImageTableName)
	defer cancel()
	var data []product_param.ProductImageParam
	if err := tx.Select(" product_id AS product_id,GROUP_CONCAT( image_id ) AS image_ids").Where("product_id in (?) and is_delete = ? and product_image_type = ? ", productMainIds, false, product_const.ProductMainType).Group("product_id").Find(&data).Error; err != nil {
		zap.L().Error("QueryProductImageByProductMainIds find error", zap.Any("productMainIds", productMainIds), zap.Any("error", err))
		return data, err
	}
	return data, nil
}

func (d ProductDao) PostProductMainDaoByParam(param product_param.PostProductMainRequestParam) error {
	db := gin_mysql.UseMysql(d.dbMap)
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Table(product_const.ProductMainTableName).Where("title = ? and is_delete = ?", param.Title, false).FirstOrCreate(&param).Error
		if err != nil {
			zap.L().Error("PostProductMainDaoByParam ProductMainTableName FirstOrCreate", zap.Any("param", param), zap.Any("error", err))
			return err
		}
		productImages := make([]product_model.ProductImage, 0)
		for index, imageId := range param.ImageIdList {
			productImages = append(productImages, product_model.ProductImage{
				ProductId:        param.Id,
				ProductImageType: product_const.ProductMainType,
				ImageId:          imageId,
				ImageSort:        index,
			})
		}
		err = tx.Table(product_const.ProductImageTableName).Create(&productImages).Error
		if err != nil {
			zap.L().Error("PostProductMainDaoByParam ProductImageTableName Create", zap.Any("param", param), zap.Any("error", err), zap.Any("productImages", productImages))
		}
		return err
	})
	return err
}
