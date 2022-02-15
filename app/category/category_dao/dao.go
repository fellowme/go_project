package category_dao

import (
	gin_mysql "github.com/fellowme/gin_common_library/mysql"
	"go.uber.org/zap"
	"go_project/app/category/category_const"
	"go_project/app/category/category_model"
	"go_project/app/category/category_param"
	"gorm.io/gorm"
)

type CategoryDao struct {
	dbMap map[string]*gorm.DB
}

type CategoryDaoInterface interface {
	GetCategoryListByParamDao(param category_param.GetCategoryListRequestParam) (int64, []category_param.CategoryResponse, error)
	CreateCategoryByParamDao(param category_param.CategoryRequestParam) error
	QueryCategoryListByParentId(parentIdList []int) ([]category_param.CategoryParam, error)
	UpdateCategoryByParamDao(param category_param.CategoryRequestParam) error
	DeleteCategoryByIdDao(id int) error
	QueryCategoryByIdsDao(ids []int) ([]category_param.CategoryResponse, error)
}

func GetCategoryDao() CategoryDao {
	return CategoryDao{
		dbMap: gin_mysql.GetMysqlMap(),
	}
}

func (d CategoryDao) GetCategoryListByParamDao(param category_param.GetCategoryListRequestParam) (int64, []category_param.CategoryResponse, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, category_const.CategoryTableName)
	defer cancel()
	tx = tx.Where("is_delete = false")
	var categoryInfos []category_param.CategoryResponse
	var total int64
	if param.Id != 0 {
		tx = tx.Where("id = ?", param.Id)
	}
	if param.CategoryName != "" {
		tx = tx.Where("category_name like ?", "%"+param.CategoryName+"%")
	}
	if err := tx.Count(&total).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&categoryInfos).Error; err != nil {
		zap.L().Error("GetCategoryListByParamDao find error", zap.Any("param", param), zap.Any("error", err))
		return total, categoryInfos, err
	}
	return total, categoryInfos, nil
}

func (d CategoryDao) CreateCategoryByParamDao(param category_param.CategoryRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, category_const.CategoryTableName)
	defer cancel()
	tx = tx.Where("category_name = ? and category_parent_id = ? and is_delete = ? ", param.CategoryName, param.CategoryParentId, false)
	if err := tx.FirstOrCreate(&param).Error; err != nil {
		zap.L().Error("CreateCategoryByParamDao FirstOrCreate error", zap.Any("param", param), zap.Any("error", err))
		return err
	}
	return nil
}

func (d CategoryDao) QueryCategoryListByParentId(parentIdList []int) ([]category_param.CategoryParam, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, category_const.CategoryTableName)
	defer cancel()
	var data []category_param.CategoryParam
	tx = tx.Where("category_parent_id in (?) and is_delete = ? ", parentIdList, false)
	if err := tx.Find(&data).Error; err != nil {
		zap.L().Error("QueryCategoryListByParentId find error", zap.Any("parentIdList", parentIdList), zap.Any("error", err))
		return nil, err
	}
	return data, nil
}

func (d CategoryDao) UpdateCategoryByParamDao(param category_param.CategoryRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, category_const.CategoryTableName)
	defer cancel()
	tx = tx.Where("id = ? and is_delete = ? ", param.Id, false)
	if err := tx.Updates(&param).Error; err != nil {
		zap.L().Error("UpdateCategoryByParamDao Update error", zap.Any("param", param), zap.Any("error", err))
		return err
	}
	return nil
}

func (d CategoryDao) DeleteCategoryByIdDao(id int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, category_const.CategoryTableName)
	defer cancel()
	tx = tx.Where("id = ? and is_delete = ? ", id, false)
	if err := tx.Delete(&category_model.Category{}).Error; err != nil {
		zap.L().Error("DeleteCategoryByIdDao delete error", zap.Any("id", id), zap.Any("error", err))
		return err
	}
	return nil
}

func (d CategoryDao) QueryCategoryByIdsDao(ids []int) ([]category_param.CategoryResponse, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, category_const.CategoryTableName)
	defer cancel()
	tx = tx.Where("id in (?) and is_delete = ? ", ids, false)
	var data []category_param.CategoryResponse
	if err := tx.Find(&data).Error; err != nil {
		zap.L().Error("QueryCategoryByIdsDao find error", zap.Any("ids", ids), zap.Any("error", err))
		return data, err
	}
	return data, nil
}
