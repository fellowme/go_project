package user_dao

import (
	gin_mysql "github.com/fellowme/gin_common_library/mysql"
	"go.uber.org/zap"
	"go_project/app/user/user_const"
	"go_project/app/user/user_model"
	"go_project/app/user/user_param"
	"gorm.io/gorm"
)

type UserDaoInterface interface {
	GetUserByIdDao(id int) (user user_param.UserRequestParam, err error)
	GetUserByIdsDao(idList []int) (list []user_param.UserParam, err error)
	GetUserByAccountIdsDao(accountIdList []int) (list []user_param.UserParam, err error)
	GetUserListDao(param user_param.UserListRequestParam) (list []user_param.UserParam, total int64)
	PostCreateUserDao(param user_param.UserRequestParam) (user user_param.UserRequestParam, err error)
	PatchUpdateUserDao(param user_param.UserPatchRequestParam) (user user_param.UserPatchRequestParam, err error)
	DeleteUserByIdListDao(id []int) error
}

type UserDao struct {
	dbMap map[string]*gorm.DB
}

func GetUserDao() *UserDao {
	return &UserDao{
		dbMap: gin_mysql.GetMysqlMap(),
	}
}

func (d UserDao) GetUserByIdsDao(idList []int) (list []user_param.UserParam, err error) {
	if len(idList) == 0 {
		return
	}
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, user_const.UserTableName)
	defer cancel()
	if err := tx.Where("id in (?) and user_status = 1 and is_delete = false", idList).Find(&list).Error; err != nil {
		zap.L().Error("getUserByIdsDao error", zap.Any("error", err), zap.Any("idList", idList))
	}
	return
}

func (d UserDao) GetUserListDao(param user_param.UserListRequestParam) (list []user_param.UserParam, total int64) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, user_const.UserTableName)
	defer cancel()
	tx = tx.Where("is_delete = false")
	if param.UserName != "" {
		tx = tx.Where("user_name like ?", "%"+param.UserName+"%")
	}
	if err := tx.Count(&total).Offset((param.Page - 1) * param.PageSize).Limit(param.PageSize).Order("id desc").Find(&list).Error; err != nil {
		zap.L().Error("getUserListDao Find error", zap.Any("error", err), zap.Any("param", param))
	}
	return
}

func (d UserDao) PostCreateUserDao(param user_param.UserRequestParam) (user_param.UserRequestParam, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, user_const.UserTableName)
	defer cancel()
	if err := tx.FirstOrCreate(&param, "user_name = ? AND account_id = ? and is_delete = false", param.UserName, param.AccountId).Error; err != nil {
		zap.L().Error("createUserDao FirstOrCreate error", zap.Any("error", err), zap.Any("param", param))
		return param, err
	}
	return param, nil
}

func (d UserDao) PatchUpdateUserDao(param user_param.UserPatchRequestParam) (user_param.UserPatchRequestParam, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, user_const.UserTableName)
	defer cancel()
	if err := tx.Where("id = ? and is_delete = false", param.Id).Updates(&param).Error; err != nil {
		zap.L().Error("createUserDao FirstOrCreate error", zap.Any("error", err), zap.Any("param", param))
		return param, err
	}
	return param, nil
}

func (d UserDao) GetUserByIdDao(id int) (user user_param.UserRequestParam, err error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, user_const.UserTableName)
	defer cancel()
	if err := tx.Where("id = ? and is_delete = false", id).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		zap.L().Error("createUserDao FirstOrCreate error", zap.Any("error", err), zap.Any("param", id))
	}
	return
}

func (d UserDao) DeleteUserByIdListDao(idList []int) (err error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, user_const.UserTableName)
	defer cancel()
	if err = tx.Where("is_delete = false").Delete(&user_model.User{}, idList).Error; err != nil && err != gorm.ErrRecordNotFound {
		zap.L().Error("deleteUserByIdDao delete error", zap.Any("error", err), zap.Any("idList", idList))
	}
	return

}

func (d UserDao) GetUserByAccountIdsDao(accountIdList []int) (list []user_param.UserParam, err error) {
	if len(accountIdList) == 0 {
		return
	}
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, user_const.UserTableName)
	defer cancel()
	if err := tx.Where("account_id in (?) and is_delete = false", accountIdList).Find(&list).Error; err != nil {
		zap.L().Error("GetUserByAccountIdsDao error", zap.Any("error", err), zap.Any("accountIdList", accountIdList))
	}
	return
}
