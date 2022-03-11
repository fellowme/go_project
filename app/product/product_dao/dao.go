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
	PostProductMainDaoByParam(param product_param.PostProductMainExtRequestParam) error
	GetProductMainDaoById(id int) (product_param.ProductMainResponse, error)
	PatchProductMainDaoByParam(param product_param.PostProductMainExtRequestParam) error
	DeleteProductMainDaoById(id int) error
	DeleteProductMainDaoByIds(ids []int) error
	PostProductDaoByParam(param product_param.PostProductExtRequestParam) error
	GetProductListDaoByParam(param product_param.GetProductRequestParam) (int64, []product_param.ProductResponse, error)
	GetProductDaoById(id int) (product_param.ProductResponse, error)
	QueryProductImageByProductMainIds(productMainIds []int) ([]product_param.ProductImageParam, error)
	DeleteProductImageByProductMainIds(ids []int) error
	DeleteProductImageByProductIds(ids []int) error
	QueryProductImageByProductMainId(productMainId int) (product_param.ProductImageParam, error)
	QueryProductImageByProductIds(productIds []int) ([]product_param.ProductImageParam, error)
	QueryProductImageByProductId(productId int) (product_param.ProductImageParam, error)
	QueryProductStockByProductMainId(productMainId int) (int64, error)
	QueryProductStockByProductMainIds(productMainIds []int) ([]product_param.ProductMainStockParam, error)
	QueryProductStockByProductId(productId int) (int64, error)
	QueryProductStockByProductIds(productIds []int) ([]product_param.ProductMainStockParam, error)
	PatchProductDaoByParam(param product_param.PostProductExtRequestParam) error
	DeleteProductDaoById(id int) error
	DeleteProductDaoByIds(ids []int) error
	DeleteProductDaoByProductMainIds(ids []int) error
	GetProductStockListDaoByParam(param product_param.GetProductStockRequestParam) (int64, []product_param.ProductStockResponse, error)
	QueryProductMainListDaoByIds(productMainIds []int) ([]product_param.ProductMainResponse, error)
	QueryProductListDaoByProductMainIds(productMainIds []int) ([]product_param.ProductResponse, error)
	QueryProductListDaoByIds(productIds []int) ([]product_param.ProductResponse, error)
	PostProductStockDaoByParam(param product_param.PostProductStockRequestParam) error
	PatchProductStockDaoByParam(param product_param.PostProductStockRequestParam) error
	DeleteProductStockDaoById(id int) error
	DeleteProductStockDaoByProductMainIds(ids []int) error
	DeleteProductStockDaoByProductIds(ids []int) error
	DeleteProductStockByIdDaoByParam(param product_param.PostProductStockByIdsRequestParam) error
	GetProductStockByIdDaoByParam(param product_param.PostProductStockByIdsRequestParam) ([]product_param.ProductStockResponse, error)
	QueryProductListDaoByParam(param product_param.PostProductIdsRequestParam) ([]product_param.ProductResponse, error)
	QueryMainProductListDaoByProductMainIds(productMainIds []int) ([]product_param.ProductResponse, error)
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

func (d ProductDao) QueryProductImageByProductIds(productIds []int) ([]product_param.ProductImageParam, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductImageTableName)
	defer cancel()
	var data []product_param.ProductImageParam
	if err := tx.Select(" product_id AS product_id,GROUP_CONCAT( image_id ) AS image_ids").Where("product_id in (?) and is_delete = ? and product_image_type = ? ", productIds, false, product_const.ProductType).Group("product_id").Find(&data).Error; err != nil {
		zap.L().Error("QueryProductImageByProductIds find error", zap.Any("productIds", productIds), zap.Any("error", err))
		return data, err
	}
	return data, nil
}

func (d ProductDao) QueryProductImageByProductMainIds(productMainIds []int) ([]product_param.ProductImageParam, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductImageTableName)
	defer cancel()
	var data []product_param.ProductImageParam
	if err := tx.Select(" product_main_id AS product_main_id,GROUP_CONCAT( image_id ) AS image_ids").Where("product_main_id in (?) and is_delete = ? and product_image_type = ? ", productMainIds, false, product_const.ProductMainType).Group("product_main_id").Find(&data).Error; err != nil {
		zap.L().Error("QueryProductImageByProductMainIds find error", zap.Any("productMainIds", productMainIds), zap.Any("error", err))
		return data, err
	}
	return data, nil
}

