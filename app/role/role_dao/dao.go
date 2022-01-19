package role_dao

import (
	"context"
	gin_mysql "github.com/fellowme/gin_common_library/mysql"
	"go.uber.org/zap"
	"go_project/app/menu/menu_model"
	"go_project/app/role/role_const"
	"go_project/app/role/role_model"
	"go_project/app/role/role_param"
	"gorm.io/gorm"
)

type RoleDaoInterface interface {
	GetRoleListByParamDao(param role_param.GetRoleListRequestParam) ([]role_param.RoleParam, int64, error)
	PostRoleByParamDao(param role_param.PostRoleRequestParam) error
	PatchRoleByParamDao(param role_param.PatchRoleRequestParam) error
	DeleteRoleByIdDao(id int) error
	QueryRoleByIdDao(id int) (role_model.Role, error)
	QueryRoleByUserIdsDao(ctx context.Context, userId []int) ([]role_model.Role, error)
	QueryRoleByIdListDao(ctx context.Context, id []int) ([]role_model.Role, error)
	GetRoleUserListByParamDao(param role_param.GetRoleUserRequestParam) ([]role_param.RoleUserParam, int64, error)
	QueryRoleMenuListByRoleIdsDao(ctx context.Context, roleIdList []int) ([]role_model.RoleMenu, error)
	PostRoleUserByParamDao(param role_param.PostRoleUserRequestParam) error
	DeleteRoleUserByIdDao(id int) error
	PostRoleMenuByParamDao(param role_param.PostRoleMenuRequestParam) error
	GetRoleMenuListByParamDao(param role_param.GetRoleMenuRequestParam) ([]role_param.RoleMenuParam, int64, error)
	DeleteRoleMenuByIdDao(id int) error
}

type RoleDao struct {
	dbMap map[string]*gorm.DB
}

func GetRoleDao() RoleDao {
	return RoleDao{
		dbMap: gin_mysql.GetMysqlV2Map(),
	}
}

func (d RoleDao) GetRoleListByParamDao(param role_param.GetRoleListRequestParam) ([]role_param.RoleParam, int64, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, role_const.RoleTableName)
	defer cancel()
	tx = tx.Where("is_delete = false")
	if param.RoleName != "" {
		tx.Where("role_name like ?", "%"+param.RoleName+"%")
	}
	var roleList []role_param.RoleParam
	var count int64
	err := tx.Count(&count).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Order("id").Find(&roleList).Error
	if err != nil {
		zap.L().Error("GetRoleListByParamDao find error ", zap.Any("error", err), zap.Any("param", param))
	}
	return roleList, count, err
}

func (d RoleDao) PostRoleByParamDao(param role_param.PostRoleRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, role_const.RoleTableName)
	defer cancel()
	err := tx.Where("role_name = ? and is_delete = ?", param.RoleName, false).FirstOrCreate(&param).Error
	if err != nil {
		zap.L().Error("PostRoleByParamDao find error ", zap.Any("error", err), zap.Any("param", param))
	}
	return err
}

func (d RoleDao) PatchRoleByParamDao(param role_param.PatchRoleRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, role_const.RoleTableName)
	defer cancel()
	err := tx.Where("id = ? and is_delete = ?", param.Id, false).Updates(&param).Error
	if err != nil {
		zap.L().Error("PatchRoleByParamDao Updates error ", zap.Any("error", err), zap.Any("param", param))
	}
	return err

}

func (d RoleDao) DeleteRoleByIdDao(id int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, role_const.RoleTableName)
	defer cancel()
	err := tx.Where("id = ? and is_delete = ?", id, false).Delete(&role_model.Role{}).Error
	if err != nil {
		zap.L().Error("DeleteRoleByIdDao delete error ", zap.Any("error", err), zap.Any("id", id))
	}
	return err
}

func (d RoleDao) QueryRoleByIdDao(id int) (role_model.Role, error) {
	var role role_model.Role
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, role_const.RoleTableName)
	defer cancel()
	err := tx.Where("id = ? and is_delete = ?", id, false).First(&role).Error
	if err != nil {
		zap.L().Error("QueryRoleByIdDao First error ", zap.Any("error", err), zap.Any("id", id))
	}
	return role, err
}

func (d RoleDao) PostRoleUserByParamDao(param role_param.PostRoleUserRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, role_const.RoleUserTableName)
	defer cancel()
	err := tx.Where("role_id = ? and user_id = ? and is_delete = ?", param.RoleId, param.UserId, false).FirstOrCreate(&param).Error
	if err != nil {
		zap.L().Error("PostRoleUserByParamDao FirstOrCreate error ", zap.Any("error", err), zap.Any("param", param))
	}
	return err
}

