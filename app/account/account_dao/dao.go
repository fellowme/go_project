package account_dao

import (
	gin_model "github.com/fellowme/gin_common_library/model"
	gin_mysql "github.com/fellowme/gin_common_library/mysql"
	gin_util "github.com/fellowme/gin_common_library/util"
	"go.uber.org/zap"
	"go_project/app/account/account_const"
	"go_project/app/account/account_model"
	"go_project/app/account/account_param"
	"gorm.io/gorm"
	"time"
)

type AccountDaoInterface interface {
	GetAccountListDaoByParamDao(param account_param.GetAccountRequestParam) ([]account_param.GetAccountResponse, int64, error)
	PostSendCodeDaoByParamDao(param account_param.PostAccountRequestParam) (error, string)
	QueryAccountByMobileDao(mobile string) (account_model.Account, error)
	QueryAccountByIdDao(id int) (account_model.Account, error)
	CreateAccountLoginByParamDao(param account_model.LoginTime) error
}

type AccountDao struct {
	dbMap map[string]*gorm.DB
}

func GetAccountDao() AccountDao {
	return AccountDao{
		dbMap: gin_mysql.GetMysqlMap(),
	}
}

func (d AccountDao) GetAccountListDaoByParamDao(param account_param.GetAccountRequestParam) ([]account_param.GetAccountResponse, int64, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, account_const.AccountTableName)
	defer cancel()
	var res []account_param.GetAccountResponse
	var total int64
	err := tx.Count(&total).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&res).Error
	if err != nil {
		zap.L().Error("getAccountListDaoByParam error", zap.Any("error", err), zap.Any("param", param))
	}
	return res, total, err
}

func (d AccountDao) PostSendCodeDaoByParamDao(param account_param.PostAccountRequestParam) (error, string) {
	err := gin_mysql.UseMysql(d.dbMap).Transaction(func(tx *gorm.DB) error {
		account := account_model.Account{
			Mobile:           param.Mobile,
			RegisterPlatform: param.RegisterPlatform,
		}
		accountError := tx.Table(account_const.AccountTableName).Where("mobile = ? and is_delete = ?", param.Mobile, false).FirstOrCreate(&account).Error
		if accountError != nil {
			zap.L().Error("getAccountListDaoByParam accountTableName Create error", zap.Any("error", accountError), zap.Any("param", param))
			return accountError
		}
		m, _ := time.ParseDuration(account_const.VerificationCodeExpireKeyTimeString)
		endTime := time.Now().Add(m)
		code := gin_util.RandomString(account_const.VerificationCodeLength)
		verificationCode := account_model.VerificationMobileCode{
			AccountId:                  account.Id,
			VerificationMobileCode:     code,
			IsVerificationMobile:       false,
			VerificationMobileSendTime: gin_model.LocalTime{Time: time.Now()},
		}
		mobileCodeError := tx.Table(account_const.VerificationMobileCodeTableName).Where("account_id = ? and is_delete = ? and verification_mobile_send_time >= ?", account.Id, false, endTime).FirstOrCreate(&verificationCode).Error
		if mobileCodeError != nil {
			zap.L().Error("getAccountListDaoByParam verificationMobileCodeTableName Create error", zap.Any("error", mobileCodeError), zap.Any("param", param))
			return mobileCodeError
		}
		param.VerificationMobileCode = verificationCode.VerificationMobileCode
		return nil
	})
	return err, param.VerificationMobileCode
}

func (d AccountDao) QueryAccountByMobileDao(mobile string) (account_model.Account, error) {
	var account account_model.Account
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, account_const.AccountTableName)
	defer cancel()
	if err := tx.Where("mobile = ? and is_delete = ?", mobile, false).First(&account).Error; err != nil {
		zap.L().Error("QueryAccountExistByMobile error", zap.Any("error", err), zap.Any("mobile", mobile))
		return account, err
	}
	return account, nil
}

func (d AccountDao) CreateAccountLoginByParamDao(param account_model.LoginTime) error {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, account_const.AccountLoginTableName)
	defer cancel()
	if err := tx.Create(&param).Error; err != nil {
		zap.L().Error("CreateAccountLoginByParam error", zap.Any("error", err), zap.Any("param", param))
		return err
	}
	return nil
}

func (d AccountDao) QueryAccountByIdDao(id int) (account_model.Account, error) {
	tx, cancel := gin_mysql.GetTxWithContext(d.dbMap, nil, account_const.AccountLoginTableName)
	defer cancel()
	var account account_model.Account
	err := tx.Where("id = ? and is_delete = ? ", id, false).First(&account).Error
	return account, err
}