func (d ProductDao) PostProductMainDaoByParam(param product_param.PostProductMainExtRequestParam) error {
	db := gin_mysql.UseMysql(d.dbMap)
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Table(product_const.ProductMainTableName).Where("title = ? and is_delete = ?", param.Title, false).FirstOrCreate(&param.PostProductMainRequestParam).Error
		if err != nil {
			zap.L().Error("PostProductMainDaoByParam ProductMainTableName FirstOrCreate", zap.Any("param", param), zap.Any("error", err))
			return err
		}
		if len(param.ImageIdList) != 0 {
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
		}
		return err
	})
	return err
}

func (d ProductDao) GetProductMainDaoById(id int) (product_param.ProductMainResponse, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductMainTableName)
	defer cancel()
	var data product_param.ProductMainResponse
	if err := tx.Where("is_delete = ? and id = ? ", false, id).First(&data).Error; err != nil {
		zap.L().Error("GetProductMainDaoById First error", zap.Any("id", id), zap.Any("error", err))
		return data, err
	}
	return data, nil
}

func (d ProductDao) QueryProductImageByProductId(productId int) (product_param.ProductImageParam, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductImageTableName)
	defer cancel()
	var data product_param.ProductImageParam
	if err := tx.Select(" product_id AS product_id, GROUP_CONCAT( image_id ORDER BY image_sort ASC) AS image_ids").Where("product_id = ? and is_delete = ? and product_image_type = ? ", productId, false, product_const.ProductType).Group("product_id").Find(&data).Error; err != nil {
		zap.L().Error("QueryProductImageByProductId find error", zap.Any("productId", productId), zap.Any("error", err))
		return data, err
	}
	return data, nil
}

func (d ProductDao) QueryProductImageByProductMainId(productMainId int) (product_param.ProductImageParam, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductImageTableName)
	defer cancel()
	var data product_param.ProductImageParam
	if err := tx.Select(" product_main_id AS product_main_id, GROUP_CONCAT( image_id ORDER BY image_sort ASC) AS image_ids").Where("product_main_id = ? and is_delete = ? and product_image_type = ? ", productMainId, false, product_const.ProductMainType).Group("product_main_id").Find(&data).Error; err != nil {
		zap.L().Error("QueryProductImageByProductMainId find error", zap.Any("productMainId", productMainId), zap.Any("error", err))
		return data, err
	}
	return data, nil
}

func (d ProductDao) PatchProductMainDaoByParam(param product_param.PostProductMainExtRequestParam) error {
	db := gin_mysql.UseMysql(d.dbMap)
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Table(product_const.ProductMainTableName).Where("id = ? and is_delete = ?", param.Id, false).Updates(&param.PostProductMainRequestParam).Error
		if err != nil {
			zap.L().Error("PatchProductMainDaoByParam ProductMainTableName Updates", zap.Any("param", param), zap.Any("error", err))
			return err
		}
		if len(param.ImageIdList) != 0 {
			productImages := make([]product_model.ProductImage, 0)
			for index, imageId := range param.ImageIdList {
				productImages = append(productImages, product_model.ProductImage{
					ProductId:        param.Id,
					ProductImageType: product_const.ProductMainType,
					ImageId:          imageId,
					ImageSort:        index,
				})
			}
			err = tx.Table(product_const.ProductImageTableName).Where("product_id = ? and is_delete = ?", param.Id, false).Delete(&product_model.ProductMain{}).Error
			if err != nil {
				zap.L().Error("PostProductMainDaoByParam ProductImageTableName Delete", zap.Any("param", param), zap.Any("error", err), zap.Any("productImages", productImages))
				return err
			}
			err = tx.Table(product_const.ProductImageTableName).Create(&productImages).Error
			if err != nil {
				zap.L().Error("PostProductMainDaoByParam ProductImageTableName Create", zap.Any("param", param), zap.Any("error", err), zap.Any("productImages", productImages))
			}
		}
		return err
	})
	return err
}

