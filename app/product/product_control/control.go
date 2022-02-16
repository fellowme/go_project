package product_control

import (
	"context"
	gin_translator "github.com/fellowme/gin_common_library/translator"
	gin_util "github.com/fellowme/gin_common_library/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_project/app/product/product_param"
	"go_project/app/product/product_service"
	"net/http"
)

type ProductControl struct {
	service product_service.ProductServiceInterface
}

func GetProductControl() *ProductControl {
	return &ProductControl{
		service: product_service.GetProductService(),
	}
}

func (receiver ProductControl) GetProductMainList(c *gin.Context) {
	var req product_param.GetProductMainRequestParam
	if err := c.ShouldBindQuery(&req); err != nil {
		zap.L().Error(" product GetProductMainList error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("GetProductMainList not get tracerContext")
	}
	data, err := receiver.service.GetProductMainListServiceByParam(ctx.(context.Context), req)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SearchSuccessTip, data, c)
}

func (receiver ProductControl) PostProductMain(c *gin.Context) {
	var req product_param.PostProductMainRequestParam
	if err := c.ShouldBind(&req); err != nil {
		zap.L().Error(" product PostProductMain error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	err := receiver.service.PostProductMainServiceByParam(req)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}
