package account_control

import (
	"context"
	gin_translator "github.com/fellowme/gin_common_library/translator"
	gin_util "github.com/fellowme/gin_common_library/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_project/app/account/account_param"
	"go_project/app/account/account_service"
	"net/http"
)

type AccountControl struct {
	accountService account_service.AccountServiceInterface
}

func GetAccountControl() AccountControl {
	return AccountControl{
		accountService: account_service.GetAccountService(),
	}
}

func (receiver *AccountControl) GetAccountList(c *gin.Context) {
	var param account_param.GetAccountRequestParam
	if err := c.BindQuery(&param); err != nil {
		zap.L().Error("account getAccountList error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	data, err := receiver.accountService.GetAccountListServiceByParam(param)
	if err != nil {
		zap.L().Error("account getAccountList error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err, nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SearchSuccessTip, data, c)
}

func (receiver *AccountControl) PostSendCode(c *gin.Context) {
	var param account_param.PostAccountRequestParam
	if err := c.Bind(&param); err != nil {
		zap.L().Error("account PostSendCode Bind error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	err := receiver.accountService.PostSendCodeServiceByParam(param)
	if err != nil {
		zap.L().Error("account PostSendCode postSendCodeServiceByParam error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err, nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (receiver *AccountControl) PostVerificationCode(c *gin.Context) {
	var param account_param.PostPostVerificationCodeRequestParam
	if err := c.Bind(&param); err != nil {
		zap.L().Error("account PostVerificationCode Bind error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("PostVerificationCode not get tracerContext")
	}
	token, err := receiver.accountService.PostVerificationCodeServiceByParam(ctx.(context.Context), param)
	if err != nil {
		zap.L().Error("account PostVerificationCode PostVerificationCodeServiceByParam error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, token, c)
}

func (receiver *AccountControl) PostLoginOut(c *gin.Context) {
	var param account_param.PostLoginOutRequestParam
	if err := c.Bind(&param); err != nil {
		zap.L().Error("account PostLoginOut Bind error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("PostLoginOut not get tracerContext")
	}
	err := receiver.accountService.PostLoginOutServiceByParam(ctx.(context.Context), param)
	if err != nil {
		zap.L().Error("account PostLoginOut PostLoginOutServiceByParam error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (receiver *AccountControl) PostLogin(c *gin.Context) {
	var param account_param.PostLoginRequestParam
	if err := c.Bind(&param); err != nil {
		zap.L().Error("account PostLogin Bind error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("PostLogin not get tracerContext")
	}
	err := receiver.accountService.PostLoginServiceByParam(ctx.(context.Context), param)
	if err != nil {
		zap.L().Error("account PostLoginOut PostLoginOutServiceByParam error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}
