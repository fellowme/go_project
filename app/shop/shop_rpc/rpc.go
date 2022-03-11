package shop_rpc

import (
	"context"
	"go.uber.org/zap"
	"go_project/app/shop/shop_dao"
	"go_project/app/shop/shop_param"
	"go_project/app/shop/shop_remote_service/remote_rpc"
	service "go_project/rpc_service"
	"strconv"
	"strings"
)

type RpcService struct {
	dao shop_dao.ShopDaoInterface
}

func GetShopRpcService() RpcService {
	return RpcService{
		dao: shop_dao.GetShopDao(),
	}
}

func (r RpcService) GetShopByIds(ctx context.Context, req *service.ShopRequest) (*service.ShopListResponse, error) {
	shopIdList := make([]int, 0)
	list := make([]*service.ShopResponse, 0)
	shopIdStringList := strings.Split(req.ShopIds, ",")
	for _, idString := range shopIdStringList {
		id, err := strconv.Atoi(idString)
		if err != nil {
			zap.L().Error(" GetShopByIds strconv.Atoi error", zap.Any("error", err), zap.Any("id", idString))
			continue
		}
		shopIdList = append(shopIdList, id)
	}
	param := shop_param.GetShopByIdsRequestParam{
		ShopIds:    "",
		ShopIdList: shopIdList,
	}
	data, err := r.dao.QueryShopByIdsDaoByParam(param)
	imageIdList := make([]int, 0)
	for _, item := range data {
		imageIdList = append(imageIdList, item.ShopImageId)
	}
	imageMap := make(map[int]*service.ImageResponse, 0)
	if req.IsImageUrl {
		imageMap, err = remote_rpc.GetImageListByImageIdsMap(ctx, imageIdList)
		if err != nil {
			return nil, err
		}
	}
	for _, value := range data {
		shopImageUrl := ""
		image, ok := imageMap[value.ShopImageId]
		if ok {
			shopImageUrl = image.ImageUrl
		}
		list = append(list, &service.ShopResponse{
			Id:           int32(value.Id),
			ShopName:     value.ShopName,
			ShopImageId:  int32(value.ShopImageId),
			ShopWeight:   int32(value.ShopWeight),
			ShopStatus:   int32(value.ShopStatus),
			ShopImageUrl: shopImageUrl,
		})
	}
	return &service.ShopListResponse{
		ShopList: list,
	}, nil
}
