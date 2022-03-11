package remote_rpc

import (
	"context"
	gin_grpc "github.com/fellowme/gin_common_library/grpc"
	gin_util "github.com/fellowme/gin_common_library/util"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go_project/app/brand/brand_const"
	service "go_project/rpc_service"
	"strconv"
	"strings"
	"time"
)

func getImageListByImageIds(ctx context.Context, imageIdList []int) (*service.ImageListResponse, error) {
	imageIdSet := gin_util.RemoveRepetitionIntSlice(imageIdList)
	if len(imageIdSet) == 0 {
		return nil, errors.New(brand_const.ImageIdEmptyTip)
	}
	clientDeadline := time.Now().Add(5 * time.Second)
	contextDeadline, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()
	conn := gin_grpc.GetGRPCConnect(contextDeadline, "127.0.0.1:18088")
	defer gin_grpc.CloseGRPCConnect(conn)
	serviceClient := service.NewImageServiceClient(conn)
	imageIdStringList := make([]string, 0)
	for _, imageId := range imageIdSet {
		imageIdStringList = append(imageIdStringList, strconv.Itoa(imageId))
	}
	resp, err := serviceClient.GetImageByIds(contextDeadline, &service.ImageRequest{
		ImageIds: strings.Join(imageIdStringList, ","),
	})
	if err != nil {
		zap.L().Error("grpc getImageListByImageIds GetImageByIds error", zap.Any("error", err), zap.Any("imageIdList", imageIdList))
		return nil, err
	}
	return resp, nil
}

func GetImageListByImageIdsMap(ctx context.Context, imageIdList []int) (map[int]*service.ImageResponse, error) {
	imageList, err := getImageListByImageIds(ctx, imageIdList)
	if err != nil {
		return nil, err
	}
	dataMap := make(map[int]*service.ImageResponse, 0)
	for _, item := range imageList.ImageList {
		dataMap[int(item.Id)] = item
	}
	return dataMap, nil
}