func (d RoleDao) DeleteRoleUserByIdDao(id int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, role_const.RoleUserTableName)
	defer cancel()
	err := tx.Where("id = ? and is_delete = ?", id, false).Delete(&role_model.RoleUser{}).Error
	if err != nil {
		zap.L().Error("DeleteRoleUserByIdDao delete error ", zap.Any("error", err), zap.Any("id", id))
	}
	return err

}

func (d RoleDao) GetRoleUserListByParamDao(param role_param.GetRoleUserRequestParam) ([]role_param.RoleUserParam, int64, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, role_const.RoleUserTableName)
	defer cancel()
	tx = tx.Where("is_delete = ?", false)
	if param.RoleId != 0 {
		tx = tx.Where("role_id = ? ", param.RoleId)
	}
	if param.UserId != 0 {
		tx = tx.Where("user_id = ? ", param.UserId)
	}
	var total int64
	var roleUserList []role_param.RoleUserParam
	err := tx.Count(&total).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&roleUserList).Error
	if err != nil {
		zap.L().Error("GetRoleUserListByParamDao find error ", zap.Any("error", err), zap.Any("param", param))
	}
	return roleUserList, total, err
}

func (d RoleDao) QueryRoleByIdListDao(ctx context.Context, idList []int) ([]role_model.Role, error) {
	var roleList []role_model.Role
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, ctx, role_const.RoleTableName)
	defer cancel()
	err := tx.Where("id in (?) and is_delete = ?", idList, false).Find(&roleList).Error
	if err != nil {
		zap.L().Error("QueryRoleByIdListDao find error ", zap.Any("error", err), zap.Any("idList", idList))
	}
	return roleList, err

}

func (d RoleDao) GetRoleMenuListByParamDao(param role_param.GetRoleMenuRequestParam) ([]role_param.RoleMenuParam, int64, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, role_const.RoleMenuTableName)
	defer cancel()
	tx = tx.Where("is_delete = ?", false)
	if param.RoleId != 0 {
		tx = tx.Where("role_id = ? ", param.RoleId)
	}
	if param.MenuId != 0 {
		tx = tx.Where("user_id = ? ", param.MenuId)
	}
	var total int64
	var roleMenuList []role_param.RoleMenuParam
	err := tx.Count(&total).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&roleMenuList).Error
	if err != nil {
		zap.L().Error("GetRoleMenuListByParamDao find error ", zap.Any("error", err), zap.Any("param", param))
	}
	return roleMenuList, total, err
}

func (d RoleDao) PostRoleMenuByParamDao(param role_param.PostRoleMenuRequestParam) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, role_const.RoleMenuTableName)
	defer cancel()
	tx = tx.Where("is_delete = ? and role_id = ? and menu_id = ? ", false, param.RoleId, param.MenuId)
	err := tx.FirstOrCreate(&param).Error
	if err != nil {
		zap.L().Error("PostRoleMenuByParamDao FirstOrCreate error ", zap.Any("error", err), zap.Any("param", param))
	}
	return err
}

func (d RoleDao) DeleteRoleMenuByIdDao(id int) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, role_const.RoleMenuTableName)
	defer cancel()
	tx = tx.Where("is_delete = ? and id = ? ", false, id)
	err := tx.Delete(&menu_model.Menu{}).Error
	if err != nil {
		zap.L().Error("DeleteRoleMenuByIdDao delete error ", zap.Any("error", err), zap.Any("id", id))
	}
	return err
}

func (d RoleDao) QueryRoleByUserIdsDao(ctx context.Context, userIdList []int) ([]role_model.Role, error) {
	var roleList []role_model.Role
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, ctx, role_const.RoleUserTableName)
	defer cancel()
	err := tx.Where("user_id in (?) and is_delete = ?", userIdList, false).Find(&roleList).Error
	if err != nil {
		zap.L().Error("QueryRoleByUserIdsDao find error ", zap.Any("error", err), zap.Any("userIdList", userIdList))
	}
	return roleList, err
}

func (d RoleDao) QueryRoleMenuListByRoleIdsDao(ctx context.Context, roleIdList []int) ([]role_model.RoleMenu, error) {
	var roleMenuList []role_model.RoleMenu
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, ctx, role_const.RoleMenuTableName)
	defer cancel()
	err := tx.Where("role_id in (?) and is_delete = ?", roleIdList, false).Find(&roleMenuList).Error
	if err != nil {
		zap.L().Error("QueryRoleMenuListByRoleIdsDao find error ", zap.Any("error", err), zap.Any("roleIdList", roleIdList))
	}
	return roleMenuList, err
}
