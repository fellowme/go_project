package remote_rpc

import (
	"context"
	"errors"
	gin_grpc "github.com/fellowme/gin_common_library/grpc"
	gin_util "github.com/fellowme/gin_common_library/util"
	"go.uber.org/zap"
	"go_project/app/account/account_const"
	"go_project/app/account/account_param"
	service "go_project/rpc_service"
	"strconv"
	"strings"
	"time"
)

func GetUserRoleMenuByUserIdList(ctx context.Context, userIdList []int) (*service.UserRoleMenuListResponse, error) {
	userIdSet := gin_util.RemoveRepetitionIntSlice(userIdList)
	if len(userIdSet) == 0 {
		return nil, errors.New(account_const.UserIdListNotEmptyTip)
	}
	clientDeadline := time.Now().Add(5 * time.Second)
	contextDeadline, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()
	conn := gin_grpc.GetGRPCConnect(contextDeadline, "127.0.0.1:18084")
	defer gin_grpc.CloseGRPCConnect(conn)
	serviceClient := service.NewUserRoleMenuServiceClient(conn)
	userIdStringList := make([]string, 0)
	for _, accountId := range userIdList {
		userIdStringList = append(userIdStringList, strconv.Itoa(accountId))
	}
	resp, err := serviceClient.GetRoleMenuByUserIds(contextDeadline, &service.UserRoleMenuRequest{
		UserIds: strings.Join(userIdStringList, ","),
	})
	if err != nil {
		zap.L().Error("grpc GetUserRoleMenuByUserIdList GetRoleMenuByUserIds error", zap.Any("error", err), zap.Any("userIdList", userIdList))
		return nil, err
	}
	return resp, nil
}

func GetUserRoleMenuByUserId(ctx context.Context, userId int) ([]account_param.SessionMenuParam, error) {
	userIdList := make([]int, 0)
	userIdList = append(userIdList, userId)
	userRoleMenuResponse, err := GetUserRoleMenuByUserIdList(ctx, userIdList)
	if err != nil {
		return nil, err
	}
	menuSession := make([]account_param.SessionMenuParam, 0)
	for _, menuInfo := range userRoleMenuResponse.MenuList {
		menuSession = append(menuSession, account_param.SessionMenuParam{
			Id:       int(menuInfo.Id),
			MenuName: menuInfo.MenuName,
			Path:     menuInfo.Path,
			Method:   menuInfo.Method,
			Handler:  menuInfo.Handler,
		})
	}
	return menuSession, nil
}
