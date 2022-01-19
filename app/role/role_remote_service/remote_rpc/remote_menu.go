package remote_rpc

import (
	"context"
	"errors"
	gin_grpc "github.com/fellowme/gin_common_library/grpc"
	gin_util "github.com/fellowme/gin_common_library/util"
	"go.uber.org/zap"
	"go_project/app/menu/menu_const"
	"go_project/app/role/role_const"
	service "go_project/rpc_service"
	"strconv"
	"strings"
	"time"
)

func GetMenuList(ctx context.Context, menuIdList []int) (map[int]*service.MenuResponse, error) {
	menuResponse := make(map[int]*service.MenuResponse, 0)
	menuIdSet := gin_util.RemoveRepetitionIntSlice(menuIdList)
	if len(menuIdSet) == 0 {
		return menuResponse, errors.New(menu_const.MenuIdListNotEmptyTip)
	}
	menuIdStringList := make([]string, 0)
	for _, menuId := range menuIdSet {
		menuIdStringList = append(menuIdStringList, strconv.Itoa(menuId))
	}
	clientDeadline := time.Now().Add(10 * time.Second)
	contextDeadline, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()
	conn := gin_grpc.GetGRPCConnect(contextDeadline, "127.0.0.1:18085")
	defer gin_grpc.CloseGRPCConnect(conn)
	serviceClient := service.NewMenuServiceClient(conn)
	resp, err := serviceClient.GetMenuByIds(contextDeadline, &service.MenuRequest{
		MenuIds: strings.Join(menuIdStringList, ","),
	})
	if err != nil {
		zap.L().Error("grpc GetMenuList GetMenuByIds error", zap.Any("error", err), zap.Any("menuIdList", menuIdList))
		return menuResponse, err
	}
	for _, menuInfo := range resp.MenuList {
		menuResponse[int(menuInfo.Id)] = menuInfo
	}
	return menuResponse, nil
}

func GetMenuExist(ctx context.Context, menuId int) error {
	menuIdList := make([]int, 0)
	menuIdList = append(menuIdList, menuId)
	menuMap, err := GetMenuList(ctx, menuIdList)
	if err != nil {
		return err
	}
	_, ok := menuMap[menuId]
	if !ok {
		return errors.New(role_const.MenuIdNotFindTip)
	}
	return nil
}
