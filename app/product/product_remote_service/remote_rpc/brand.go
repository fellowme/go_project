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

func getBrandListByBrandIds(ctx context.Context, brandIdList []int) (*service.BrandListResponse, error) {
	brandIdSet := gin_util.RemoveRepetitionIntSlice(brandIdList)
	if len(brandIdSet) == 0 {
		return nil, errors.New(product_const.BrandIdEmptyTip)
	}
	clientDeadline := time.Now().Add(5 * time.Second)
	contextDeadline, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()
	conn := gin_grpc.GetGRPCConnect(contextDeadline, "127.0.0.1:18090")
	defer gin_grpc.CloseGRPCConnect(conn)
	serviceClient := service.NewBrandServiceClient(conn)
	brandIdStringList := make([]string, 0)
	for _, brandId := range brandIdSet {
		brandIdStringList = append(brandIdStringList, strconv.Itoa(brandId))
	}
	resp, err := serviceClient.GetBrandByIds(contextDeadline, &service.BrandRequest{
		BrandIds: strings.Join(brandIdStringList, ","),
	})
	if err != nil {
		zap.L().Error("grpc getBrandListByBrandIds GetBrandByIds error", zap.Any("error", err), zap.Any("brandIdList", brandIdList))
		return nil, err
	}
	return resp, nil
}

func GetBrandListByBrandIdsChannel(ctx context.Context, brandIdList []int, brandChan chan map[int]product_param.BrandParam) {
	defer close(brandChan)
	data := make(map[int]product_param.BrandParam, 0)
	resp, err := getBrandListByBrandIds(ctx, brandIdList)
	if err != nil {
		brandChan <- data
		return
	}
	for _, item := range resp.BrandList {
		data[int(item.Id)] = product_param.BrandParam{
			Id:        int(item.Id),
			BrandName: item.BrandName,
		}
	}
	brandChan <- data
}