func (d ProductDao) DeleteProductMainDaoById(id int) error {
	db := gin_mysql.UseMysql(d.dbMap)
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Table(product_const.ProductMainTableName).Where("id = ? and is_delete = ?", id, false).Delete(&product_model.ProductMain{}).Error
		if err != nil {
			zap.L().Error("DeleteProductMainDaoById ProductMainTableName Delete", zap.Any("id", id), zap.Any("error", err))
			return err
		}
		err = tx.Table(product_const.ProductImageTableName).Where("product_id = ? and is_delete = ?", id, false).Delete(&product_model.ProductImage{}).Error
		if err != nil {
			zap.L().Error("DeleteProductMainDaoById ProductImageTableName Delete", zap.Any("id", id), zap.Any("error", err))
		}
		return err
	})
	return err
}

func (d ProductDao) PostProductDaoByParam(param product_param.PostProductExtRequestParam) error {
	db := gin_mysql.UseMysql(d.dbMap)
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Table(product_const.ProductTableName).Where("title = ? and is_delete = ? and product_main_id = ?", param.Title, false, param.ProductMainId).FirstOrCreate(&param.PostProductRequestParam).Error
		if err != nil {
			zap.L().Error("PostProductDaoByParam ProductTableName FirstOrCreate", zap.Any("param", param), zap.Any("error", err))
			return err
		}
		if len(param.ImageIdList) != 0 {
			productImages := make([]product_model.ProductImage, 0)
			for index, imageId := range param.ImageIdList {
				productImages = append(productImages, product_model.ProductImage{
					ProductId:        param.Id,
					ProductImageType: product_const.ProductType,
					ImageId:          imageId,
					ImageSort:        index,
				})
			}
			err = tx.Table(product_const.ProductImageTableName).Create(&productImages).Error
			if err != nil {
				zap.L().Error("PostProductMainDaoByParam ProductImageTableName Create", zap.Any("param", param), zap.Any("error", err), zap.Any("productImages", productImages))
				return err
			}
		}
		if param.Stock != 0 {
			productStock := product_model.Stock{
				ProductMainId: param.ProductMainId,
				ProductId:     param.Id,
				StockNumber:   param.Stock,
			}
			err = tx.Table(product_const.StockTableName).Where("is_delete = ? and product_id = ? ", false, param.Id).FirstOrCreate(&productStock).Error
			if err != nil {
				zap.L().Error("PostProductDaoByParam StockTableName FirstOrCreate", zap.Any("param", param), zap.Any("error", err), zap.Any("productStock", productStock))
			}
		}
		return err
	})
	return err
}

func (d ProductDao) GetProductListDaoByParam(param product_param.GetProductRequestParam) (int64, []product_param.ProductResponse, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductTableName)
	defer cancel()
	var total int64
	var data []product_param.ProductResponse
	if err := tx.Where("is_delete = ?", false).Count(&total).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&data).Error; err != nil {
		zap.L().Error("GetProductDaoByParam find error", zap.Any("param", param), zap.Any("error", err))
		return total, data, err
	}
	return total, data, nil
}

func (d ProductDao) QueryProductStockByProductMainId(productMainId int) (int64, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.StockTableName)
	defer cancel()
	var stockTotal int64
	if err := tx.Select("SUM(stock_number) AS stock_total").Where("is_delete = ? and product_main_id = ? ", false, productMainId).Find(&stockTotal).Error; err != nil {
		zap.L().Error("QueryProductStockByProductMainId find error", zap.Any("productMainId", productMainId), zap.Any("error", err))
		return stockTotal, err
	}
	return stockTotal, nil
}

func (d ProductDao) QueryProductStockByProductMainIds(productMainIds []int) ([]product_param.ProductMainStockParam, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.StockTableName)
	defer cancel()
	var productMainStock []product_param.ProductMainStockParam
	if err := tx.Select("product_main_id AS product_main_id, SUM(stock_number) AS stock_total").Where("is_delete = ? and product_main_id in (?)", false, productMainIds).Group("product_main_id").Find(&productMainStock).Error; err != nil {
		zap.L().Error("QueryProductStockByProductMainId find error", zap.Any("productMainIds", productMainIds), zap.Any("error", err))
		return productMainStock, err
	}
	return productMainStock, nil
}

