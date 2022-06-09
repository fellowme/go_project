package role_srevice

import (
	"context"
	"encoding/json"
	"errors"
	"go_project/app/role/role_cache"
	"go_project/app/role/role_const"
	"regexp"

	gin_util "github.com/fellowme/gin_common_library/util"
	"go_project/app/role/role_dao"
	"go_project/app/role/role_param"
	"go_project/app/role/role_remote_service/remote_rpc"
	"strconv"
	"strings"
)

type RoleServiceInterface interface {
	GetRoleListService(param role_param.GetRoleListRequestParam) (role_param.RoleListResponse, error)
	PostRoleService(param role_param.PostRoleRequestParam) error
	PatchRoleService(param role_param.PatchRoleRequestParam) error
	DeleteRoleService(id int) error
	PostRoleUserService(ctx context.Context, param role_param.PostRoleUserRequestParam) error
	DeleteRoleUserService(id int) error
	GetRoleUserService(ctx context.Context, param role_param.GetRoleUserRequestParam) (role_param.RoleUserListResponse, error)
	GetRoleMenuService(ctx context.Context, param role_param.GetRoleMenuRequestParam) (role_param.RoleMenuListResponse, error)
	PostRoleMenuService(ctx context.Context, param role_param.PostRoleMenuRequestParam) error
	DeleteRoleMenuService(id int) error
	RebuildRoleMenuService(ctx context.Context, param role_param.RebuildRoleMenuRequestParam) error
	RoleMenuMapMatchService(param role_param.PostRoleMenuMatchRequestParam) error
}

type RoleService struct {
	dao role_dao.RoleDaoInterface
}

func GetRoleService() RoleService {
	return RoleService{
		dao: role_dao.GetRoleDao(),
	}
}

func (s RoleService) GetRoleListService(param role_param.GetRoleListRequestParam) (role_param.RoleListResponse, error) {
	list, total, err := s.dao.GetRoleListByParamDao(param)
	return role_param.RoleListResponse{
		Total: total,
		List:  list,
	}, err
}

func (s RoleService) PostRoleService(param role_param.PostRoleRequestParam) error {
	return s.dao.PostRoleByParamDao(param)
}

func (s RoleService) PatchRoleService(param role_param.PatchRoleRequestParam) error {
	return s.dao.PatchRoleByParamDao(param)
}

func (s RoleService) DeleteRoleService(id int) error {
	return s.dao.DeleteRoleByIdDao(id)

}

func (s RoleService) PostRoleUserService(ctx context.Context, param role_param.PostRoleUserRequestParam) error {
	RpcError := remote_rpc.GetUserExist(ctx, param.UserId)
	if RpcError != nil {
		return RpcError
	}
	_, MysqlError := s.dao.QueryRoleByIdDao(param.RoleId)
	if MysqlError != nil {
		return MysqlError
	}
	return s.dao.PostRoleUserByParamDao(param)
}

func (s RoleService) DeleteRoleUserService(id int) error {
	return s.dao.DeleteRoleUserByIdDao(id)
}

func (s RoleService) GetRoleUserService(ctx context.Context, param role_param.GetRoleUserRequestParam) (role_param.RoleUserListResponse, error) {
	var roleUserInfo role_param.RoleUserListResponse
	roleUserList, total, err := s.dao.GetRoleUserListByParamDao(param)
	if err != nil {
		return roleUserInfo, err
	}
	if total == 0 {
		return roleUserInfo, errors.New(gin_util.NotFindTip)
	}
	userIdList := make([]int, 0)
	roleIdList := make([]int, 0)
	for _, roleUser := range roleUserList {
		roleIdList = append(roleIdList, roleUser.RoleId)
		userIdList = append(userIdList, roleUser.UserId)
	}
	roleList, queryError := s.dao.QueryRoleByIdListDao(ctx, roleIdList)
	if queryError != nil {
		return roleUserInfo, queryError
	}
	roleMap := make(map[int]string)
	for _, role := range roleList {
		roleMap[role.Id] = role.RoleName
	}
	userMap, rpcError := remote_rpc.GetUserList(ctx, userIdList)
	if rpcError != nil {
		return roleUserInfo, rpcError
	}
	for index, userRole := range roleUserList {
		roleUserList[index].RoleName = roleMap[userRole.RoleId]
		menuInfo, ok := userMap[userRole.UserId]
		if !ok {
			roleUserList[index].UserName = ""
		}
		roleUserList[index].UserName = menuInfo.UserName
	}
	return role_param.RoleUserListResponse{
		Total: total,
		List:  roleUserList,
	}, nil
}

