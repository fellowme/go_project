package remote_rpc

import (
	"context"
	"errors"
	gin_grpc "github.com/fellowme/gin_common_library/grpc"
	gin_util "github.com/fellowme/gin_common_library/util"
	"go.uber.org/zap"
	"go_project/app/role/role_const"
	service "go_project/rpc_service"
	"strconv"
	"strings"
	"time"
)

func GetUserList(ctx context.Context, userIdList []int) (map[int]*service.UserResponse, error) {
	userBool := make(map[int]*service.UserResponse, 0)
	userIdSet := gin_util.RemoveRepetitionIntSlice(userIdList)
	if len(userIdSet) == 0 {
		return userBool, errors.New(role_const.UserIdListNotEmptyTip)
	}
	clientDeadline := time.Now().Add(10 * time.Second)
	contextDeadline, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()
	conn := gin_grpc.GetGRPCConnect(contextDeadline, "127.0.0.1:18082")
	defer gin_grpc.CloseGRPCConnect(conn)
	serviceClient := service.NewUserServiceClient(conn)
	userIdStringList := make([]string, 0)
	for _, userId := range userIdSet {
		userIdStringList = append(userIdStringList, strconv.Itoa(userId))
	}
	resp, err := serviceClient.GetUserByIds(contextDeadline, &service.UserRequest{
		UserIds: strings.Join(userIdStringList, ","),
	})
	if err != nil {
		zap.L().Error("grpc GetUseListrExist GetUserByIds error", zap.Any("error", err), zap.Any("userIdList", userIdList))
		return userBool, err
	}
	for _, userInfo := range resp.UserList {
		userBool[int(userInfo.Id)] = userInfo
	}
	return userBool, nil
}

func GetUserExist(ctx context.Context, userId int) error {
	userIdList := make([]int, 0)
	userIdList = append(userIdList, userId)
	userBool, err := GetUserList(ctx, userIdList)
	if err != nil {
		return err
	}
	_, ok := userBool[userId]
	if !ok {
		return errors.New(role_const.UserIdNotFindTip)
	}
	return nil
}
