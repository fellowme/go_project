package image_control

import (
	gin_translator "github.com/fellowme/gin_common_library/translator"
	gin_util "github.com/fellowme/gin_common_library/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_project/app/image/image_param"
	"go_project/app/image/image_service"
	"net/http"
	"strconv"
)

type ImageControl struct {
	service image_service.ImageServiceInterface
}

func GetImageControl() *ImageControl {
	return &ImageControl{
		service: image_service.GetImageService(),
	}
}

func (receiver ImageControl) GetImageList(c *gin.Context) {
	var param image_param.GetImageRequestParam
	if err := c.ShouldBindQuery(&param); err != nil {
		zap.L().Error(" image GetImageList error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, gin_translator.GetErrorMessage(err), nil, c)
		return
	}
	data, err := receiver.service.GetImageListByParamService(param)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, nil, data, c)
}

func (receiver ImageControl) CreateImage(c *gin.Context) {
	form, err := c.MultipartForm()
	file := form.File["image"]
	if err != nil {
		zap.L().Error(" image CreateImage error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	data := receiver.service.CreateImageByImageService(c, file)
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, nil, data, c)
}

func (receiver ImageControl) GetImage(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		zap.L().Error(" image GetImage error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	data, err := receiver.service.GetImageByIdService(id)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, nil, data, c)
}

func (receiver ImageControl) DeleteImage(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		zap.L().Error(" image GetImage error", zap.Any("error", err))
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	err = receiver.service.DeleteImageByIdService(id)
	if err != nil {
		gin_util.ReturnResponse(http.StatusOK, gin_util.FailCode, err.Error(), nil, c)
		return
	}
	gin_util.ReturnResponse(http.StatusOK, gin_util.SuccessCode, nil, nil, c)
}
