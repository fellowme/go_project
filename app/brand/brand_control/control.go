package brand_control

import (
	"context"
	gin_translator "github.com/fellowme/gin_common_library/translator"
	gin_util "github.com/fellowme/gin_common_library/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_project/app/brand/brand_param"
	"go_project/app/brand/brand_service"
	"net/http"
	"strconv"
)

type BrandControl struct {
	service brand_service.BrandServiceInterface
}

func GetBrandControl() *BrandControl {
	return &BrandControl{
		service: brand_service.GetBrandService(),
	}
}

func (receiver BrandControl) GetBrandList(c *gin.Context) {
	var param brand_param.GetBrandRequestParam
	if err := c.ShouldBindQuery(&param); err != nil {
		zap.L().Error(" brand GetBrandList error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("PostVerificationCode not get tracerContext")
	}
	data, err := receiver.service.GetBrandListServiceByParam(ctx.(context.Context), param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, nil, data, c)
}

func (receiver BrandControl) PostBrand(c *gin.Context) {
	var param brand_param.PostBrandRequestParam
	if err := c.ShouldBind(&param); err != nil {
		zap.L().Error(" brand PostBrand error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	err := receiver.service.PostBrandServiceByParam(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, nil, nil, c)
}

func (receiver BrandControl) PatchBrand(c *gin.Context) {
	var param brand_param.PostBrandRequestParam
	if err := c.ShouldBind(&param); err != nil {
		zap.L().Error(" brand PatchBrand error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error(" brand strconv.Atoi error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	param.Id = id
	err = receiver.service.PatchBrandServiceByParam(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, nil, nil, c)
}

func (receiver BrandControl) DeleteBrand(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error(" brand strconv.Atoi error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	err = receiver.service.DeleteBrandServiceById(id)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, nil, nil, c)
}

func (receiver BrandControl) GetBrandByIds(c *gin.Context) {
	var param brand_param.GetBrandByIdsRequestParam
	if err := c.ShouldBindQuery(&param); err != nil {
		zap.L().Error(" brand GetBrandByIds error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("PostVerificationCode not get tracerContext")
	}
	data, err := receiver.service.GetBrandListByIdsServiceByParam(ctx.(context.Context), param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, nil, data, c)
}

func (receiver BrandControl) DeleteBrandByIds(c *gin.Context) {
	var param brand_param.DeleteBrandByIdsRequestParam
	if err := c.ShouldBind(&param); err != nil {
		zap.L().Error(" brand DeleteBrandByIds error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	err := receiver.service.DeleteBrandListByIdsServiceByParam(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, nil, nil, c)
}
