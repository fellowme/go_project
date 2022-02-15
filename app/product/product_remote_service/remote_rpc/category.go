package remote_rpc

import (
	"context"
	"errors"
	gin_grpc "github.com/fellowme/gin_common_library/grpc"
	gin_util "github.com/fellowme/gin_common_library/util"
	"go.uber.org/zap"
	"go_project/app/product/product_const"
	service "go_project/rpc_service"
	"strconv"
	"strings"
	"time"
)

func getCategoryListByCategoryIds(ctx context.Context, categoryIdList []int) (*service.CategoryListResponse, error) {
	categoryIdSet := gin_util.RemoveRepetitionIntSlice(categoryIdList)
	if len(categoryIdSet) == 0 {
		return nil, errors.New(product_const.CategoryIdEmptyTip)
	}
	clientDeadline := time.Now().Add(5 * time.Second)
	contextDeadline, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()
	conn := gin_grpc.GetGRPCConnect(contextDeadline, "127.0.0.1:18089")
	defer gin_grpc.CloseGRPCConnect(conn)
	serviceClient := service.NewCategoryServiceClient(conn)
	categoryIdStringList := make([]string, 0)
	for _, categoryId := range categoryIdSet {
		categoryIdStringList = append(categoryIdStringList, strconv.Itoa(categoryId))
	}
	resp, err := serviceClient.GetCategoryByIds(contextDeadline, &service.CategoryRequest{
		CategoryIds: strings.Join(categoryIdStringList, ","),
	})
	if err != nil {
		zap.L().Error("grpc getCategoryListByCategoryId GetCategoryByIds error", zap.Any("error", err), zap.Any("categoryIdList", categoryIdList))
		return nil, err
	}
	return resp, nil
}

func GetCategoryListByCategoryIdsChannel(ctx context.Context, categoryIdList []int, categoryChan chan map[int]string) {
	defer close(categoryChan)
	data := make(map[int]string, 0)
	resp, err := getCategoryListByCategoryIds(ctx, categoryIdList)
	if err != nil {
		categoryChan <- data
		return
	}
	for _, item := range resp.CategoryList {
		data[int(item.Id)] = item.CategoryName
	}
	categoryChan <- data
	return
}