func (d ProductDao) QueryProductStockByProductId(productId int) (int64, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.StockTableName)
	defer cancel()
	var stockTotal int64
	if err := tx.Select("SUM(stock_number) AS stock_total").Where("is_delete = ? and product_id = ? ", false, productId).Find(&stockTotal).Error; err != nil {
		zap.L().Error("QueryProductStockByProductId find error", zap.Any("productMainId", productId), zap.Any("error", err))
		return stockTotal, err
	}
	return stockTotal, nil
}

func (d ProductDao) QueryProductStockByProductIds(productIds []int) ([]product_param.ProductMainStockParam, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.StockTableName)
	defer cancel()
	var productMainStock []product_param.ProductMainStockParam
	if err := tx.Select("product_id AS product_id, SUM(stock_number) AS stock_total").Where("is_delete = ? and product_id in (?)", false, productIds).Group("product_id").Find(&productMainStock).Error; err != nil {
		zap.L().Error("QueryProductStockByProductIds find error", zap.Any("productIds", productIds), zap.Any("error", err))
		return productMainStock, err
	}
	return productMainStock, nil
}

func (d ProductDao) GetProductDaoById(id int) (product_param.ProductResponse, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductTableName)
	defer cancel()
	var data product_param.ProductResponse
	if err := tx.Where("is_delete = ? and id = ? ", false, id).First(&data).Error; err != nil {
		zap.L().Error("GetProductDaoById find error", zap.Any("id", id), zap.Any("error", err))
		return data, err
	}
	return data, nil
}

func (d ProductDao) PatchProductDaoByParam(param product_param.PostProductExtRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductTableName)
	defer cancel()
	if err := tx.Where("is_delete = ? and id = ? ", false, param.Id).Updates(&param).Error; err != nil {
		zap.L().Error("PatchProductDaoByParam Updates error", zap.Any("param", param), zap.Any("error", err))
		return err
	}
	return nil
}

func (d ProductDao) DeleteProductDaoById(id int) error {
	db := gin_mysql.UseMysql(d.dbMap)
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Table(product_const.ProductTableName).Where("id = ? and is_delete = ?", id, false).Delete(&product_model.Product{}).Error
		if err != nil {
			zap.L().Error("DeleteProductDaoById ProductTableName Delete", zap.Any("id", id), zap.Any("error", err))
			return err
		}
		err = tx.Table(product_const.ProductImageTableName).Where("product_id = ? and is_delete = ?", id, false).Delete(&product_model.ProductImage{}).Error
		if err != nil {
			zap.L().Error("DeleteProductDaoById ProductImageTableName Delete", zap.Any("id", id), zap.Any("error", err))
		}
		return err
	})
	return err
}

func (d ProductDao) GetProductStockListDaoByParam(param product_param.GetProductStockRequestParam) (int64, []product_param.ProductStockResponse, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.StockTableName)
	defer cancel()
	var total int64
	var data []product_param.ProductStockResponse
	if param.ProductId != 0 {
		tx = tx.Where("product_id = ? ", param.ProductId)
	}
	if param.ProductMainId != 0 {
		tx = tx.Where("product_main_id = ? ", param.ProductMainId)
	}
	if err := tx.Where("is_delete = ? ", false).Count(&total).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&data).Error; err != nil {
		zap.L().Error("GetProductStockListDaoByParam find error", zap.Any("param", param), zap.Any("error", err))
		return total, data, err
	}
	return total, data, nil
}

func (d ProductDao) QueryProductMainListDaoByIds(productMainIds []int) ([]product_param.ProductMainResponse, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductMainTableName)
	defer cancel()
	var data []product_param.ProductMainResponse
	if err := tx.Where("is_delete = ? and id in (?)", false, productMainIds).Find(&data).Error; err != nil {
		zap.L().Error("QueryProductMainListDaoByIds find error", zap.Any("productMainIds", productMainIds), zap.Any("error", err))
		return data, err
	}
	return data, nil
}

