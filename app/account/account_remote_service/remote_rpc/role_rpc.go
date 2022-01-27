package remote_rpc

import (
	"context"
	"errors"
	gin_grpc "github.com/fellowme/gin_common_library/grpc"
	gin_util "github.com/fellowme/gin_common_library/util"
	"go.uber.org/zap"
	"go_project/app/account/account_const"
	service "go_project/rpc_service"
	"strconv"
	"strings"
	"time"
)

func GetUserRoleByUserIdList(ctx context.Context, userIdList []int) (*service.UserRoleListResponse, error) {
	userIdSet := gin_util.RemoveRepetitionIntSlice(userIdList)
	if len(userIdSet) == 0 {
		return nil, errors.New(account_const.UserIdListNotEmptyTip)
	}
	clientDeadline := time.Now().Add(5 * time.Second)
	contextDeadline, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()
	conn := gin_grpc.GetGRPCConnect(contextDeadline, "127.0.0.1:18084")
	defer gin_grpc.CloseGRPCConnect(conn)
	serviceClient := service.NewUserRoleServiceClient(conn)
	userIdStringList := make([]string, 0)
	for _, accountId := range userIdList {
		userIdStringList = append(userIdStringList, strconv.Itoa(accountId))
	}
	resp, err := serviceClient.GetRoleByUserIds(contextDeadline, &service.UserRoleRequest{
		UserIds: strings.Join(userIdStringList, ","),
	})
	if err != nil {
		zap.L().Error("grpc GetUserRoleByUserIdList GetRoleByUserIds error", zap.Any("error", err), zap.Any("userIdList", userIdList))
		return nil, err
	}
	return resp, nil
}

func GetUserRoleMenuByUserId(ctx context.Context, userId int) ([]int32, error) {
	userIdList := make([]int, 0)
	userIdList = append(userIdList, userId)
	userRoleResponse, err := GetUserRoleByUserIdList(ctx, userIdList)
	if err != nil {
		return nil, err
	}
	roleIdList := make([]int32, 0)
	for _, roleId := range userRoleResponse.RoleIdList {
		roleIdList = append(roleIdList, roleId)
	}
	return roleIdList, nil
}
