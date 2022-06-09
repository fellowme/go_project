package stock_service

import (
	"context"
	"errors"
	"go_project/app/product/product_const"
	"go_project/app/stock/stock_cache"
	"go_project/app/stock/stock_dao"
	"go_project/app/stock/stock_param"
	"go_project/app/stock/stock_remote_service/remote_rpc"
	"strconv"
	"strings"

	gin_util "github.com/fellowme/gin_common_library/util"
	"go.uber.org/zap"
)

type StockService struct {
	dao stock_dao.StockDaoInterface
}

func GetStockService() *StockService {
	return &StockService{
		dao: stock_dao.GetStockDao(),
	}
}

type StockServiceInterface interface {
	GetStockListServiceByParam(ctx context.Context, param stock_param.GetStockRequestParam) (stock_param.StockListResponse, error)
	PostStockServiceByParam(param stock_param.PostStockRequestParam) error
	PatchStockServiceByParam(param stock_param.PostStockRequestParam) error
	DeleteStockServiceById(id int) error
	DeleteStockServiceByParam(param stock_param.PostStockByIdsRequestParam) error
	GetStockServiceByParam(ctx context.Context, param stock_param.PostStockByIdsRequestParam) ([]stock_param.StockResponse, error)
	PostStockToRedisByParam(param stock_param.PostStockTorRedisByIdsRequestParam) []error
}

/*
	GetStockServiceStockListServiceByParam  获取库存
*/
func (s StockService) GetStockListServiceByParam(ctx context.Context, param stock_param.GetStockRequestParam) (stock_param.StockListResponse, error) {
	total, stockInfos, err := s.dao.QueryStockListDaoByParam(param)
	if err != nil {
		return stock_param.StockListResponse{}, err
	}
	if total == 0 {
		return stock_param.StockListResponse{}, errors.New(gin_util.NotFindTip)
	}
	productIdList := make([]int, 0)
	productMainIdList := make([]int, 0)
	for _, stock := range stockInfos {
		productIdList = append(productIdList, stock.ProductId)
		productMainIdList = append(productMainIdList, stock.ProductMainId)
	}

	productChanMap := make(chan map[int]stock_param.ProductParam, 1)
	productMainChanMap := make(chan map[int]stock_param.ProductParam, 1)
	go remote_rpc.GetProductListByProductIdsChannel(ctx, productIdList, productChanMap)
	go remote_rpc.GetProductMainListByProductMainIdsChannel(ctx, productMainIdList, productMainChanMap)
	productInfoMap := <-productChanMap
	productMainInfoMap := <-productMainChanMap
	list := make([]stock_param.StockResponse, 0)
	for _, stock := range stockInfos {
		productName := ""
		productMainName := ""
		product, ok := productInfoMap[stock.ProductId]
		if ok {
			productName = product.ProductName
		}
		productMain, ok := productMainInfoMap[stock.ProductMainId]
		if ok {
			productName = productMain.ProductMainName
		}
		list = append(list, stock_param.StockResponse{
			Id:              stock.Id,
			ProductMainId:   stock.ProductMainId,
			ProductId:       stock.ProductId,
			ProductName:     productName,
			ProductMainName: productMainName,
			StockNumber:     stock.StockNumber,
		})
	}
	return stock_param.StockListResponse{
		Total: total,
		List:  list,
	}, nil
}

/*
	PostStockServiceByParam  新建库存
*/
func (s StockService) PostStockServiceByParam(param stock_param.PostStockRequestParam) error {
	return s.dao.PostStockDaoByParam(param)
}

/*
	PatchProductStockServiceByParam  更新库存
*/
func (s StockService) PatchStockServiceByParam(param stock_param.PostStockRequestParam) error {
	return s.dao.PatchStockDaoByParam(param)
}

/*
	DeleteStockServiceById  删除库存
*/
func (s StockService) DeleteStockServiceById(id int) error {
	return s.dao.DeleteStockDaoById(id)
}

