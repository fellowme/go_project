package menu_rpc

import (
	"context"
	"go_project/app/menu/menu_dao"
	service "go_project/rpc_service"
	"strconv"
	"strings"
)

type RpcService struct {
	dao menu_dao.MenuDaoInterface
}

func GetMenuRpcService() RpcService {
	return RpcService{
		dao: menu_dao.GetMenuDao(),
	}
}

func (u RpcService) GetMenuByIds(c context.Context, req *service.MenuRequest) (*service.MenuListResponse, error) {
	idList := make([]int, 0)
	menuIdListString := strings.Split(req.MenuIds, ",")
	for _, idString := range menuIdListString {
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			continue
		}
		idList = append(idList, idInt)
	}
	menuList := make([]*service.MenuResponse, 0)
	menuInfoList, err := u.dao.GetMenuListByIdsDao(idList)
	for _, menuInfo := range menuInfoList {
		menuList = append(menuList, &service.MenuResponse{
			Id:       int32(menuInfo.Id),
			MenuName: menuInfo.MenuName,
			MenuPath: menuInfo.MenuPath,
			Remark:   menuInfo.Remark,
			MenuType: int32(menuInfo.MenuType),
		})
	}
	return &service.MenuListResponse{
		MenuList: menuList,
	}, err
}