func (s RoleService) GetRoleMenuService(ctx context.Context, param role_param.GetRoleMenuRequestParam) (role_param.RoleMenuListResponse, error) {
	var roleMenuResponse role_param.RoleMenuListResponse
	roleMenuList, total, err := s.dao.GetRoleMenuListByParamDao(param)
	if err != nil {
		return roleMenuResponse, err
	}
	if total == 0 {
		return roleMenuResponse, errors.New(gin_util.NotFindTip)
	}
	menuIdList := make([]int, 0)
	roleIdList := make([]int, 0)
	for _, roleUser := range roleMenuList {
		roleIdList = append(roleIdList, roleUser.RoleId)
		menuIdList = append(menuIdList, roleUser.MenuId)
	}
	roleList, queryError := s.dao.QueryRoleByIdListDao(ctx, roleIdList)
	if queryError != nil {
		return roleMenuResponse, queryError
	}
	roleMap := make(map[int]string)
	for _, role := range roleList {
		roleMap[role.Id] = role.RoleName
	}
	userMap, rpcError := remote_rpc.GetMenuList(ctx, menuIdList)
	if rpcError != nil {
		return roleMenuResponse, rpcError
	}
	for index, userRole := range roleMenuList {
		roleMenuList[index].RoleName = roleMap[userRole.RoleId]
		menuInfo, ok := userMap[userRole.MenuId]
		if !ok {
			roleMenuList[index].Path = ""
			roleMenuList[index].MenuName = ""
			roleMenuList[index].Handler = ""
			roleMenuList[index].Method = ""
			continue
		}
		roleMenuList[index].Path = menuInfo.Path
		roleMenuList[index].MenuName = menuInfo.MenuName
		roleMenuList[index].Handler = menuInfo.Handler
		roleMenuList[index].Method = menuInfo.Method
	}
	roleMenuResponse.Total = total
	roleMenuResponse.List = roleMenuList
	return roleMenuResponse, nil
}

func (s RoleService) PostRoleMenuService(ctx context.Context, param role_param.PostRoleMenuRequestParam) error {
	_, err := s.dao.QueryRoleByIdDao(param.RoleId)
	if err != nil {
		return err
	}
	rpcError := remote_rpc.GetMenuExist(ctx, param.MenuId)
	if rpcError != nil {
		return rpcError
	}
	return s.dao.PostRoleMenuByParamDao(param)
}

func (s RoleService) DeleteRoleMenuService(id int) error {
	return s.dao.DeleteRoleMenuByIdDao(id)

}

func (s RoleService) RebuildRoleMenuService(ctx context.Context, param role_param.RebuildRoleMenuRequestParam) error {
	roleMenuIds, err := s.dao.QueryRoleGroup(ctx, param)
	if err != nil {
		return err
	}
	roleMenuIdMap := make(map[int][]int, 0)
	totalMenuIdList := make([]int, 0)
	for _, roleMenuId := range roleMenuIds {
		menuIdList := make([]int, 0)
		menuIdStringList := strings.Split(roleMenuId.MenuIds, ",")
		for _, menuIdString := range menuIdStringList {
			menuId, _ := strconv.Atoi(menuIdString)
			menuIdList = append(menuIdList, menuId)
		}
		totalMenuIdList = append(totalMenuIdList, menuIdList...)
		roleMenuIdMap[roleMenuId.RoleId] = menuIdList
	}
	res, err := remote_rpc.GetMenuList(ctx, totalMenuIdList)
	if err != nil {
		return err
	}
	roleMenuMap := make(map[int][]role_param.MenuResponseParam, 0)
	for key, value := range roleMenuIdMap {
		menuList := make([]role_param.MenuResponseParam, 0)
		for _, item := range value {
			menuInfo := res[item]
			path := menuInfo.Path
			newPath := ""
			isRegex := false
			if strings.Contains(path, ":") {
				pathList := strings.Split(path, "/")
				for _, item := range pathList {
					if strings.HasPrefix(item, ":") {
						item = "*"
						isRegex = true
					}
					newPath += item + "/"
				}
			} else {
				newPath = path
			}

			menuList = append(menuList, role_param.MenuResponseParam{
				Id:      menuInfo.Id,
				Path:    newPath,
				Method:  menuInfo.Method,
				IsRegex: isRegex,
			})
		}
		roleMenuMap[key] = menuList
	}
	_ = role_cache.SetRoleMenuMapRedisKeyCache(roleMenuMap)
	return nil
}

func (s RoleService) RoleMenuMapMatchService(param role_param.PostRoleMenuMatchRequestParam) error {
	data, err := role_cache.GetRoleMenuMapRedisKeyCache(strconv.Itoa(param.RoleId))
	if err != nil {
		return err
	}
	var menuInfoList []role_param.MenuResponseParam
	err = json.Unmarshal(data, &menuInfoList)
	if err != nil {
		return err
	}
	flag := false
	for _, menuInfo := range menuInfoList {
		if menuInfo.Method == param.Method {
			if menuInfo.IsRegex {
				reg := regexp.MustCompile(menuInfo.Path)
				flag = reg.MatchString(param.Path)
				if flag {
					break
				}
			} else {
				if menuInfo.Path == param.Path {
					flag = true
					break
				}
			}
		}
	}
	if !flag {
		return errors.New(role_const.RoleMenuMapNotMatchTip)
	}
	return nil
}
