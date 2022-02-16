package shop_control

import (
	"context"
	gin_translator "github.com/fellowme/gin_common_library/translator"
	gin_util "github.com/fellowme/gin_common_library/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_project/app/shop/shop_param"
	"go_project/app/shop/shop_service"
	"net/http"
	"strconv"
)

type shopControl struct {
	service shop_service.ShopServiceInterface
}

func GetShopControl() *shopControl {
	return &shopControl{
		service: shop_service.GetShopService(),
	}
}

func (s *shopControl) GetShopList(c *gin.Context) {
	var param shop_param.GetShopListRequestParam
	if err := c.ShouldBindQuery(&param); err != nil {
		zap.L().Error("shop GetShopList error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("GetShopList not get tracerContext")
	}
	data, err := s.service.GetShopListServiceByParam(ctx.(context.Context), param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SearchSuccessTip, data, c)
}

func (s *shopControl) PostShop(c *gin.Context) {
	var param shop_param.PostShopRequestParam
	if err := c.ShouldBind(&param); err != nil {
		zap.L().Error("shop PostShop error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	err := s.service.PostShopServiceByParam(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (s *shopControl) PatchShop(c *gin.Context) {
	var param shop_param.PatchShopRequestParam
	if err := c.ShouldBind(&param); err != nil {
		zap.L().Error("shop PostShop error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error("shop PatchShop strconv.Atoi error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	param.Id = id
	err = s.service.PatchShopServiceByParam(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (s *shopControl) DeleteShop(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error("shop PatchShop strconv.Atoi error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	err = s.service.DeleteShopServiceById(id)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (s *shopControl) GetShopByIds(c *gin.Context) {
	var param shop_param.GetShopByIdsRequestParam
	if err := c.ShouldBindQuery(&param); err != nil {
		zap.L().Error("shop GetShopByIds error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("GetShopByIds not get tracerContext")
	}
	data, err := s.service.GetShopByIdsServiceByParam(ctx.(context.Context), param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SearchSuccessTip, data, c)
}

func (s *shopControl) DeleteShopByIds(c *gin.Context) {
	var param shop_param.DeleteShopByIdsRequestParam
	if err := c.ShouldBind(&param); err != nil {
		zap.L().Error("shop DeleteShopByIds error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	err := s.service.DeleteShopByIdsServiceByParam(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}
