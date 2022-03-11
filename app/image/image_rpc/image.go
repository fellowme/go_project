package image_rpc

import (
	"context"
	"go_project/app/image/image_dao"
	"go_project/rpc_service"
	"strconv"
	"strings"
)

type ImageRpcService struct {
	dao image_dao.ImageDaoInterface
}

func GetImageRpcService() *ImageRpcService {
	return &ImageRpcService{
		dao: image_dao.GetImageDao(),
	}
}

func (s ImageRpcService) GetImageByIds(ctx context.Context, req *rpc_service.ImageRequest) (*rpc_service.ImageListResponse, error) {
	idList := make([]int, 0)
	imageIdListString := strings.Split(req.ImageIds, ",")
	for _, idString := range imageIdListString {
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			continue
		}
		idList = append(idList, idInt)
	}
	imageList := make([]*rpc_service.ImageResponse, 0)
	data, err := s.dao.GetImageByIdsDao(idList)
	if err != nil {
		return nil, err
	}
	for _, item := range data {
		imageList = append(imageList, &rpc_service.ImageResponse{
			Id:        int32(item.Id),
			ImageUrl:  item.ImageUrl,
			ImageName: item.ImageName,
			ImageSort: int32(item.ImageSort),
			ImageType: int32(item.ImageType),
			ImageSize: item.ImageSize,
		})
	}
	return &rpc_service.ImageListResponse{ImageList: imageList}, nil
}
