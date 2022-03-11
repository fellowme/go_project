package brand_service

import (
	"context"
	"go.uber.org/zap"
	"go_project/app/brand/brand_dao"
	"go_project/app/brand/brand_param"
	"go_project/app/brand/brand_remote_service/remote_rpc"
	"strconv"
	"strings"
)

type BrandServiceInterface interface {
	GetBrandListServiceByParam(ctx context.Context, param brand_param.GetBrandRequestParam) (brand_param.BrandListResponse, error)
	PostBrandServiceByParam(param brand_param.PostBrandRequestParam) error
	PatchBrandServiceByParam(param brand_param.PostBrandRequestParam) error
	DeleteBrandServiceById(id int) error
	GetBrandListByIdsServiceByParam(ctx context.Context, param brand_param.GetBrandByIdsRequestParam) ([]brand_param.BrandResponse, error)
	DeleteBrandListByIdsServiceByParam(param brand_param.DeleteBrandByIdsRequestParam) error
}

type BrandService struct {
	dao brand_dao.BrandDaoInterface
}

func GetBrandService() *BrandService {
	return &BrandService{dao: brand_dao.GetBrandDao()}
}
func (s BrandService) GetBrandListServiceByParam(ctx context.Context, param brand_param.GetBrandRequestParam) (brand_param.BrandListResponse, error) {
	total, data, err := s.dao.GetBrandListDaoByParam(param)
	imageIdList := make([]int, 0)
	list := make([]brand_param.BrandResponse, 0)
	for _, item := range data {
		imageIdList = append(imageIdList, item.BrandImageId)
	}
	imageMap, err := remote_rpc.GetImageListByImageIdsMap(ctx, imageIdList)
	if err != nil {
		return brand_param.BrandListResponse{}, err
	}
	for _, item := range data {
		imageInfo, ok := imageMap[item.BrandImageId]
		if ok {
			item.BrandImageUrl = imageInfo.ImageUrl
		}
		list = append(list, item)
	}
	return brand_param.BrandListResponse{
		Total: total,
		List:  list,
	}, err
}

func (s BrandService) PostBrandServiceByParam(param brand_param.PostBrandRequestParam) error {
	return s.dao.PostBrandDaoByParam(param)

}

func (s BrandService) PatchBrandServiceByParam(param brand_param.PostBrandRequestParam) error {
	return s.dao.PatchBrandDaoByParam(param)
}

func (s BrandService) DeleteBrandServiceById(id int) error {
	return s.dao.DeleteBrandDaoById(id)

}

func (s BrandService) GetBrandListByIdsServiceByParam(ctx context.Context, param brand_param.GetBrandByIdsRequestParam) ([]brand_param.BrandResponse, error) {
	idStringList := strings.Split(param.BrandIds, ",")
	idList := make([]int, 0)
	for _, value := range idStringList {
		id, err := strconv.Atoi(value)
		if err != nil {
			zap.L().Error("GetBrandListByIdsServiceByParam strconv.Atoi error", zap.Any("error", err), zap.Any("value", value))
			continue
		}
		idList = append(idList, id)
	}
	param.BrandIdList = idList
	data, err := s.dao.GetBrandListByIdsDaoByParam(param)
	imageIdList := make([]int, 0)
	list := make([]brand_param.BrandResponse, 0)
	for _, item := range data {
		imageIdList = append(imageIdList, item.BrandImageId)
	}
	imageMap, err := remote_rpc.GetImageListByImageIdsMap(ctx, imageIdList)
	if err != nil {
		return list, err
	}
	for _, item := range data {
		imageInfo, ok := imageMap[item.BrandImageId]
		if ok {
			item.BrandImageUrl = imageInfo.ImageUrl
		}
		list = append(list, item)
	}
	return list, nil
}

func (s BrandService) DeleteBrandListByIdsServiceByParam(param brand_param.DeleteBrandByIdsRequestParam) error {
	idStringList := strings.Split(param.BrandIds, ",")
	idList := make([]int, 0)
	for _, value := range idStringList {
		id, err := strconv.Atoi(value)
		if err != nil {
			zap.L().Error("DeleteBrandListByIdsServiceByParam strconv.Atoi error", zap.Any("error", err), zap.Any("value", value))
			continue
		}
		idList = append(idList, id)
	}
	param.BrandIdList = idList
	return s.dao.DeleteBrandListByIdsDaoByParam(param)
}
