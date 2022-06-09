package remote_rpc

import (
	"context"
	"errors"
	gin_grpc "github.com/fellowme/gin_common_library/grpc"
	gin_util "github.com/fellowme/gin_common_library/util"
	"go.uber.org/zap"
	"go_project/app/product/product_const"
	"go_project/app/product/product_param"
	service "go_project/rpc_service"
	"strconv"
	"strings"
	"time"
)

func getStockListByProductIds(ctx context.Context, productIdList []int) (*service.StockListResponse, error) {
	productIdSet := gin_util.RemoveRepetitionIntSlice(productIdList)
	if len(productIdSet) == 0 {
		return nil, errors.New(product_const.ProductIdEmptyTip)
	}
	clientDeadline := time.Now().Add(5 * time.Second)
	contextDeadline, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()
	conn := gin_grpc.GetGRPCConnect(contextDeadline, "127.0.0.1:18093")
	defer gin_grpc.CloseGRPCConnect(conn)
	serviceClient := service.NewStockServiceClient(conn)
	productIdStringList := make([]string, 0)
	for _, productId := range productIdSet {
		productIdStringList = append(productIdStringList, strconv.Itoa(productId))
	}
	resp, err := serviceClient.GetStockByIds(contextDeadline, &service.StockRequest{
		ProductIds: strings.Join(productIdStringList, ","),
	})
	if err != nil {
		zap.L().Error("grpc getStockListByProductIds GetStockByIds error", zap.Any("error", err), zap.Any("productIdList", productIdList))
		return nil, err
	}
	return resp, nil
}

func GetStockListByProductIdsChannel(ctx context.Context, productIdList []int, stockChan chan map[int]product_param.StockParam) {
	defer close(stockChan)
	data := make(map[int]product_param.StockParam, 0)
	resp, err := getStockListByProductIds(ctx, productIdList)
	if err != nil {
		stockChan <- data
		return
	}
	for _, item := range resp.StockList {
		id := 0
		if item.ProductId != 0 {
			id = int(item.ProductId)
		} else if item.ProductMainId != 0 {
			id = int(item.ProductMainId)
		}
		data[id] = product_param.StockParam{
			TotalStock:    int32(item.TotalStock),
			ProductId:     int(item.ProductId),
			ProductMainId: int(item.ProductMainId),
		}
	}
	stockChan <- data
}
