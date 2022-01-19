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

func GetUserAccountByAccountIdList(ctx context.Context, accountIdList []int) (map[int]*service.UserResponse, error) {
	userMap := make(map[int]*service.UserResponse, 0)
	accountIdSet := gin_util.RemoveRepetitionIntSlice(accountIdList)
	if len(accountIdSet) == 0 {
		return nil, errors.New(account_const.AccountIdListNotEmptyTip)
	}
	clientDeadline := time.Now().Add(5 * time.Second)
	contextDeadline, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()
	conn := gin_grpc.GetGRPCConnect(contextDeadline, "127.0.0.1:18082")
	defer gin_grpc.CloseGRPCConnect(conn)
	serviceClient := service.NewUserServiceClient(conn)
	accountIdStringList := make([]string, 0)
	for _, accountId := range accountIdList {
		accountIdStringList = append(accountIdStringList, strconv.Itoa(accountId))
	}
	resp, err := serviceClient.GetUserByAccountIds(contextDeadline, &service.UserAccountRequest{
		AccountIds: strings.Join(accountIdStringList, ","),
	})
	if err != nil {
		zap.L().Error("grpc UserList error", zap.Any("error", err), zap.Any("accountIdList", accountIdList))
		return nil, err
	}
	for _, userInfo := range resp.UserList {
		userMap[int(userInfo.AccountId)] = userInfo
	}

	return userMap, nil
}

func GetUserAccountById(ctx context.Context, accountId int) (*service.UserResponse, error) {
	accountList := make([]int, 0)
	accountList = append(accountList, accountId)
	accountMap, err := GetUserAccountByAccountIdList(ctx, accountList)
	if err != nil {
		return nil, err
	}
	userInfo, ok := accountMap[accountId]
	if !ok {
		return nil, errors.New(account_const.AccountNotRegisterErrorTip)
	}
	return userInfo, nil
}
