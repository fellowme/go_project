package remote_rpc

import (
	"context"
	"errors"
	gin_grpc "github.com/fellowme/gin_common_library/grpc"
	gin_util "github.com/fellowme/gin_common_library/util"
	"go.uber.org/zap"
	"go_project/app/stock/stock_const"
	"go_project/app/stock/stock_param"
	service "go_project/rpc_service"
	"strconv"
	"strings"
	"time"
)

func getProductListByProductIds(ctx context.Context, productIdList []int) (*service.ProductListResponse, error) {
	productIdSet := gin_util.RemoveRepetitionIntSlice(productIdList)
	if len(productIdSet) == 0 {
		return nil, errors.New(stock_const.ProductIdEmptyTip)
	}
	clientDeadline := time.Now().Add(5 * time.Second)
	contextDeadline, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()
	conn := gin_grpc.GetGRPCConnect(contextDeadline, "127.0.0.1:18092")
	defer gin_grpc.CloseGRPCConnect(conn)
	serviceClient := service.NewProductServiceClient(conn)
	productIdStringList := make([]string, 0)
	for _, productId := range productIdSet {
		productIdStringList = append(productIdStringList, strconv.Itoa(productId))
	}
	resp, err := serviceClient.GetProductByIds(contextDeadline, &service.ProductRequest{
		ProductIds: strings.Join(productIdStringList, ","),
	})
	if err != nil {
		zap.L().Error("grpc getProductListByProductIds GetProductByIds error", zap.Any("error", err), zap.Any("productIdList", productIdList))
		return nil, err
	}
	return resp, nil
}

func getProductMainListByProductMainIds(ctx context.Context, productMainIdList []int) (*service.ProductMainListResponse, error) {
	productMainIdSet := gin_util.RemoveRepetitionIntSlice(productMainIdList)
	if len(productMainIdSet) == 0 {
		return nil, errors.New(stock_const.ProductMainIdEmptyTip)
	}
	clientDeadline := time.Now().Add(5 * time.Second)
	contextDeadline, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()
	conn := gin_grpc.GetGRPCConnect(contextDeadline, "127.0.0.1:18092")
	defer gin_grpc.CloseGRPCConnect(conn)
	serviceClient := service.NewProductServiceClient(conn)
	productMainIdStringList := make([]string, 0)
	for _, productMainId := range productMainIdSet {
		productMainIdStringList = append(productMainIdStringList, strconv.Itoa(productMainId))
	}
	resp, err := serviceClient.GetProductMainByIds(contextDeadline, &service.ProductMainRequest{
		ProductMainIds: strings.Join(productMainIdStringList, ","),
	})
	if err != nil {
		zap.L().Error("grpc getProductMainListByProductMainIds GetProductMainByIds error", zap.Any("error", err), zap.Any("productMainIdList", productMainIdList))
		return nil, err
	}
	return resp, nil
}

func GetProductListByProductIdsChannel(ctx context.Context, productIdList []int, productChan chan map[int]stock_param.ProductParam) {
	defer close(productChan)
	data := make(map[int]stock_param.ProductParam, 0)
	resp, err := getProductListByProductIds(ctx, productIdList)
	if err != nil {
		productChan <- data
		return
	}
	for _, item := range resp.ProductList {
		data[int(item.Id)] = stock_param.ProductParam{
			ProductMainId:   int(item.ProductMainId),
			ProductId:       int(item.Id),
			ProductName:     item.Title,
			ProductMainName: "",
		}
	}
	productChan <- data
}

func GetProductMainListByProductMainIdsChannel(ctx context.Context, productMainIdList []int, productMainChan chan map[int]stock_param.ProductParam) {
	defer close(productMainChan)
	data := make(map[int]stock_param.ProductParam, 0)
	resp, err := getProductMainListByProductMainIds(ctx, productMainIdList)
	if err != nil {
		productMainChan <- data
		return
	}
	for _, item := range resp.ProductMainList {
		data[int(item.Id)] = stock_param.ProductParam{
			ProductMainId:   int(item.Id),
			ProductId:       0,
			ProductName:     "",
			ProductMainName: item.Title,
		}
	}
	productMainChan <- data
}
