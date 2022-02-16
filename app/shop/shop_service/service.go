package shop_service

import (
	"context"
	"go.uber.org/zap"
	"go_project/app/shop/shop_dao"
	"go_project/app/shop/shop_param"
	"go_project/app/shop/shop_remote_service/remote_rpc"
	"strconv"
	"strings"
)

type ShopServiceInterface interface {
	GetShopListServiceByParam(ctx context.Context, param shop_param.GetShopListRequestParam) (shop_param.ShopListResponse, error)
	PostShopServiceByParam(param shop_param.PostShopRequestParam) error
	PatchShopServiceByParam(param shop_param.PatchShopRequestParam) error
	DeleteShopServiceById(id int) error
	GetShopByIdsServiceByParam(ctx context.Context, param shop_param.GetShopByIdsRequestParam) ([]shop_param.ShopResponse, error)
	DeleteShopByIdsServiceByParam(param shop_param.DeleteShopByIdsRequestParam) error
}

type shopService struct {
	dao shop_dao.ShopDaoInterface
}

func GetShopService() *shopService {
	return &shopService{
		dao: shop_dao.GetShopDao(),
	}
}

func (s shopService) GetShopListServiceByParam(ctx context.Context, param shop_param.GetShopListRequestParam) (shop_param.ShopListResponse, error) {
	total, data, err := s.dao.GetShopListDaoByParam(param)
	list := make([]shop_param.ShopResponse, 0)
	imageIdList := make([]int, 0)
	for _, item := range data {
		imageIdList = append(imageIdList, item.ShopImageId)
	}
	imageMap, err := remote_rpc.GetImageListByImageIdsMap(ctx, imageIdList)
	if err != nil {
		return shop_param.ShopListResponse{}, err
	}
	for _, value := range data {
		image, ok := imageMap[value.ShopImageId]
		if ok {
			value.ShopImageUrl = image.ImageUrl
		} else {
			value.ShopImageUrl = ""
		}
		list = append(list, value)
	}
	return shop_param.ShopListResponse{
		Total: total,
		List:  list,
	}, err
}

func (s shopService) PostShopServiceByParam(param shop_param.PostShopRequestParam) error {
	return s.dao.PostShopDaoByParam(param)
}

func (s shopService) PatchShopServiceByParam(param shop_param.PatchShopRequestParam) error {
	return s.dao.PatchShopDaoByParam(param)
}

func (s shopService) DeleteShopServiceById(id int) error {
	return s.dao.DeleteShopDaoById(id)

}

func (s shopService) GetShopByIdsServiceByParam(ctx context.Context, param shop_param.GetShopByIdsRequestParam) ([]shop_param.ShopResponse, error) {
	shopIdList := make([]int, 0)
	list := make([]shop_param.ShopResponse, 0)
	shopIdStringList := strings.Split(param.ShopIds, ",")
	for _, idString := range shopIdStringList {
		id, err := strconv.Atoi(idString)
		if err != nil {
			zap.L().Error(" GetShopByIdsServiceByParam strconv.Atoi error", zap.Any("error", err), zap.Any("id", idString))
			continue
		}
		shopIdList = append(shopIdList, id)
	}
	param.ShopIdList = shopIdList
	data, err := s.dao.QueryShopByIdsDaoByParam(param)
	imageIdList := make([]int, 0)
	for _, item := range data {
		imageIdList = append(imageIdList, item.ShopImageId)
	}
	imageMap, err := remote_rpc.GetImageListByImageIdsMap(ctx, imageIdList)
	if err != nil {
		return list, err
	}
	for _, value := range data {
		image, ok := imageMap[value.ShopImageId]
		if ok {
			value.ShopImageUrl = image.ImageUrl
		} else {
			value.ShopImageUrl = ""
		}
		list = append(list, value)
	}
	return list, nil
}

func (s shopService) DeleteShopByIdsServiceByParam(param shop_param.DeleteShopByIdsRequestParam) error {
	shopIdList := make([]int, 0)
	shopIdStringList := strings.Split(param.ShopIds, ",")
	for _, idString := range shopIdStringList {
		id, err := strconv.Atoi(idString)
		if err != nil {
			zap.L().Error(" DeleteShopByIdsServiceByParam strconv.Atoi error", zap.Any("error", err), zap.Any("id", idString))
			continue
		}
		shopIdList = append(shopIdList, id)
	}
	param.ShopIdList = shopIdList
	return s.dao.DeleteShopByIdsDaoByParam(param)
}
