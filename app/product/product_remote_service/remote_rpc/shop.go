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

func getShopListByShopIds(ctx context.Context, shopIdList []int) (*service.ShopListResponse, error) {
	shopIdSet := gin_util.RemoveRepetitionIntSlice(shopIdList)
	if len(shopIdSet) == 0 {
		return nil, errors.New(product_const.ShopIdEmptyTip)
	}
	clientDeadline := time.Now().Add(5 * time.Second)
	contextDeadline, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()
	conn := gin_grpc.GetGRPCConnect(contextDeadline, "127.0.0.1:18091")
	defer gin_grpc.CloseGRPCConnect(conn)
	serviceClient := service.NewShopServiceClient(conn)
	shopIdStringList := make([]string, 0)
	for _, shopId := range shopIdSet {
		shopIdStringList = append(shopIdStringList, strconv.Itoa(shopId))
	}
	resp, err := serviceClient.GetShopByIds(contextDeadline, &service.ShopRequest{
		ShopIds: strings.Join(shopIdStringList, ","),
	})
	if err != nil {
		zap.L().Error("grpc getShopListByShopIds GetShopByIds error", zap.Any("error", err), zap.Any("shopIdList", shopIdList))
		return nil, err
	}
	return resp, nil
}

func GetShopListByShopIdsChannel(ctx context.Context, shopIdList []int, shopChan chan map[int]product_param.ShopParam) {
	defer close(shopChan)
	data := make(map[int]product_param.ShopParam, 0)
	resp, err := getShopListByShopIds(ctx, shopIdList)
	if err != nil {
		shopChan <- data
		return
	}
	for _, item := range resp.ShopList {
		data[int(item.Id)] = product_param.ShopParam{
			Id:       int(item.Id),
			ShopName: item.ShopName,
		}
	}
	shopChan <- data
}