func (d ProductDao) QueryProductListDaoByIds(productIds []int) ([]product_param.ProductResponse, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductTableName)
	defer cancel()
	var data []product_param.ProductResponse
	if err := tx.Where("is_delete = ? and id in (?)", false, productIds).Find(&data).Error; err != nil {
		zap.L().Error("QueryProductListDaoByIds find error", zap.Any("productIds", productIds), zap.Any("error", err))
		return data, err
	}
	return data, nil
}

func (d ProductDao) PostProductStockDaoByParam(param product_param.PostProductStockRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.StockTableName)
	defer cancel()
	if err := tx.Where("is_delete = ? and product_id = ?", false, param.ProductId).FirstOrCreate(&param).Error; err != nil {
		zap.L().Error("PostProductStockDaoByParam find error", zap.Any("param", param), zap.Any("error", err))
		return err
	}
	return nil
}

func (d ProductDao) PatchProductStockDaoByParam(param product_param.PostProductStockRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.StockTableName)
	defer cancel()
	if err := tx.Where("is_delete = ? and id = ?", false, param.Id).Updates(&param).Error; err != nil {
		zap.L().Error("PatchProductStockDaoByParam Updates error", zap.Any("param", param), zap.Any("error", err))
		return err
	}
	return nil
}

func (d ProductDao) DeleteProductStockDaoById(id int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.StockTableName)
	defer cancel()
	if err := tx.Where("is_delete = ? and id = ?", false, id).Delete(&product_model.Stock{}).Error; err != nil {
		zap.L().Error("DeleteProductStockDaoById Delete error", zap.Any("id", id), zap.Any("error", err))
		return err
	}
	return nil
}

func (d ProductDao) DeleteProductStockByIdDaoByParam(param product_param.PostProductStockByIdsRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.StockTableName)
	defer cancel()
	if len(param.IdList) != 0 {
		tx = tx.Where("id in (?)", param.IdList)
	}
	if len(param.ProductIdList) != 0 {
		tx = tx.Where("product_id in (?)", param.ProductIdList)
	}
	if len(param.ProductMainIdList) != 0 {
		tx = tx.Where("product_main_id in (?)", param.ProductMainIdList)
	}
	if err := tx.Where("is_delete = ? ", false).Delete(&product_model.Stock{}).Error; err != nil {
		zap.L().Error("DeleteProductStockDaoById Delete error", zap.Any("param", param), zap.Any("error", err))
		return err
	}
	return nil
}

func (d ProductDao) GetProductStockByIdDaoByParam(param product_param.PostProductStockByIdsRequestParam) ([]product_param.ProductStockResponse, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.StockTableName)
	defer cancel()
	if len(param.IdList) != 0 {
		tx = tx.Where("id in (?)", param.IdList)
	}
	if len(param.ProductIdList) != 0 {
		tx = tx.Where("product_id in (?)", param.ProductIdList)
	}
	if len(param.ProductMainIdList) != 0 {
		tx = tx.Where("product_main_id in (?)", param.ProductMainIdList)
	}
	var data []product_param.ProductStockResponse
	if err := tx.Where("is_delete = ? ", false).Find(&data).Error; err != nil {
		zap.L().Error("GetProductStockByIdDaoById find error", zap.Any("param", param), zap.Any("error", err))
		return data, err
	}
	return data, nil
}

func (d ProductDao) QueryProductListDaoByParam(param product_param.PostProductIdsRequestParam) ([]product_param.ProductResponse, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductTableName)
	defer cancel()
	if len(param.IdList) != 0 {
		tx = tx.Where("id in (?)", param.IdList)
	}
	if len(param.ProductMainIdList) != 0 {
		tx = tx.Where("product_main_id in (?)", param.ProductMainIdList)
	}
	var data []product_param.ProductResponse
	if err := tx.Where("is_delete = ? ", false).Find(&data).Error; err != nil {
		zap.L().Error("QueryProductListDaoByParam find error", zap.Any("param", param), zap.Any("error", err))
		return data, err
	}
	return data, nil
}

