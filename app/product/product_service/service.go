package product_service

import (
	"context"
	"errors"
	gin_util "github.com/fellowme/gin_common_library/util"
	"go.uber.org/zap"
	"go_project/app/product/product_dao"
	"go_project/app/product/product_param"
	"go_project/app/product/product_remote_service/remote_rpc"
	"strconv"
	"strings"
)

type ProductServiceInterface interface {
	GetProductMainListServiceByParam(ctx context.Context, req product_param.GetProductMainRequestParam) (product_param.ProductMainListResponse, error)
	PostProductMainServiceByParam(param product_param.PostProductMainRequestParam) error
}

type ProductService struct {
	dao product_dao.ProductDaoInterface
}

func GetProductService() *ProductService {
	return &ProductService{
		dao: product_dao.GetProductDao(),
	}
}

func (s ProductService) GetProductMainListServiceByParam(ctx context.Context, req product_param.GetProductMainRequestParam) (product_param.ProductMainListResponse, error) {
	total, data, err := s.dao.GetProductMainListDaoByParam(req)
	if err != nil {
		return product_param.ProductMainListResponse{}, err
	}
	if total == 0 {
		return product_param.ProductMainListResponse{}, errors.New(gin_util.NotFindTip)
	}
	brandIdList := make([]int, 0)
	categoryIdList := make([]int, 0)
	shopIdList := make([]int, 0)
	productMainIdList := make([]int, 0)
	for _, item := range data {
		brandIdList = append(brandIdList, item.BrandId)
		categoryIdList = append(categoryIdList, item.CategoryId)
		shopIdList = append(shopIdList, item.ShopId)
		productMainIdList = append(productMainIdList, item.Id)
	}
	images, err := s.dao.QueryProductImageByProductMainIds(productMainIdList)
	if err != nil {
		return product_param.ProductMainListResponse{}, err
	}
	imageMapList := make(map[int][]int, 0)
	imageIdAllList := make([]int, 0)
	for _, image := range images {
		imageIdList := make([]int, 0)
		imageIdStringList := strings.Split(image.ImageIds, ",")
		for _, imageIdString := range imageIdStringList {
			imageId, err := strconv.Atoi(imageIdString)
			if err != nil {
				zap.L().Error("GetProductMainListServiceByParam strconv.Atoi error", zap.Any("error", err))
				continue
			}
			imageIdList = append(imageIdList, imageId)
			imageIdAllList = append(imageIdAllList, imageId)
		}
		imageMapList[image.ProductId] = imageIdList
	}
	categoryChanMap := make(chan map[int]string, 1)
	imageChanMap := make(chan map[int]product_param.ImageParam, 1)
	shopChanMap := make(chan map[int]product_param.ShopParam, 1)
	brandChanMap := make(chan map[int]product_param.BrandParam, 1)
	go remote_rpc.GetCategoryListByCategoryIdsChannel(ctx, categoryIdList, categoryChanMap)
	go remote_rpc.GetImageListByImageIdsChannel(ctx, imageIdAllList, imageChanMap)
	go remote_rpc.GetShopListByShopIdsChannel(ctx, shopIdList, shopChanMap)
	go remote_rpc.GetBrandListByBrandIdsChannel(ctx, brandIdList, brandChanMap)
	categoryMap := <-categoryChanMap
	imageMap := <-imageChanMap
	shopMap := <-shopChanMap
	brandMap := <-brandChanMap
	list := make([]product_param.ProductMainResponse, 0)
	for _, item := range data {
		item.CategoryName = categoryMap[item.CategoryId]
		imageIdList := imageMapList[item.Id]
		item.Images = imageIdList
		imageMapList := make([]product_param.ImageParam, 0)
		for _, imageId := range imageIdList {
			data, ok := imageMap[imageId]
			if ok {
				imageMapList = append(imageMapList, data)
			}
		}
		item.ImageMapList = imageMapList
		shop, ok := shopMap[item.ShopId]
		if ok {
			item.ShopName = shop.ShopName
		}
		brand, ok := brandMap[item.BrandId]
		if ok {
			item.BrandName = brand.BrandName
		}
		list = append(list, item)
	}
	return product_param.ProductMainListResponse{
		Total: total,
		List:  list,
	}, nil
}

func (s ProductService) PostProductMainServiceByParam(param product_param.PostProductMainRequestParam) error {
	imageIdList := make([]int, 0)
	imageIdStringList := strings.Split(param.Images, ",")
	for _, imageIdString := range imageIdStringList {
		imageId, err := strconv.Atoi(imageIdString)
		if err != nil {
			zap.L().Error("GetProductMainListServiceByParam strconv.Atoi error", zap.Any("error", err))
			continue
		}
		imageIdList = append(imageIdList, imageId)
	}
	param.ImageIdList = imageIdList
	return s.dao.PostProductMainDaoByParam(param)
}
