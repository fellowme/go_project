package role_control

import (
	"context"
	gin_translator "github.com/fellowme/gin_common_library/translator"
	gin_util "github.com/fellowme/gin_common_library/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_project/app/role/role_param"
	"go_project/app/role/role_srevice"
	"net/http"
	"strconv"
)

type RoleControl struct {
	service role_srevice.RoleServiceInterface
}

func GetRoleControl() RoleControl {
	return RoleControl{
		service: role_srevice.GetRoleService(),
	}
}

func (r RoleControl) GetRoleList(c *gin.Context) {
	var param role_param.GetRoleListRequestParam
	if err := c.ShouldBindQuery(&param); err != nil {
		zap.L().Error("role GetRoleList error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	data, err := r.service.GetRoleListService(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SearchSuccessTip, data, c)
}

func (r RoleControl) PostRole(c *gin.Context) {
	var param role_param.PostRoleRequestParam
	if err := c.ShouldBind(&param); err != nil {
		zap.L().Error("role PostRole error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	err := r.service.PostRoleService(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SuccessTip, nil, c)
}

func (r RoleControl) PatchRole(c *gin.Context) {
	var param role_param.PatchRoleRequestParam
	if err := c.ShouldBind(&param); err != nil {
		zap.L().Error("role PatchRole error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	id, strconvError := strconv.Atoi(c.Param("id"))
	if strconvError != nil {
		zap.L().Error("role PatchRole strconv error", zap.Any("error", strconvError))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, strconvError.Error(), nil, c)
		return
	}
	param.Id = id
	err := r.service.PatchRoleService(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SuccessTip, nil, c)
}

func (r RoleControl) DeleteRole(c *gin.Context) {
	id, strconvError := strconv.Atoi(c.Param("id"))
	if strconvError != nil {
		zap.L().Error("role DeleteRole strconv error", zap.Any("error", strconvError))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, strconvError.Error(), nil, c)
		return
	}
	err := r.service.DeleteRoleService(id)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (r RoleControl) PostRoleUser(c *gin.Context) {
	var param role_param.PostRoleUserRequestParam
	if err := c.ShouldBind(&param); err != nil {
		zap.L().Error("role PostRoleUser error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("role PostRoleUser not get tracerContext")
	}
	err := r.service.PostRoleUserService(ctx.(context.Context), param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (r RoleControl) DeleteRoleUser(c *gin.Context) {
	id, strconvError := strconv.Atoi(c.Param("id"))
	if strconvError != nil {
		zap.L().Error("role DeleteRoleUser strconv error", zap.Any("error", strconvError))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, strconvError.Error(), nil, c)
		return
	}
	err := r.service.DeleteRoleUserService(id)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (r RoleControl) GetRoleUser(c *gin.Context) {
	var param role_param.GetRoleUserRequestParam
	if err := c.ShouldBindQuery(&param); err != nil {
		zap.L().Error("role GetRoleUser error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("role PostRoleUser not get tracerContext")
	}
	data, err := r.service.GetRoleUserService(ctx.(context.Context), param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, data, c)
}

func (r RoleControl) PostRoleMenu(c *gin.Context) {
	var param role_param.PostRoleMenuRequestParam
	if err := c.ShouldBind(&param); err != nil {
		zap.L().Error("role PostRoleMenu error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("role PostRoleMenu not get tracerContext")
	}
	err := r.service.PostRoleMenuService(ctx.(context.Context), param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (r RoleControl) DeleteRoleMenu(c *gin.Context) {
	id, strconvError := strconv.Atoi(c.Param("id"))
	if strconvError != nil {
		zap.L().Error("role DeleteRoleMenu strconv error", zap.Any("error", strconvError))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, strconvError.Error(), nil, c)
		return
	}
	err := r.service.DeleteRoleMenuService(id)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (r RoleControl) GetRoleMenu(c *gin.Context) {
	var param role_param.GetRoleMenuRequestParam
	if err := c.ShouldBindQuery(&param); err != nil {
		zap.L().Error("role GetRoleMenu error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("role GetRoleMenu not get tracerContext")
	}
	data, err := r.service.GetRoleMenuService(ctx.(context.Context), param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, data, c)
}

func (r RoleControl) RebuildRoleMenu(c *gin.Context) {
	var param role_param.RebuildRoleMenuRequestParam
	if err := c.ShouldBindQuery(&param); err != nil {
		zap.L().Error("role RebuildRoleMenu error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("role RebuildRoleMenu not get tracerContext")
	}
	err := r.service.RebuildRoleMenuService(ctx.(context.Context), param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}

func (r RoleControl) RoleMenuMapMatch(c *gin.Context) {
	var param role_param.PostRoleMenuMatchRequestParam
	if err := c.ShouldBind(&param); err != nil {
		zap.L().Error("role RoleMenuMapMatch error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}

	err := r.service.RoleMenuMapMatchService(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.ActionSuccessTip, nil, c)
}
