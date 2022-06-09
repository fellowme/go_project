package stock_rpc

import (
	"context"
	"errors"
	"go_project/app/stock/stock_const"
	"go_project/app/stock/stock_dao"
	"go_project/app/stock/stock_param"
	"go_project/rpc_service"
	"strconv"
	"strings"
)

type StockRpcService struct {
	dao stock_dao.StockDaoInterface
}

func GetStockRpcService() *StockRpcService {
	return &StockRpcService{
		dao: stock_dao.GetStockDao(),
	}
}

func (s StockRpcService) GetStockByIds(ctx context.Context, req *rpc_service.StockRequest) (*rpc_service.StockListResponse, error) {
	if req.ProductMainIds == "" && req.ProductIds == "" {
		return nil, errors.New(stock_const.ParamEmptyTip)
	}
	productIdList := make([]int, 0)
	if req.ProductIds != "" {
		productIdStringList := strings.Split(req.ProductIds, ",")
		for _, productIdString := range productIdStringList {
			productId, _ := strconv.Atoi(productIdString)
			productIdList = append(productIdList, productId)
		}

	}
	productMainList := make([]int, 0)
	if req.ProductMainIds != "" {
		productMainStringList := strings.Split(req.ProductMainIds, ",")
		for _, productMainIdString := range productMainStringList {
			productMainId, _ := strconv.Atoi(productMainIdString)
			productMainList = append(productMainList, productMainId)
		}
	}
	data := make([]stock_param.ProductMainStockParam, 0)
	var err error
	if len(productIdList) != 0 {
		data, err = s.dao.QueryStockByProductIds(productIdList)
	}
	if len(productMainList) != 0 {
		data, err = s.dao.QueryStockByProductMainIds(productMainList)
	}
	if err != nil {
		return nil, err
	}
	list := make([]*rpc_service.StockResponse, 0)
	for _, stockInfo := range data {
		list = append(list, &rpc_service.StockResponse{
			ProductId:     int32(stockInfo.ProductId),
			ProductMainId: int32(stockInfo.ProductMainId),
			TotalStock:    int32(stockInfo.StockTotal),
		})
	}
	return &rpc_service.StockListResponse{StockList: list}, nil
}
