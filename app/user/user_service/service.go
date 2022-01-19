package user_service

import (
	"errors"
	"go_project/app/user/user_const"
	"go_project/app/user/user_dao"
	"go_project/app/user/user_model"
	"go_project/app/user/user_param"
	"strconv"
	"strings"

	gin_util "github.com/fellowme/gin_common_library/util"
)

type UserServiceInterface interface {
	GetUserByIdService(id int) (user_param.UserInfoResponse, error)
	GetUserByIdsService(param user_param.UserByIdsRequestParam) ([]user_param.UserInfoResponse, error)
	GetUserListService(param user_param.UserListRequestParam) user_param.UserListResponse
	PostCreateUserService(param user_param.UserRequestParam) (user_param.UserInfoResponse, error)
	PatchUpdateUserService(param user_param.UserPatchRequestParam) (user_param.UserInfoResponse, error)
	DeleteUserByIdService(id int) error
	DeleteUserByParamService(param user_param.DeleteUserListRequestParam) error
}

type UserService struct {
	dao user_dao.UserDaoInterface
}

func GetUserService() *UserService {
	return &UserService{
		dao: user_dao.GetUserDao(),
	}
}

func (service UserService) GetUserByIdsService(param user_param.UserByIdsRequestParam) ([]user_param.UserInfoResponse, error) {
	idStringList := strings.Split(param.UserIds, ",")
	idIntList := make([]int, 0)
	for _, idString := range idStringList {
		if idString != "" {
			id, _ := strconv.Atoi(idString)
			idIntList = append(idIntList, id)
		}

	}
	data, err := service.dao.GetUserByIdsDao(idIntList)
	list := make([]user_param.UserInfoResponse, 0)
	for _, userInfo := range data {
		list = append(list, user_param.UserInfoResponse{
			Id:         userInfo.Id,
			UserName:   userInfo.UserName,
			NickName:   userInfo.NickName,
			RealName:   userInfo.RealName,
			Gender:     user_model.GetGenderMapCode(userInfo.Gender),
			UserStatus: user_model.GetUserStatusMapCode(userInfo.UserStatus),
		})
	}
	return list, err

}

func (service UserService) GetUserListService(param user_param.UserListRequestParam) user_param.UserListResponse {
	if param.Page == 0 {
		param.Page = gin_util.DefaultPage
	}
	if param.PageSize == 0 {
		param.PageSize = gin_util.DefaultPageSize
	}
	data, total := service.dao.GetUserListDao(param)
	userInfoList := make([]user_param.UserInfoResponse, 0)
	for _, userInfo := range data {
		userInfoList = append(userInfoList, user_param.UserInfoResponse{
			Id:         userInfo.Id,
			UserName:   userInfo.UserName,
			NickName:   userInfo.NickName,
			RealName:   userInfo.RealName,
			Gender:     user_model.GetGenderMapCode(userInfo.Gender),
			UserStatus: user_model.GetUserStatusMapCode(userInfo.UserStatus),
		})
	}
	return user_param.UserListResponse{
		Total: total,
		List:  userInfoList,
	}
}

func (service UserService) PostCreateUserService(param user_param.UserRequestParam) (user_param.UserInfoResponse, error) {
	userInfo, err := service.dao.PostCreateUserDao(param)
	if err != nil {
		return user_param.UserInfoResponse{}, err
	}
	return user_param.UserInfoResponse{
		Id:       userInfo.Id,
		UserName: userInfo.UserName,
		Gender:   user_model.GetGenderMapCode(userInfo.Gender),
	}, nil
}

func (service UserService) PatchUpdateUserService(param user_param.UserPatchRequestParam) (user_param.UserInfoResponse, error) {
	userInfo, err := service.dao.PatchUpdateUserDao(param)
	if err != nil {
		return user_param.UserInfoResponse{}, err
	}
	return user_param.UserInfoResponse{
		Id:       userInfo.Id,
		UserName: userInfo.UserName,
		Gender:   user_model.GetGenderMapCode(userInfo.Gender),
	}, nil
}

func (service UserService) GetUserByIdService(id int) (user_param.UserInfoResponse, error) {
	userInfo, err := service.dao.GetUserByIdDao(id)
	if userInfo.Id == 0 {
		return user_param.UserInfoResponse{}, err
	}
	return user_param.UserInfoResponse{
		Id:       userInfo.Id,
		UserName: userInfo.UserName,
		NickName: userInfo.NickName,
		RealName: userInfo.RealName,
		Gender:   user_model.GetGenderMapCode(userInfo.Gender),
	}, nil

}

func (service UserService) DeleteUserByIdService(id int) error {
	idList := make([]int, 0)
	idList = append(idList, id)
	return service.dao.DeleteUserByIdListDao(idList)
}

func (service UserService) DeleteUserByParamService(param user_param.DeleteUserListRequestParam) error {
	idListString := strings.Split(param.Ids, ",")
	if len(idListString) == 0 {
		return errors.New(user_const.ParamErrorTip)
	}
	idList := make([]int, 0)
	for _, idString := range idListString {
		id, _ := strconv.Atoi(idString)
		idList = append(idList, id)
	}
	return service.dao.DeleteUserByIdListDao(idList)
}
