package role_rpc

import (
	"context"
	"errors"
	gin_util "github.com/fellowme/gin_common_library/util"
	"go_project/app/role/role_const"
	"go_project/app/role/role_dao"
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

func (s RpcService) GetRoleByUserIds(ctx context.Context, req *service.UserRoleRequest) (*service.UserRoleListResponse, error) {
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
	roleIdListInt32 := make([]int32, 0)
	for _, roleId := range roleIdSet {
		roleIdListInt32 = append(roleIdListInt32, int32(roleId))
	}
	return &service.UserRoleListResponse{RoleIdList: roleIdListInt32}, nil
}
