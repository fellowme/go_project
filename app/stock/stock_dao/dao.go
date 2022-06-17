package stock_dao

import (
	"context"
	gin_mysql "github.com/fellowme/gin_common_library/mysql"
	"go.uber.org/zap"
	"go_project/app/stock/stock_const"
	"go_project/app/stock/stock_model"
	"go_project/app/stock/stock_param"
	"gorm.io/gorm"
)

type StockDao struct {
	dbMap map[string]*gorm.DB
}

func GetStockDao() StockDao {
	return StockDao{
		dbMap: gin_mysql.GetMysqlMap(),
	}
}

type StockDaoInterface interface {
	QueryStockByProductMainId(productMainId int) (int64, error)
	QueryStockByProductMainIds(productMainIds []int) ([]stock_param.ProductMainStockParam, error)
	QueryStockByProductId(productId int) (int64, error)
	QueryStockByProductIds(productIds []int) ([]stock_param.ProductMainStockParam, error)
	QueryStockListDaoByParam(param stock_param.GetStockRequestParam) (int64, []stock_param.StockResponse, error)
	PostStockDaoByParam(param stock_param.PostStockRequestParam) error
	PatchStockDaoByParam(param stock_param.PostStockRequestParam) error
	DeleteStockDaoById(id int) error
	DeleteStockDaoByProductMainIds(ids []int) error
	DeleteStockDaoByProductIds(ids []int) error
	DeleteStockDaoByParam(param stock_param.PostStockByIdsRequestParam) error
	GetStockDaoByParam(param stock_param.PostStockByIdsRequestParam) ([]stock_param.StockResponse, error)
	QueryStockToRedisDaoByParam(ctx context.Context, param stock_param.PostStockTorRedisByIdsRequestParam) ([]stock_param.StockParam, error)
}

func (d StockDao) QueryStockByProductMainId(productMainId int) (int64, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, stock_const.StockTableName)
	defer cancel()
	var stockTotal int64
	if err := tx.Select("SUM(stock_number) AS stock_total").Where("is_delete = ? and product_main_id = ? ", false, productMainId).Find(&stockTotal).Error; err != nil {
		zap.L().Error("QueryStockByProductMainId find error", zap.Any("productMainId", productMainId), zap.Any("error", err))
		return stockTotal, err
	}
	return stockTotal, nil
}

func (d StockDao) QueryStockByProductMainIds(productMainIds []int) ([]stock_param.ProductMainStockParam, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, stock_const.StockTableName)
	defer cancel()
	var productMainStock []stock_param.ProductMainStockParam
	if err := tx.Select("product_main_id AS product_main_id, SUM(stock_number) AS stock_total").Where("is_delete = ? and product_main_id in (?)", false, productMainIds).Group("product_main_id").Find(&productMainStock).Error; err != nil {
		zap.L().Error("QueryStockByProductMainIds find error", zap.Any("productMainIds", productMainIds), zap.Any("error", err))
		return productMainStock, err
	}
	return productMainStock, nil
}

func (d StockDao) QueryStockByProductId(productId int) (int64, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, stock_const.StockTableName)
	defer cancel()
	var stockTotal int64
	if err := tx.Select("SUM(stock_number) AS stock_total").Where("is_delete = ? and product_id = ? ", false, productId).Find(&stockTotal).Error; err != nil {
		zap.L().Error("QueryStockByProductId find error", zap.Any("productMainId", productId), zap.Any("error", err))
		return stockTotal, err
	}
	return stockTotal, nil
}

func (d StockDao) QueryStockByProductIds(productIds []int) ([]stock_param.ProductMainStockParam, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, stock_const.StockTableName)
	defer cancel()
	var productMainStock []stock_param.ProductMainStockParam
	if err := tx.Select("product_id AS product_id, SUM(stock_number) AS stock_total").Where("is_delete = ? and product_id in (?)", false, productIds).Group("product_id").Find(&productMainStock).Error; err != nil {
		zap.L().Error("QueryStockByProductIds find error", zap.Any("productIds", productIds), zap.Any("error", err))
		return productMainStock, err
	}
	return productMainStock, nil
}

