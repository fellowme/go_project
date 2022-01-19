package role_rpc

import (
	"context"
	"errors"
	gin_util "github.com/fellowme/gin_common_library/util"
	"go_project/app/role/role_const"
	"go_project/app/role/role_dao"
	"go_project/app/role/role_remote_service/remote_rpc"
	service "go_project/rpc_service"
	"strconv"
	"strings"
)

type RpcService struct {
	dao role_dao.RoleDaoInterface
}

func GetRoleRpcService() RpcService {
	return RpcService{
		dao: role_dao.GetRoleDao(),
	}
}

func (s RpcService) GetRoleMenuByUserIds(ctx context.Context, req *service.UserRoleMenuRequest) (*service.UserRoleMenuListResponse, error) {
	idList := make([]int, 0)
	userIdListString := strings.Split(req.UserIds, ",")
	for _, idString := range userIdListString {
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			continue
		}
		idList = append(idList, idInt)
	}
	roleList, err := s.dao.QueryRoleByUserIdsDao(ctx, idList)
	if err != nil {
		return nil, err
	}
	roleIdList := make([]int, 0)
	for _, role := range roleList {
		roleIdList = append(roleIdList, role.Id)
	}
	roleIdSet := gin_util.RemoveRepetitionIntSlice(roleIdList)
	if len(roleIdSet) == 0 {
		return nil, errors.New(role_const.RoleIdNotFindTip)
	}
	roleMenuInfoList, queryError := s.dao.QueryRoleMenuListByRoleIdsDao(ctx, roleIdSet)
	if queryError != nil {
		return nil, queryError
	}
	if len(roleMenuInfoList) == 0 {
		return nil, errors.New(role_const.MenuIdNotFindTip)
	}
	menuIdList := make([]int, 0)
	userRoleMenuList := make([]*service.UserRoleMenuResponse, 0)
	for _, menuInfo := range roleMenuInfoList {
		menuIdList = append(menuIdList, menuInfo.MenuId)
	}
	menuIdSet := gin_util.RemoveRepetitionIntSlice(menuIdList)
	menuList, rpcError := remote_rpc.GetMenuList(ctx, menuIdSet)
	if rpcError != nil {
		return nil, rpcError
	}
	for _, menuInfo := range menuList {
		userRoleMenuList = append(userRoleMenuList, &service.UserRoleMenuResponse{
			Id:       menuInfo.Id,
			MenuName: menuInfo.MenuName,
			MenuPath: menuInfo.MenuPath,
			MenuType: menuInfo.MenuType,
		})
	}
	return &service.UserRoleMenuListResponse{MenuList: userRoleMenuList}, nil
}
