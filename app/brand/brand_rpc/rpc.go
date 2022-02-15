package brand_rpc

import (
	"context"
	"go_project/app/brand/brand_dao"
	"go_project/app/brand/brand_param"
	"go_project/app/brand/brand_remote_service/remote_rpc"
	service "go_project/rpc_service"
	"strconv"
	"strings"
)

type BrandRpcService struct {
	dao brand_dao.BrandDaoInterface
}

func GetBrandRpcService() *BrandRpcService {
	return &BrandRpcService{
		dao: brand_dao.GetBrandDao(),
	}
}

func (s BrandRpcService) GetBrandByIds(ctx context.Context, req *service.BrandRequest) (*service.BrandListResponse, error) {
	idList := make([]int, 0)
	brandIdListString := strings.Split(req.BrandIds, ",")
	for _, idString := range brandIdListString {
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			continue
		}
		idList = append(idList, idInt)
	}
	param := brand_param.GetBrandByIdsRequestParam{
		BrandIds:    "",
		BrandIdList: idList,
	}
	brandList := make([]*service.BrandResponse, 0)
	data, err := s.dao.GetBrandListByIdsDaoByParam(param)
	if err != nil {
		return nil, err
	}
	imageMap := make(map[int]*service.ImageResponse, 0)
	if req.IsImageUrl {
		imageIdList := make([]int, 0)
		for _, item := range data {
			imageIdList = append(imageIdList, item.BrandImageId)
		}
		imageMap, _ = remote_rpc.GetImageListByImageIdsMap(ctx, imageIdList)
	}
	for _, item := range data {
		imageUrl := ""
		image, ok := imageMap[item.BrandImageId]
		if ok {
			imageUrl = image.ImageUrl
		}
		brandList = append(brandList, &service.BrandResponse{
			Id:            int32(item.Id),
			BrandName:     item.BrandName,
			BrandImageId:  int32(item.BrandImageId),
			BrandWeight:   int32(item.BrandWeight),
			BrandStatus:   int32(item.BrandStatus),
			BrandImageUrl: imageUrl,
		})
	}
	return &service.BrandListResponse{
		BrandList: brandList,
	}, nil
}
