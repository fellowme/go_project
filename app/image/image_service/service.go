package image_service

import (
	gin_util "github.com/fellowme/gin_common_library/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_project/app/image/image_dao"
	"go_project/app/image/image_model"
	"go_project/app/image/image_param"
	"mime/multipart"
	"path/filepath"
)

type ImageServiceInterface interface {
	GetImageListByParamService(param image_param.GetImageRequestParam) (image_param.ImageListResponse, error)
	CreateImageByImageService(c *gin.Context, files []*multipart.FileHeader) map[string]string
	GetImageByIdService(id int) (image_param.ImageResponse, error)
	DeleteImageByIdService(id int) error
}

type ImageService struct {
	dao image_dao.ImageDaoInterface
}

func GetImageService() *ImageService {
	return &ImageService{
		dao: image_dao.GetImageDao(),
	}
}

func (s ImageService) GetImageListByParamService(param image_param.GetImageRequestParam) (image_param.ImageListResponse, error) {
	total, data, err := s.dao.GetImageListByParamDao(param)
	return image_param.ImageListResponse{
		Total: total,
		Data:  data,
	}, err
}

func (s ImageService) CreateImageByImageService(c *gin.Context, files []*multipart.FileHeader) map[string]string {
	basePath := gin_util.GetPath()
	imageMap := make(map[string]string, 0)
	userId := c.GetInt("user_id")
	for index, file := range files {
		filename := "/upload/" + filepath.Base(file.Filename)
		osName := basePath + filename
		if err := c.SaveUploadedFile(file, osName); err != nil {
			zap.L().Error("CreateImageByImageService error", zap.Any("error", err), zap.String("fileName", filename))
			imageMap[filename] = err.Error()
		} else {
			err := s.dao.CreateImageDao(image_model.Image{
				ImageUrl:        "/upload/" + filepath.Base(file.Filename),
				ImageName:       file.Filename,
				ImageUniqueName: "",
				ImageSort:       index,
				ImageType:       0,
				ImageHeight:     "",
				ImageWidth:      "",
				ImageSize:       file.Size,
				CreateUserId:    userId,
			})
			if err != nil {
				imageMap[filename] = err.Error()
			} else {
				imageMap[filename] = ""
			}
		}
	}
	return imageMap
}

func (s ImageService) GetImageByIdService(id int) (image_param.ImageResponse, error) {
	return s.dao.GetImageByIdDao(id)

}

func (s ImageService) DeleteImageByIdService(id int) error {
	return s.dao.DeleteImageByIdDao(id)

}
