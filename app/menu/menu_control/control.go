package menu_control

import (
	"context"
	gin_translator "github.com/fellowme/gin_common_library/translator"
	gin_util "github.com/fellowme/gin_common_library/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_project/app/menu/menu_param"
	"go_project/app/menu/menu_service"
	"net/http"
	"strconv"
)

type MenuControl struct {
	service menu_service.MenuServiceInterface
}

func GetMenuControl() MenuControl {
	return MenuControl{
		service: menu_service.GetMenuService(),
	}
}

func (receiver MenuControl) PostMenu(c *gin.Context) {
	var param menu_param.PostMenuRequestParam
	if err := c.ShouldBind(&param); err != nil {
		zap.L().Error("menu PostMenu error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	err := receiver.service.PostMenuService(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SuccessTip, nil, c)
}

func (receiver MenuControl) GetMenuList(c *gin.Context) {
	var param menu_param.GetMenuRequestParam
	if err := c.ShouldBindQuery(&param); err != nil {
		zap.L().Error("menu GetMenuList error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	ctx, ok := c.Get("tracerContext")
	if !ok {
		zap.L().Warn("GetMenuList not get tracerContext")
	}
	data, err := receiver.service.GetMenuListService(ctx.(context.Context), param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SearchSuccessTip, data, c)
}

func (receiver MenuControl) PatchMenu(c *gin.Context) {
	var param menu_param.PatchMenuRequestParam
	if err := c.ShouldBind(&param); err != nil {
		zap.L().Error("menu PatchMenu error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	id, strconvError := strconv.Atoi(c.Param("id"))
	if strconvError != nil {
		zap.L().Error("menu PatchMenu strconv error", zap.Any("error", strconvError))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, strconvError.Error(), nil, c)
		return
	}
	param.Id = id
	err := receiver.service.PatchMenuService(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SuccessTip, nil, c)
}

func (receiver MenuControl) DeleteMenu(c *gin.Context) {
	id, strconvError := strconv.Atoi(c.Param("id"))
	if strconvError != nil {
		zap.L().Error("menu PatchMenu strconv error", zap.Any("error", strconvError))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, strconvError.Error(), nil, c)
		return
	}
	err := receiver.service.DeleteMenuService(id)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, gin_util.SuccessTip, nil, c)
}
