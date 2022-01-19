package user_rpc

import (
	"context"
	"go_project/app/user/user_dao"
	"go_project/app/user/user_model"
	service "go_project/rpc_service"
	"strconv"
	"strings"
)

type RpcService struct {
	dao user_dao.UserDaoInterface
}

func GetUserRpcService() RpcService {
	return RpcService{
		dao: user_dao.GetUserDao(),
	}
}

func (u RpcService) GetUserByIds(c context.Context, req *service.UserRequest) (*service.UserListResponse, error) {
	idList := make([]int, 0)
	userIdListString := strings.Split(req.UserIds, ",")
	for _, idString := range userIdListString {
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			continue
		}
		idList = append(idList, idInt)
	}
	userList := make([]*service.UserResponse, 0)
	userInfoList, err := u.dao.GetUserByIdsDao(idList)
	for _, userInfo := range userInfoList {
		userList = append(userList, &service.UserResponse{
			Id:         int32(userInfo.Id),
			UserName:   userInfo.UserName,
			NickName:   userInfo.NickName,
			RealName:   userInfo.RealName,
			Gender:     user_model.GetGenderMapCode(userInfo.Gender),
			UserStatus: user_model.GetUserStatusMapCode(userInfo.UserStatus),
		})
	}
	return &service.UserListResponse{
		UserList: userList,
	}, err
}

func (u RpcService) GetUserByAccountIds(c context.Context, req *service.UserAccountRequest) (*service.UserListResponse, error) {
	accountIdList := make([]int, 0)
	accountIdListString := strings.Split(req.AccountIds, ",")
	for _, accountIdString := range accountIdListString {
		accountId, err := strconv.Atoi(accountIdString)
		if err != nil {
			continue
		}
		accountIdList = append(accountIdList, accountId)
	}
	userList := make([]*service.UserResponse, 0)
	userInfoList, err := u.dao.GetUserByAccountIdsDao(accountIdList)
	for _, userInfo := range userInfoList {
		userList = append(userList, &service.UserResponse{
			Id:         int32(userInfo.Id),
			AccountId:  int32(userInfo.AccountId),
			UserName:   userInfo.UserName,
			NickName:   userInfo.NickName,
			RealName:   userInfo.RealName,
			Gender:     user_model.GetGenderMapCode(userInfo.Gender),
			UserStatus: user_model.GetUserStatusMapCode(userInfo.UserStatus),
		})
	}
	return &service.UserListResponse{
		UserList: userList,
	}, err
}
