package product_es

import (
	"context"
	"encoding/json"
	"github.com/apache/pulsar-client-go/pulsar"
	gin_es "github.com/fellowme/gin_common_library/elastic"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"go_project/app/product/product_const"
	"go_project/app/product/product_dao"
	"go_project/app/product/product_param"
	"go_project/app/product/product_remote_service/remote_rpc"
	"go_project/app/product/product_util"
	"strconv"
	"strings"
)

type ProductEsService struct {
	dao product_dao.ProductDaoInterface
}

func GetProductEsService() *ProductEsService {
	return &ProductEsService{
		dao: product_dao.GetProductDao(),
	}
}

func (p ProductEsService) SendProductMainToEs(message pulsar.Message) {
	zap.L().Info("SendProductMain message", zap.Any("message", message))
	if len(message.Payload()) != 0 {
		var ids []int
		if err := json.Unmarshal(message.Payload(), &ids); err != nil {
			zap.L().Error("SendProductMain json.Unmarsha error", zap.Any("error", err), zap.Any("message", message))
			return
		}
		productMainInfoList, err := p.dao.QueryProductMainListDaoByIds(ids)
		if err != nil {
			zap.L().Error("SendProductMain QueryProductMainListDaoByIds error", zap.Any("error", err), zap.Any("message", message))
			return
		}
		if len(productMainInfoList) == 0 {
			zap.L().Error("SendProductMain QueryProductMainListDaoByIds empty", zap.Any("message", message))
			return
		}
		productInfoList, err := p.dao.QueryMainProductListDaoByProductMainIds(ids)
		if err != nil {
			zap.L().Error("SendProductMain QueryMainProductListDaoByProductMainIds error", zap.Any("error", err), zap.Any("message", message))
			return
		}
		if len(productInfoList) == 0 {
			zap.L().Error("SendProductMain QueryMainProductListDaoByProductMainIds empty", zap.Any("message", message))
			return
		}
		productMap := make(map[int]product_param.ProductResponse, 0)
		for _, productInfo := range productInfoList {
			productMap[productInfo.ProductMainId] = productInfo
		}
		brandIdList := make([]int, 0)
		categoryIdList := make([]int, 0)
		shopIdList := make([]int, 0)
		for _, item := range productMainInfoList {
			brandIdList = append(brandIdList, item.BrandId)
			categoryIdList = append(categoryIdList, item.CategoryId)
			shopIdList = append(shopIdList, item.ShopId)
		}
		images, _ := p.dao.QueryProductImageByProductIds(ids)
		imageMapList := make(map[int][]int, 0)
		imageIdAllList := make([]int, 0)
		for _, image := range images {
			imageIdList := make([]int, 0)
			imageIdStringList := strings.Split(image.ImageIds, ",")
			for _, imageIdString := range imageIdStringList {
				imageId, err := strconv.Atoi(imageIdString)
				if err != nil {
					zap.L().Error("SendProductMainToEs strconv.Atoi error", zap.Any("error", err))
					continue
				}
				imageIdList = append(imageIdList, imageId)
				imageIdAllList = append(imageIdAllList, imageId)
			}
			imageMapList[image.ProductId] = imageIdList
		}
		ctx := context.Background()
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
		list := make([]product_param.ProductMainExtEsResponse, 0)
		for _, item := range productMainInfoList {
			ProductMainExtEsInfo := product_param.ProductMainExtEsResponse{
				ProductMainResponse: item,
			}
			ProductMainExtEsInfo.SaleTimeString = item.SaleTime.String()
			ProductMainExtEsInfo.CategoryName = categoryMap[item.CategoryId]
			ProductMainExtEsInfo.ProductMainTypeName = product_util.GetProductMainTypeNameByCode(item.ProductMainType)
			ProductMainExtEsInfo.ProductMainStatusString = product_util.GetProductMainStatusNameByCode(item.ProductMainStatus)
			imageIdList := imageMapList[item.Id]
			ProductMainExtEsInfo.Images = imageIdList
			imageMapList := make([]product_param.ImageParam, 0)
			for _, imageId := range imageIdList {
				data, ok := imageMap[imageId]
				if ok {
					imageMapList = append(imageMapList, data)
				}
			}
			ProductMainExtEsInfo.ImageMapList = imageMapList
			shop, ok := shopMap[item.ShopId]
			if ok {
				ProductMainExtEsInfo.ShopName = shop.ShopName
			}
			brand, ok := brandMap[item.BrandId]
			if ok {
				ProductMainExtEsInfo.BrandName = brand.BrandName
			}
			product, ok := productMap[item.Id]
			if ok {
				ProductMainExtEsInfo.Product = product
			}
			list = append(list, ProductMainExtEsInfo)
		}
		es := gin_es.GetElasticClient()
		req := es.Bulk().Index(product_const.ProductMainIndex)
		for _, productMain := range list {
			doc := elastic.NewBulkIndexRequest().Id(strconv.Itoa(productMain.Id)).Doc(productMain)
			req.Add(doc)
		}
		if req.NumberOfActions() < 0 {
			zap.L().Info("SendProductMainToEs NewBulkIndexRequest empty")
			return
		}
		if _, err := req.Do(ctx); err != nil {
			zap.L().Error("SendProductMainToEs req NewBulkIndexRequest error ", zap.Any("error", err))
		}
	}
}