func (d ProductDao) DeleteProductMainDaoByIds(ids []int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductMainTableName)
	defer cancel()
	if err := tx.Where("is_delete = ? AND id in (?)", false, ids).Delete(&product_model.ProductMain{}).Error; err != nil {
		zap.L().Error("DeleteProductMainDaoByIds delete error", zap.Any("ids", ids), zap.Any("error", err))
		return err
	}
	return nil
}

func (d ProductDao) DeleteProductDaoByProductMainIds(ids []int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductTableName)
	defer cancel()
	if err := tx.Where("is_delete = ? AND product_main_id in (?)", false, ids).Delete(&product_model.Product{}).Error; err != nil {
		zap.L().Error("DeleteProductDaoByProductMainId delete error", zap.Any("ids", ids), zap.Any("error", err))
		return err
	}
	return nil
}

func (d ProductDao) DeleteProductStockDaoByProductMainIds(ids []int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.StockTableName)
	defer cancel()
	if err := tx.Where("is_delete = ? AND product_main_id in (?)", false, ids).Delete(&product_model.Stock{}).Error; err != nil {
		zap.L().Error("DeleteProductStockDaoByProductMainId delete error", zap.Any("ids", ids), zap.Any("error", err))
		return err
	}
	return nil
}

func (d ProductDao) DeleteProductImageByProductMainIds(ids []int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductImageTableName)
	defer cancel()
	if err := tx.Where("is_delete = ? AND product_id in (?) AND product_image_type = ? ", false, ids, product_const.ProductMainType).Delete(&product_model.ProductImage{}).Error; err != nil {
		zap.L().Error("DeleteProductImageByProductMainId delete error", zap.Any("ids", ids), zap.Any("error", err))
		return err
	}
	return nil
}

func (d ProductDao) DeleteProductImageByProductIds(ids []int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductImageTableName)
	defer cancel()
	if err := tx.Where("is_delete = ? AND product_id in (?) AND product_image_type = ? ", false, ids, product_const.ProductType).Delete(&product_model.ProductImage{}).Error; err != nil {
		zap.L().Error("DeleteProductImageByProductId delete error", zap.Any("ids", ids), zap.Any("error", err))
		return err
	}
	return nil
}

func (d ProductDao) QueryProductListDaoByProductMainIds(productMainIds []int) ([]product_param.ProductResponse, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductTableName)
	defer cancel()
	var data []product_param.ProductResponse
	if err := tx.Where("is_delete = ? AND product_main_id in (?) ", false, productMainIds).Find(&data).Error; err != nil {
		zap.L().Error("QueryProductListDaoByProductMainIds find error", zap.Any("productMainIds", productMainIds), zap.Any("error", err))
		return data, err
	}
	return data, nil
}

func (d ProductDao) DeleteProductDaoByIds(ids []int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductTableName)
	defer cancel()
	if err := tx.Where("is_delete = ? AND id in (?) ", false, ids).Delete(&product_model.Product{}).Error; err != nil {
		zap.L().Error("DeleteProductDaoByIds delete error", zap.Any("ids", ids), zap.Any("error", err))
		return err
	}
	return nil
}

func (d ProductDao) DeleteProductStockDaoByProductIds(ids []int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.StockTableName)
	defer cancel()
	if err := tx.Where("is_delete = ? AND product_id in (?)", false, ids).Delete(&product_model.Stock{}).Error; err != nil {
		zap.L().Error("DeleteProductStockDaoByProductId delete error", zap.Any("ids", ids), zap.Any("error", err))
		return err
	}
	return nil
}

func (d ProductDao) QueryMainProductListDaoByProductMainIds(productMainIds []int) ([]product_param.ProductResponse, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, product_const.ProductTableName)
	defer cancel()
	var data []product_param.ProductResponse
	if err := tx.Where("is_delete = ? AND product_main_id in (?) and is_main_product = ? ", false, productMainIds, true).Find(&data).Error; err != nil {
		zap.L().Error("QueryMainProductListDaoByProductMainIds find error", zap.Any("productMainIds", productMainIds), zap.Any("error", err))
		return data, err
	}
	return data, nil
}
