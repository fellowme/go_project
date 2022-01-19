package user_control

import (
	"go_project/app/user/user_param"
	"go_project/app/user/user_service"
	"net/http"
	"strconv"

	gin_translator "github.com/fellowme/gin_common_library/translator"
	gin_util "github.com/fellowme/gin_common_library/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserControl struct {
	service user_service.UserServiceInterface
}

func GetUserControl() *UserControl {
	return &UserControl{
		service: user_service.GetUserService(),
	}
}

/*
	getUserByIds  根据id 获取 用户信息
*/

func (u UserControl) GetUserByIds(c *gin.Context) {
	var param user_param.UserByIdsRequestParam
	if err := c.BindQuery(&param); err != nil {
		zap.L().Error("user getUserByIds error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	data, err := u.service.GetUserByIdsService(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SearchSuccessTip, data, c)
}

/*
	getUserList  获取用户列表
*/

func (u UserControl) GetUserList(c *gin.Context) {
	var param user_param.UserListRequestParam
	if err := c.BindQuery(&param); err != nil {
		zap.L().Error("user getUserList error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	data := u.service.GetUserListService(param)
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SearchSuccessTip, data, c)

}

/*
	createUser  创建用户
*/

func (u UserControl) CreateUser(c *gin.Context) {
	var param user_param.UserRequestParam
	if err := c.ShouldBind(&param); err != nil {
		zap.L().Error("user createUser error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	data, err := u.service.PostCreateUserService(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SearchSuccessTip, data, c)
}

/*
	updateUser  更新用户
*/

func (u UserControl) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var param user_param.UserPatchRequestParam
	if err := c.ShouldBind(&param); err != nil {
		zap.L().Error("user createUser error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	param.Id = id
	data, err := u.service.PatchUpdateUserService(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SearchSuccessTip, data, c)
}

/*
	getUser  获取单个用户
*/

func (u UserControl) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error("user getUser error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	data, err := u.service.GetUserByIdService(id)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SearchSuccessTip, data, c)
}

/*
	getUser  删除单个用户
*/

func (u UserControl) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error("user getUser error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	err = u.service.DeleteUserByIdService(id)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SuccessTip, nil, c)
}

/*
	getUser  删除多个用户
*/

func (u UserControl) DeleteUserByIdList(c *gin.Context) {
	var param user_param.DeleteUserListRequestParam
	if err := c.ShouldBind(&param); err != nil {
		zap.L().Error("user deleteUserByIdList error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	err := u.service.DeleteUserByParamService(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SuccessTip, nil, c)
}
