package stock_control

import (
	"context"
	"go_project/app/stock/stock_param"
	"go_project/app/stock/stock_service"
	"net/http"
	"strconv"

	gin_translator "github.com/fellowme/gin_common_library/translator"
	gin_util "github.com/fellowme/gin_common_library/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StockControl struct {
	service stock_service.StockServiceInterface
}

func GetStockControl() *StockControl {
	return &StockControl{
		service: stock_service.GetStockService(),
	}
}

func (receiver StockControl) GetStockList(c *gin.Context) {
	var req stock_param.GetStockRequestParam
	if err := c.ShouldBindQuery(&req); err != nil {
		zap.L().Error(" stock GetStockList error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("GetStockList not get tracerContext")
	}
	data, err := receiver.service.GetStockListServiceByParam(ctx.(context.Context), req)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SearchSuccessTip, data, c)
}

func (receiver StockControl) PostStock(c *gin.Context) {
	var req stock_param.PostStockRequestParam
	if err := c.ShouldBind(&req); err != nil {
		zap.L().Error(" stock PostStock error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	err := receiver.service.PostStockServiceByParam(req)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (receiver StockControl) PatchStock(c *gin.Context) {
	var req stock_param.PostStockRequestParam
	if err := c.ShouldBind(&req); err != nil {
		zap.L().Error("stock  PatchStock error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error("stock  Patch error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	req.Id = id
	err = receiver.service.PatchStockServiceByParam(req)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (receiver StockControl) DeleteStock(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error(" stock Patch error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	err = receiver.service.DeleteStockServiceById(id)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (receiver StockControl) DeleteStockById(c *gin.Context) {
	var req stock_param.PostStockByIdsRequestParam
	if err := c.ShouldBind(&req); err != nil {
		zap.L().Error(" stock DeleteStockById error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}

	err := receiver.service.DeleteStockServiceByParam(req)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (receiver StockControl) GetStockById(c *gin.Context) {
	var req stock_param.PostStockByIdsRequestParam
	if err := c.ShouldBindQuery(&req); err != nil {
		zap.L().Error(" stock GetStockById error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("GetStockById not get tracerContext")
	}
	data, err := receiver.service.GetStockServiceByParam(ctx.(context.Context), req)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SearchSuccessTip, data, c)
}

func (receiver StockControl) PostToRedis(c *gin.Context) {
	var req stock_param.PostStockTorRedisByIdsRequestParam
	if err := c.ShouldBind(&req); err != nil {
		zap.L().Error(" stock PostToRedis error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	errorList := receiver.service.PostStockToRedisByParam(req)
	if errorList != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, errorList, nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}
