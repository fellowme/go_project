package category_control

import (
	gin_translator "github.com/fellowme/gin_common_library/translator"
	gin_util "github.com/fellowme/gin_common_library/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_project/app/category/category_param"
	"go_project/app/category/category_service"
	"net/http"
	"strconv"
)

type CategoryControl struct {
	service category_service.CategoryServiceInterface
}

func GetCategoryControl() *CategoryControl {
	return &CategoryControl{
		service: category_service.GetCategoryService(),
	}
}

func (receiver CategoryControl) GetCategoryList(c *gin.Context) {
	var param category_param.GetCategoryListRequestParam
	if err := c.ShouldBindQuery(&param); err != nil {
		zap.L().Error(" category GetCategoryList error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	data, err := receiver.service.GetCategoryListByParamService(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, nil, data, c)
}

func (receiver CategoryControl) CreateCategory(c *gin.Context) {
	var param category_param.CategoryRequestParam
	if err := c.ShouldBind(&param); err != nil {
		zap.L().Error(" category CreateCategory error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	err := receiver.service.CreateCategoryByParamService(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, nil, nil, c)
}

func (receiver CategoryControl) UpdateCategory(c *gin.Context) {
	var param category_param.CategoryRequestParam
	if err := c.ShouldBind(&param); err != nil {
		zap.L().Error(" category UpdateCategory error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	param.Id = id
	err = receiver.service.UpdateCategoryByParamService(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, nil, nil, c)
}

func (receiver CategoryControl) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	err = receiver.service.DeleteCategoryByIdService(id)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, nil, nil, c)
}

func (receiver CategoryControl) RebuildCategory(c *gin.Context) {
	err := receiver.service.RebuildCategoryService()
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, nil, nil, c)
}
