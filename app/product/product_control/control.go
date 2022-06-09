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
	"strconv"
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

func (receiver ProductControl) GetProductMain(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error(" product GetProductMain error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("GetProductMainList not get tracerContext")
	}
	data, err := receiver.service.GetProductMainServiceById(ctx.(context.Context), id)
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

func (receiver ProductControl) PatchProductMain(c *gin.Context) {
	var req product_param.PostProductMainRequestParam
	if err := c.ShouldBind(&req); err != nil {
		zap.L().Error(" product PatchProductMain error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error(" product PostProductMain error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	req.Id = id
	err = receiver.service.PatchProductMainServiceByParam(req)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (receiver ProductControl) DeleteProductMain(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error(" product DeleteProductMain error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	err = receiver.service.DeleteProductMainServiceById(id)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (receiver ProductControl) PostProduct(c *gin.Context) {
	var req product_param.PostProductRequestParam
	if err := c.ShouldBind(&req); err != nil {
		zap.L().Error(" product PostProduct error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	err := receiver.service.PostProductServiceByParam(req)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (receiver ProductControl) GetProductList(c *gin.Context) {
	var req product_param.GetProductRequestParam
	if err := c.ShouldBindQuery(&req); err != nil {
		zap.L().Error(" product GetProductList error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("GetProductMainList not get tracerContext")
	}
	data, err := receiver.service.GetProductServiceByParam(ctx.(context.Context), req)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SearchSuccessTip, data, c)
}

func (receiver ProductControl) GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error(" product GetProduct error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("GetProduct not get tracerContext")
	}
	data, err := receiver.service.GetProductServiceById(ctx.(context.Context), id)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SearchSuccessTip, data, c)
}

func (receiver ProductControl) PatchProduct(c *gin.Context) {
	var req product_param.PostProductRequestParam
	if err := c.ShouldBind(&req); err != nil {
		zap.L().Error(" product PatchProduct error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error(" product PatchProduct error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	req.Id = id
	err = receiver.service.PatchProductServiceByParam(req)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (receiver ProductControl) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error(" product PatchProduct error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	err = receiver.service.DeleteProductServiceById(id)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (receiver ProductControl) GetProductByProductMainIds(c *gin.Context) {
	var req product_param.PostProductIdsRequestParam
	if err := c.ShouldBind(&req); err != nil {
		zap.L().Error(" product GetProductByProductMainIds error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("GetProductByProductMainIds not get tracerContext")
	}
	data, err := receiver.service.GetProductByProductMainIdsServiceByParam(ctx.(context.Context), req)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SearchSuccessTip, data, c)
}

func (receiver ProductControl) PostDeleteProductMainAll(c *gin.Context) {
	var req product_param.PostDeleteProductMainAllRequestParam
	if err := c.ShouldBind(&req); err != nil {
		zap.L().Error(" product PostDeleteProductMainAll error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	err := receiver.service.PostDeleteProductMainAllServiceByParam(req)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (receiver ProductControl) PostDeleteProductByIds(c *gin.Context) {
	var req product_param.DeletePostProductIdsRequestParam
	if err := c.ShouldBind(&req); err != nil {
		zap.L().Error(" product PostDeleteProductByIds error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	err := receiver.service.PostDeleteProductServiceByParam(req)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (receiver ProductControl) PostProductMainToMq(c *gin.Context) {
	var req product_param.PostProductMainIdsToMqRequestParam
	if err := c.ShouldBind(&req); err != nil {
		zap.L().Error(" product PostProductMainToMq error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	messageId, err := receiver.service.PostProductMainToMqServiceByParam(req)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, messageId, c)
}

func (receiver ProductControl) PostProductToMq(c *gin.Context) {
	var req product_param.PostProductIdsToMqRequestParam
	if err := c.ShouldBind(&req); err != nil {
		zap.L().Error(" product PostProductToMq error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	messageId, err := receiver.service.PostProductToMqServiceByParam(req)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, messageId, c)
}