/*
	DeleteProductStockServiceByParam  批量删除库存
*/
func (s StockService) DeleteStockServiceByParam(param stock_param.PostStockByIdsRequestParam) error {
	if param.Ids != "" {
		idStringList := strings.Split(param.Ids, ",")
		for _, idString := range idStringList {
			id, err := strconv.Atoi(idString)
			if err != nil {
				zap.L().Error("DeleteProductStockByIdServiceById Ids strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
				continue
			}
			param.IdList = append(param.IdList, id)
		}
	}
	if param.ProductIds != "" {
		idStringList := strings.Split(param.ProductIds, ",")
		for _, idString := range idStringList {
			id, err := strconv.Atoi(idString)
			if err != nil {
				zap.L().Error("DeleteProductStockByIdServiceById ProductIds strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
				continue
			}
			param.ProductIdList = append(param.ProductIdList, id)
		}
	}
	if param.ProductMainIds != "" {
		idStringList := strings.Split(param.ProductMainIds, ",")
		for _, idString := range idStringList {
			id, err := strconv.Atoi(idString)
			if err != nil {
				zap.L().Error("DeleteProductStockByIdServiceById ProductMainIds strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
				continue
			}
			param.ProductMainIdList = append(param.ProductMainIdList, id)
		}
	}
	if len(param.IdList) == 0 && len(param.ProductIdList) == 0 && len(param.ProductMainIdList) == 0 {
		return errors.New(product_const.ParamEmptyTip)
	}
	return s.dao.DeleteStockDaoByParam(param)
}

/*
	GetProductStockServiceByParam  根据 id product_id  product_main_id 批量查询库存
*/
func (s StockService) GetStockServiceByParam(ctx context.Context, param stock_param.PostStockByIdsRequestParam) ([]stock_param.StockResponse, error) {
	if param.Ids != "" {
		idStringList := strings.Split(param.Ids, ",")
		for _, idString := range idStringList {
			id, err := strconv.Atoi(idString)
			if err != nil {
				zap.L().Error("GetProductStockByIdServiceById Ids strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
				continue
			}
			param.IdList = append(param.IdList, id)
		}
	}
	if param.ProductIds != "" {
		idStringList := strings.Split(param.ProductIds, ",")
		for _, idString := range idStringList {
			id, err := strconv.Atoi(idString)
			if err != nil {
				zap.L().Error("GetProductStockByIdServiceById ProductIds strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
				continue
			}
			param.ProductIdList = append(param.ProductIdList, id)
		}
	}
	if param.ProductMainIds != "" {
		idStringList := strings.Split(param.ProductMainIds, ",")
		for _, idString := range idStringList {
			id, err := strconv.Atoi(idString)
			if err != nil {
				zap.L().Error("GetProductStockByIdServiceById ProductMainIds strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
				continue
			}
			param.ProductMainIdList = append(param.ProductMainIdList, id)
		}
	}
	if len(param.IdList) == 0 && len(param.ProductIdList) == 0 && len(param.ProductMainIdList) == 0 {
		return nil, errors.New(product_const.ParamEmptyTip)
	}
	stockInfos, err := s.dao.GetStockDaoByParam(param)
	if err != nil {
		return nil, err
	}
	productIdList := make([]int, 0)
	productMainIdList := make([]int, 0)
	for _, stock := range stockInfos {
		productIdList = append(productIdList, stock.ProductId)
		productMainIdList = append(productMainIdList, stock.ProductMainId)
	}
	productChanMap := make(chan map[int]stock_param.ProductParam, 1)
	productMainChanMap := make(chan map[int]stock_param.ProductParam, 1)
	go remote_rpc.GetProductListByProductIdsChannel(ctx, productIdList, productChanMap)
	go remote_rpc.GetProductMainListByProductMainIdsChannel(ctx, productMainIdList, productMainChanMap)
	productInfoMap := <-productChanMap
	productMainInfoMap := <-productMainChanMap
	list := make([]stock_param.StockResponse, 0)
	for _, stock := range stockInfos {
		productName := ""
		productMainName := ""
		product, ok := productInfoMap[stock.ProductId]
		if ok {
			productName = product.ProductName
		}
		productMain, ok := productMainInfoMap[stock.ProductMainId]
		if ok {
			productName = productMain.ProductMainName
		}
		list = append(list, stock_param.StockResponse{
			Id:              stock.Id,
			ProductMainId:   stock.ProductMainId,
			ProductId:       stock.ProductId,
			ProductName:     productName,
			ProductMainName: productMainName,
			StockNumber:     stock.StockNumber,
		})
	}
	return list, nil
}

/*
	PostStockToRedisByParam  根据 id product_id  product_main_id 批量将库存保存到redis
*/
func (s StockService) PostStockToRedisByParam(param stock_param.PostStockTorRedisByIdsRequestParam) (list []error) {
	if param.ProductIds != "" {
		idStringList := strings.Split(param.ProductIds, ",")
		for _, idString := range idStringList {
			id, err := strconv.Atoi(idString)
			if err != nil {
				list = append(list, err)
				zap.L().Error("PostStockToRedisByParam ProductIds strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
				continue
			}
			param.ProductIdList = append(param.ProductIdList, id)
		}
	}
	if param.ProductMainIds != "" {
		idStringList := strings.Split(param.ProductMainIds, ",")
		for _, idString := range idStringList {
			id, err := strconv.Atoi(idString)
			if err != nil {
				list = append(list, err)
				zap.L().Error("PostToRedisByParam ProductMainIds strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
				continue
			}
			param.ProductMainIdList = append(param.ProductMainIdList, id)
		}
	}
	if len(param.ProductIdList) == 0 && len(param.ProductMainIdList) == 0 {

		return append(list, errors.New(product_const.ParamEmptyTip))
	}
	stockInfos, err := s.dao.QueryStockToRedisDaoByParam(param)
	if err != nil {
		return append(list, errors.New(product_const.ParamEmptyTip))
	}
	productStockMap := make(map[int]int64, 0)
	for _, stockInfo := range stockInfos {
		productStock, ok := productStockMap[stockInfo.ProductId]
		if ok {
			productStockMap[stockInfo.ProductId] = productStock + stockInfo.StockNumber
		} else {
			productStockMap[stockInfo.ProductId] = productStock
		}
	}

	return stock_cache.SetProductStockToRedis(productStockMap)
}
