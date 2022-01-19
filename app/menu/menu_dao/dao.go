package menu_dao

import (
	gin_mysql "github.com/fellowme/gin_common_library/mysql"
	"go.uber.org/zap"
	"go_project/app/menu/menu_const"
	"go_project/app/menu/menu_model"
	"go_project/app/menu/menu_param"
	"gorm.io/gorm"
)

type MenuDaoInterface interface {
	PostMenuByParamDao(param menu_param.PostMenuRequestParam) error
	GetMenuListByParamDao(param menu_param.GetMenuRequestParam) ([]menu_param.MenuResponse, int64, error)
	PatchMenuByParamDao(param menu_param.PatchMenuRequestParam) error
	DeleteMenuByIdDao(id int) error
	GetMenuListByIdsDao(idList []int) ([]menu_param.MenuResponse, error)
}

type MenuDao struct {
	dbMap map[string]*gorm.DB
}

func GetMenuDao() MenuDao {
	return MenuDao{
		dbMap: gin_mysql.GetMysqlV2Map(),
	}
}

func (d MenuDao) PostMenuByParamDao(param menu_param.PostMenuRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, menu_const.MenuTableName)
	defer cancel()
	if err := tx.Where("menu_path = ? and is_delete = ? ", param.MenuPath, false).FirstOrCreate(&param).Error; err != nil {
		zap.L().Error("PostMenuDao error", zap.Any("param", param), zap.Any("error", err))
		return err
	}
	return nil
}

func (d MenuDao) GetMenuListByParamDao(param menu_param.GetMenuRequestParam) ([]menu_param.MenuResponse, int64, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, menu_const.MenuTableName)
	defer cancel()
	var total int64
	var menuList []menu_param.MenuResponse
	if param.MenuName != "" {
		tx = tx.Where("menu_name like ?", "%"+param.MenuName+"%")
	}
	if param.MenuPath != "" {
		tx = tx.Where("menu_path like ?", "%"+param.MenuPath+"%")
	}
	if param.MenuType != 0 {
		tx = tx.Where("menu_type = ?", param.MenuType)
	}
	err := tx.Count(&total).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&menuList).Error
	if err != nil {
		zap.L().Error("GetMenuListDao error", zap.Any("error", err), zap.Any("param", param))
	}
	return menuList, total, err
}

func (d MenuDao) PatchMenuByParamDao(param menu_param.PatchMenuRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, menu_const.MenuTableName)
	defer cancel()
	if err := tx.Where("id = ? and is_delete = ? ", param.Id, false).Updates(&param).Error; err != nil {
		zap.L().Error("PatchMenuDao error", zap.Any("param", param), zap.Any("error", err))
		return err
	}
	return nil
}

func (d MenuDao) DeleteMenuByIdDao(id int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, menu_const.MenuTableName)
	defer cancel()
	if err := tx.Where("id = ? and is_delete = ? ", id, false).Delete(&menu_model.Menu{}).Error; err != nil {
		zap.L().Error("DeleteMenuDao error", zap.Any("id", id), zap.Any("error", err))
		return err
	}
	return nil
}

func (d MenuDao) GetMenuListByIdsDao(idList []int) ([]menu_param.MenuResponse, error) {
	var menuList []menu_param.MenuResponse
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, menu_const.MenuTableName)
	defer cancel()
	if err := tx.Where("id in (?) and is_delete = ? ", idList, false).Find(&menuList).Error; err != nil {
		zap.L().Error("DeleteMenuDao error", zap.Any("idList", idList), zap.Any("error", err))
		return menuList, err
	}
	return menuList, nil
}
