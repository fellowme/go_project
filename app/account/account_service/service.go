package account_service

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	gin_const "github.com/fellowme/gin_common_library/const"
	gin_jwt "github.com/fellowme/gin_common_library/jwt"
	gin_model "github.com/fellowme/gin_common_library/model"
	gin_util "github.com/fellowme/gin_common_library/util"
	"go_project/app/account/account_cache"
	"go_project/app/account/account_const"
	"go_project/app/account/account_dao"
	"go_project/app/account/account_model"
	"go_project/app/account/account_param"
	"go_project/app/account/account_remote_service/remote_rpc"
	"go_project/app/account/account_remote_service/remote_web"
	"time"
)

type AccountServiceInterface interface {
	GetAccountListServiceByParam(param account_param.GetAccountRequestParam) (account_param.GetAccountListResponse, error)
	PostSendCodeServiceByParam(param account_param.PostAccountRequestParam) error
	PostVerificationCodeServiceByParam(ctx context.Context, param account_param.PostPostVerificationCodeRequestParam) (string, error)
	PostLoginOutServiceByParam(ctx context.Context, param account_param.PostLoginOutRequestParam) error
	PostLoginServiceByParam(ctx context.Context, param account_param.PostLoginRequestParam) error
}

type AccountService struct {
	accountDao account_dao.AccountDaoInterface
}

func GetAccountService() AccountService {
	return AccountService{
		accountDao: account_dao.GetAccountDao(),
	}
}

func (receiver AccountService) GetAccountListServiceByParam(param account_param.GetAccountRequestParam) (account_param.GetAccountListResponse, error) {
	if param.Page == 0 {
		param.Page = gin_util.DefaultPage
	}
	if param.PageSize == 0 {
		param.PageSize = gin_util.DefaultPageSize
	}
	data, total, err := receiver.accountDao.GetAccountListDaoByParamDao(param)
	return account_param.GetAccountListResponse{
		Total: total,
		List:  data,
	}, err

}

func (receiver AccountService) PostSendCodeServiceByParam(param account_param.PostAccountRequestParam) error {
	var code string
	code = account_cache.GetPhoneRedisKeyCache(param.Mobile)
	if code == "" {
		var err error
		err, code = receiver.accountDao.PostSendCodeDaoByParamDao(param)
		if err != nil {
			return err
		}
		account_cache.SetPhoneRedisKeyCache(param.Mobile, code)
	}
	go func() {
		_ = remote_web.SendPhoneCode(code, param.Mobile)
	}()
	return nil
}

func (receiver AccountService) PostVerificationCodeServiceByParam(ctx context.Context, param account_param.PostPostVerificationCodeRequestParam) (string, error) {
	code := account_cache.GetPhoneRedisKeyCache(param.Mobile)
	if code == "" {
		return "", errors.New(account_const.VerificationCodeExpireTimeOutTip)
	}
	if code != param.VerificationMobileCode {
		return "", errors.New(account_const.VerificationCodeErrorTip)
	}
	accountInfo, err := receiver.accountDao.QueryAccountByMobileDao(param.Mobile)
	if err != nil {
		return "", err
	}
	userInfo, userError := remote_rpc.GetUserAccountById(ctx, accountInfo.Id)
	if userError != nil {
		return "", userError
	}
	menuList, menuError := remote_rpc.GetUserRoleMenuByUserId(ctx, int(userInfo.Id))
	if menuError != nil {
		return "", menuError
	}
	redisError := account_cache.SetUserRedisKeyCache(userInfo.Id, account_param.SessionUserParam{
		UserName:   userInfo.UserName,
		NickName:   userInfo.NickName,
		RealName:   userInfo.RealName,
		Gender:     userInfo.Gender,
		UserStatus: userInfo.UserStatus,
		AccountId:  userInfo.AccountId,
		Mobile:     accountInfo.Mobile,
		Email:      accountInfo.Email,
		Menu:       menuList,
	})
	if redisError != nil {
		return "", redisError
	}
	loginError := receiver.accountDao.CreateAccountLoginByParamDao(account_model.LoginTime{
		AccountId:     accountInfo.Id,
		LoginType:     1,
		LoginPlatform: param.LoginPlatform,
		LoginTime:     gin_model.LocalTime{Time: time.Now()},
	})
	_ = account_cache.DeletePhoneRedisKeyCache(param.Mobile)
	ginJwt := gin_jwt.NewJwt()
	token, _ := ginJwt.CreateJwtToken(gin_jwt.CustomClaims{
		UserId: int(userInfo.Id),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(gin_const.DefaultJwtExpiresAt).Unix(),
		},
	})
	return token, loginError
}

func (receiver AccountService) PostLoginOutServiceByParam(ctx context.Context, param account_param.PostLoginOutRequestParam) error {
	sessionUser, err := account_cache.GetUserRedisKeyCache(int32(param.UserId))
	if err != nil {
		return err
	}
	deleteError := account_cache.DeleteUserRedisKeyCache(int32(param.UserId))
	if deleteError != nil {
		return deleteError
	}
	return receiver.accountDao.CreateAccountLoginByParamDao(account_model.LoginTime{
		AccountId:     int(sessionUser.AccountId),
		LoginPlatform: param.LoginPlatform,
		LoginType:     -1,
		LoginTime:     gin_model.LocalTime{Time: time.Now()},
	})
}

func (receiver AccountService) PostLoginServiceByParam(ctx context.Context, param account_param.PostLoginRequestParam) error {
	accountInfo, err := receiver.accountDao.QueryAccountByMobileDao(param.Mobile)
	if err != nil {
		return err
	}
	if gin_util.HmacSha256(accountInfo.SaltKey, param.Password) != accountInfo.Password {
		return errors.New(account_const.PasswordErrorTip)
	}
	userInfo, userError := remote_rpc.GetUserAccountById(ctx, accountInfo.Id)
	if userError != nil {
		return userError
	}
	menuList, menuError := remote_rpc.GetUserRoleMenuByUserId(ctx, int(userInfo.Id))
	if menuError != nil {
		return menuError
	}
	redisError := account_cache.SetUserRedisKeyCache(userInfo.Id, account_param.SessionUserParam{
		UserName:   userInfo.UserName,
		NickName:   userInfo.NickName,
		RealName:   userInfo.RealName,
		Gender:     userInfo.Gender,
		UserStatus: userInfo.UserStatus,
		AccountId:  userInfo.AccountId,
		Mobile:     accountInfo.Mobile,
		Email:      accountInfo.Email,
		Menu:       menuList,
	})
	if redisError != nil {
		return redisError
	}
	loginError := receiver.accountDao.CreateAccountLoginByParamDao(account_model.LoginTime{
		AccountId:     accountInfo.Id,
		LoginType:     1,
		LoginPlatform: param.LoginPlatform,
		LoginTime:     gin_model.LocalTime{Time: time.Now()},
	})
	return loginError
}