func (d StockDao) QueryStockListDaoByParam(param stock_param.GetStockRequestParam) (int64, []stock_param.StockResponse, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, stock_const.StockTableName)
	defer cancel()
	var total int64
	var data []stock_param.StockResponse
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

func (d StockDao) PostStockDaoByParam(param stock_param.PostStockRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, stock_const.StockTableName)
	defer cancel()
	if err := tx.Where("is_delete = ? and product_id = ?", false, param.ProductId).FirstOrCreate(&param).Error; err != nil {
		zap.L().Error("PostProductStockDaoByParam find error", zap.Any("param", param), zap.Any("error", err))
		return err
	}
	return nil
}

func (d StockDao) PatchStockDaoByParam(param stock_param.PostStockRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, stock_const.StockTableName)
	defer cancel()
	if err := tx.Where("is_delete = ? and id = ?", false, param.Id).Updates(&param).Error; err != nil {
		zap.L().Error("PatchStockDaoByParam Updates error", zap.Any("param", param), zap.Any("error", err))
		return err
	}
	return nil
}

func (d StockDao) DeleteStockDaoById(id int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, stock_const.StockTableName)
	defer cancel()
	if err := tx.Where("is_delete = ? and id = ?", false, id).Delete(&stock_model.Stock{}).Error; err != nil {
		zap.L().Error("DeleteStockDaoById Delete error", zap.Any("id", id), zap.Any("error", err))
		return err
	}
	return nil
}

func (d StockDao) DeleteStockDaoByParam(param stock_param.PostStockByIdsRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, stock_const.StockTableName)
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
	if err := tx.Where("is_delete = ? ", false).Delete(&stock_model.Stock{}).Error; err != nil {
		zap.L().Error("DeleteProductStockByIdDaoByParam Delete error", zap.Any("param", param), zap.Any("error", err))
		return err
	}
	return nil
}

func (d StockDao) GetStockDaoByParam(param stock_param.PostStockByIdsRequestParam) ([]stock_param.StockResponse, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, stock_const.StockTableName)
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
	var data []stock_param.StockResponse
	if err := tx.Where("is_delete = ? ", false).Find(&data).Error; err != nil {
		zap.L().Error("GetStockDaoByParam find error", zap.Any("param", param), zap.Any("error", err))
		return data, err
	}
	return data, nil
}

func (d StockDao) DeleteStockDaoByProductMainIds(ids []int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, stock_const.StockTableName)
	defer cancel()
	if err := tx.Where("is_delete = ? AND product_main_id in (?)", false, ids).Delete(&stock_model.Stock{}).Error; err != nil {
		zap.L().Error("DeleteStockDaoByProductMainIds delete error", zap.Any("ids", ids), zap.Any("error", err))
		return err
	}
	return nil
}

func (d StockDao) DeleteStockDaoByProductIds(ids []int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, stock_const.StockTableName)
	defer cancel()
	if err := tx.Where("is_delete = ? AND product_id in (?)", false, ids).Delete(&stock_model.Stock{}).Error; err != nil {
		zap.L().Error("DeleteStockDaoByProductIds delete error", zap.Any("ids", ids), zap.Any("error", err))
		return err
	}
	return nil
}

func (d StockDao) QueryStockToRedisDaoByParam(ctx context.Context, param stock_param.PostStockTorRedisByIdsRequestParam) ([]stock_param.StockParam, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, ctx, stock_const.StockTableName)
	defer cancel()
	if len(param.ProductIdList) != 0 {
		tx = tx.Where("product_id in (?)", param.ProductIdList)
	}
	if len(param.ProductMainIdList) != 0 {
		tx = tx.Where("product_main_id in (?)", param.ProductMainIdList)
	}
	var data []stock_param.StockParam
	if err := tx.Where("is_delete = ? ", false).Find(&data).Error; err != nil {
		zap.L().Error("GetStockDaoByParam find error", zap.Any("param", param), zap.Any("error", err))
		return data, err
	}
	return data, nil
}
