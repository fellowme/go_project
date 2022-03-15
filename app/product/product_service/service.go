package product_service

import (
	"context"
	"errors"
	gin_util "github.com/fellowme/gin_common_library/util"
	"go.uber.org/zap"
	"go_project/app/product/product_const"
	"go_project/app/product/product_dao"
	"go_project/app/product/product_mq"
	"go_project/app/product/product_param"
	"go_project/app/product/product_remote_service/remote_rpc"
	"go_project/app/product/product_util"
	"strconv"
	"strings"
)

type ProductServiceInterface interface {
	GetProductMainListServiceByParam(ctx context.Context, req product_param.GetProductMainRequestParam) (product_param.ProductMainListResponse, error)
	PostProductMainServiceByParam(param product_param.PostProductMainRequestParam) error
	GetProductMainServiceById(ctx context.Context, id int) (product_param.ProductMainExtResponse, error)
	PatchProductMainServiceByParam(param product_param.PostProductMainRequestParam) error
	DeleteProductMainServiceById(id int) error
	PostDeleteProductMainAllServiceByParam(param product_param.PostDeleteProductMainAllRequestParam) error
	PostProductMainToMqServiceByParam(param product_param.PostProductIdsToMqRequestParam) (int64, error)

	PostProductServiceByParam(param product_param.PostProductRequestParam) error
	GetProductServiceByParam(ctx context.Context, param product_param.GetProductRequestParam) (product_param.ProductListResponse, error)
	GetProductServiceById(ctx context.Context, id int) (product_param.ProductExtResponse, error)
	PatchProductServiceByParam(param product_param.PostProductRequestParam) error
	DeleteProductServiceById(id int) error
	GetProductByProductMainIdsServiceByParam(ctx context.Context, param product_param.PostProductIdsRequestParam) ([]product_param.ProductExtResponse, error)
	PostDeleteProductServiceByParam(param product_param.DeletePostProductIdsRequestParam) error

	GetProductStockListServiceByParam(param product_param.GetProductStockRequestParam) (product_param.ProductStockListResponse, error)
	PostProductStockServiceByParam(param product_param.PostProductStockRequestParam) error
	PatchProductStockServiceByParam(param product_param.PostProductStockRequestParam) error
	DeleteProductStockServiceById(id int) error
	DeleteProductStockByIdServiceById(param product_param.PostProductStockByIdsRequestParam) error
	GetProductStockByIdServiceById(param product_param.PostProductStockByIdsRequestParam) ([]product_param.ProductStockResponse, error)
}

type ProductService struct {
	dao product_dao.ProductDaoInterface
}

func GetProductService() *ProductService {
	return &ProductService{
		dao: product_dao.GetProductDao(),
	}
}

/*
	GetProductMainListServiceByParam 获取spu列表
*/
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
	stockMap := make(map[int]int64, 0)
	images, _ := s.dao.QueryProductImageByProductIds(productMainIdList)
	stocks, _ := s.dao.QueryProductStockByProductMainIds(productMainIdList)
	for _, item := range stocks {
		stockMap[item.ProductMainId] = item.StockTotal
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
	list := make([]product_param.ProductMainExtResponse, 0)
	for _, item := range data {
		productMainExtInfo := product_param.ProductMainExtResponse{
			ProductMainResponse: item,
		}
		productMainExtInfo.Stock = stockMap[item.Id]
		productMainExtInfo.SaleTimeString = item.SaleTime.String()
		productMainExtInfo.CategoryName = categoryMap[item.CategoryId]
		productMainExtInfo.ProductMainTypeName = product_util.GetProductMainTypeNameByCode(item.ProductMainType)
		productMainExtInfo.ProductMainStatusString = product_util.GetProductMainStatusNameByCode(item.ProductMainStatus)
		imageIdList := imageMapList[item.Id]
		productMainExtInfo.Images = imageIdList
		imageMapList := make([]product_param.ImageParam, 0)
		for _, imageId := range imageIdList {
			data, ok := imageMap[imageId]
			if ok {
				imageMapList = append(imageMapList, data)
			}
		}
		productMainExtInfo.ImageMapList = imageMapList
		shop, ok := shopMap[item.ShopId]
		if ok {
			productMainExtInfo.ShopName = shop.ShopName
		}
		brand, ok := brandMap[item.BrandId]
		if ok {
			productMainExtInfo.BrandName = brand.BrandName
		}
		list = append(list, productMainExtInfo)
	}
	return product_param.ProductMainListResponse{
		Total: total,
		List:  list,
	}, nil
}

/*
	PostProductMainServiceByParam 新建spu信息
*/
func (s ProductService) PostProductMainServiceByParam(param product_param.PostProductMainRequestParam) error {
	imageIdList := make([]int, 0)
	if param.Images != "" {
		imageIdStringList := strings.Split(param.Images, ",")
		for _, imageIdString := range imageIdStringList {
			imageId, err := strconv.Atoi(imageIdString)
			if err != nil {
				zap.L().Error("PostProductMainServiceByParam strconv.Atoi error", zap.Any("error", err))
				continue
			}
			imageIdList = append(imageIdList, imageId)
		}
	}
	return s.dao.PostProductMainDaoByParam(product_param.PostProductMainExtRequestParam{
		PostProductMainRequestParam: param,
		ImageIdList:                 imageIdList,
	})
}

/*
	GetProductMainServiceById 根据id获取spu的信息
*/
func (s ProductService) GetProductMainServiceById(ctx context.Context, id int) (product_param.ProductMainExtResponse, error) {
	productMainInfo, err := s.dao.GetProductMainDaoById(id)
	if err != nil {
		return product_param.ProductMainExtResponse{}, err
	}
	brandIdList := []int{productMainInfo.BrandId}
	categoryIdList := []int{productMainInfo.CategoryId}
	shopIdList := []int{productMainInfo.ShopId}

	image, _ := s.dao.QueryProductImageByProductId(productMainInfo.Id)
	stock, _ := s.dao.QueryProductStockByProductMainId(productMainInfo.Id)
	imageIdList := make([]int, 0)
	imageIdStringList := strings.Split(image.ImageIds, ",")
	for _, imageIdString := range imageIdStringList {
		imageId, err := strconv.Atoi(imageIdString)
		if err != nil {
			zap.L().Error("GetProductMainServiceById strconv.Atoi error", zap.Any("error", err))
			continue
		}
		imageIdList = append(imageIdList, imageId)
	}
	categoryChanMap := make(chan map[int]string, 1)
	imageChanMap := make(chan map[int]product_param.ImageParam, 1)
	shopChanMap := make(chan map[int]product_param.ShopParam, 1)
	brandChanMap := make(chan map[int]product_param.BrandParam, 1)
	go remote_rpc.GetCategoryListByCategoryIdsChannel(ctx, categoryIdList, categoryChanMap)
	go remote_rpc.GetImageListByImageIdsChannel(ctx, imageIdList, imageChanMap)
	go remote_rpc.GetShopListByShopIdsChannel(ctx, shopIdList, shopChanMap)
	go remote_rpc.GetBrandListByBrandIdsChannel(ctx, brandIdList, brandChanMap)
	categoryMap := <-categoryChanMap
	imageMap := <-imageChanMap
	shopMap := <-shopChanMap
	brandMap := <-brandChanMap
	categoryName := categoryMap[productMainInfo.CategoryId]
	shop, ok := shopMap[productMainInfo.ShopId]
	shopName := ""
	if ok {
		shopName = shop.ShopName
	}
	brand, ok := brandMap[productMainInfo.BrandId]
	brandName := ""
	if ok {
		brandName = brand.BrandName
	}
	imageMapList := make([]product_param.ImageParam, 0)
	for _, imageId := range imageIdList {
		data, ok := imageMap[imageId]
		if ok {
			imageMapList = append(imageMapList, data)
		}
	}
	return product_param.ProductMainExtResponse{
		ProductMainResponse:     productMainInfo,
		Stock:                   stock,
		ProductMainStatusString: product_util.GetProductMainStatusNameByCode(productMainInfo.ProductMainStatus),
		SaleTimeString:          productMainInfo.SaleTime.String(),
		ProductMainTypeName:     product_util.GetProductMainTypeNameByCode(productMainInfo.ProductMainType),
		BrandName:               brandName,
		CategoryName:            categoryName,
		ShopName:                shopName,
		Images:                  imageIdList,
		ImageMapList:            imageMapList,
	}, err
}

/*
	PatchProductMainServiceByParam 根据id更新spu的信息
*/
func (s ProductService) PatchProductMainServiceByParam(param product_param.PostProductMainRequestParam) error {
	imageIdList := make([]int, 0)
	if param.Images != "" {
		imageIdStringList := strings.Split(param.Images, ",")
		for _, imageIdString := range imageIdStringList {
			imageId, err := strconv.Atoi(imageIdString)
			if err != nil {
				zap.L().Error("PatchProductMainServiceByParam strconv.Atoi error", zap.Any("error", err))
				continue
			}
			imageIdList = append(imageIdList, imageId)
		}
	}
	return s.dao.PatchProductMainDaoByParam(product_param.PostProductMainExtRequestParam{
		PostProductMainRequestParam: param,
		ImageIdList:                 imageIdList,
	})
}

/*
	DeleteProductMainServiceById 根据id删除spu的信息
*/
func (s ProductService) DeleteProductMainServiceById(id int) error {
	return s.dao.DeleteProductMainDaoById(id)
}

/*
	PostProductServiceByParam 新建sku
*/
func (s ProductService) PostProductServiceByParam(param product_param.PostProductRequestParam) error {
	imageIdList := make([]int, 0)
	if param.Images != "" {
		imageIdStringList := strings.Split(param.Images, ",")
		for _, imageIdString := range imageIdStringList {
			imageId, err := strconv.Atoi(imageIdString)
			if err != nil {
				zap.L().Error("PatchProductMainServiceByParam strconv.Atoi error", zap.Any("error", err))
				continue
			}
			imageIdList = append(imageIdList, imageId)
		}
	}
	return s.dao.PostProductDaoByParam(product_param.PostProductExtRequestParam{
		PostProductRequestParam: param,
		ImageIdList:             imageIdList,
	})
}

/*
	GetProductServiceByParam 获取sku list
*/
func (s ProductService) GetProductServiceByParam(ctx context.Context, param product_param.GetProductRequestParam) (product_param.ProductListResponse, error) {
	total, data, err := s.dao.GetProductListDaoByParam(param)
	if err != nil {
		return product_param.ProductListResponse{}, err
	}
	if total == 0 {
		return product_param.ProductListResponse{}, errors.New(gin_util.NotFindTip)
	}
	productIdList := make([]int, 0)
	for _, item := range data {
		productIdList = append(productIdList, item.Id)
	}
	stockMap := make(map[int]int64, 0)
	stocks, _ := s.dao.QueryProductStockByProductIds(productIdList)
	for _, stock := range stocks {
		stockMap[stock.ProductId] = stock.StockTotal
	}
	imageMap := make(map[int]product_param.ImageParam, 0)
	images, _ := s.dao.QueryProductImageByProductIds(productIdList)
	imageMapList := make(map[int][]int, 0)
	if len(images) != 0 {
		imageIdAllList := make([]int, 0)
		for _, image := range images {
			imageIdList := make([]int, 0)
			imageIdStringList := strings.Split(image.ImageIds, ",")
			for _, imageIdString := range imageIdStringList {
				imageId, err := strconv.Atoi(imageIdString)
				if err != nil {
					zap.L().Error("GetProductServiceByParam strconv.Atoi error", zap.Any("error", err))
					continue
				}
				imageIdList = append(imageIdList, imageId)
				imageIdAllList = append(imageIdAllList, imageId)
			}
			imageMapList[image.ProductId] = imageIdList
		}

		imageChanMap := make(chan map[int]product_param.ImageParam, 1)
		go remote_rpc.GetImageListByImageIdsChannel(ctx, imageIdAllList, imageChanMap)
		imageMap = <-imageChanMap
	}
	list := make([]product_param.ProductExtResponse, 0)
	for _, item := range data {
		productExtResponse := product_param.ProductExtResponse{
			ProductResponse: item,
		}
		imageIdList := imageMapList[item.Id]
		productExtResponse.Images = imageIdList
		imageMapList := make([]product_param.ImageParam, 0)
		for _, imageId := range imageIdList {
			data, ok := imageMap[imageId]
			if ok {
				imageMapList = append(imageMapList, data)
			}
		}
		productExtResponse.ImageMapList = imageMapList
		productExtResponse.Stock = stockMap[item.Id]
		list = append(list, productExtResponse)
	}
	return product_param.ProductListResponse{
		Total: total,
		List:  list,
	}, err
}

/*
	GetProductServiceById 获取sku 信息
*/
func (s ProductService) GetProductServiceById(ctx context.Context, id int) (product_param.ProductExtResponse, error) {
	productInfo, err := s.dao.GetProductDaoById(id)
	if err != nil {
		return product_param.ProductExtResponse{}, err
	}
	stock, _ := s.dao.QueryProductStockByProductId(productInfo.Id)
	imageParam, _ := s.dao.QueryProductImageByProductId(productInfo.Id)
	imageIdStringList := strings.Split(imageParam.ImageIds, ",")
	imageIdList := make([]int, 0)
	for _, imageIdString := range imageIdStringList {
		imageId, err := strconv.Atoi(imageIdString)
		if err != nil {
			zap.L().Error("GetProductServiceById strconv.Atoi error", zap.Any("error", err))
			continue
		}
		imageIdList = append(imageIdList, imageId)
	}
	imageChanMap := make(chan map[int]product_param.ImageParam, 1)
	go remote_rpc.GetImageListByImageIdsChannel(ctx, imageIdList, imageChanMap)
	imageMap := <-imageChanMap
	productExtResponse := product_param.ProductExtResponse{
		ProductResponse: productInfo,
	}
	productExtResponse.Images = imageIdList
	imageMapList := make([]product_param.ImageParam, 0)
	for _, imageId := range imageIdList {
		data, ok := imageMap[imageId]
		if ok {
			imageMapList = append(imageMapList, data)
		}
	}
	productExtResponse.ImageMapList = imageMapList
	productExtResponse.Stock = stock
	return productExtResponse, nil
}

/*
	PatchProductServiceByParam 修改sku 信息
*/
func (s ProductService) PatchProductServiceByParam(param product_param.PostProductRequestParam) error {
	imageIdList := make([]int, 0)
	if param.Images != "" {
		imageIdStringList := strings.Split(param.Images, ",")
		for _, imageIdString := range imageIdStringList {
			imageId, err := strconv.Atoi(imageIdString)
			if err != nil {
				zap.L().Error("PatchProductMainServiceByParam strconv.Atoi error", zap.Any("error", err))
				continue
			}
			imageIdList = append(imageIdList, imageId)
		}
	}
	return s.dao.PatchProductDaoByParam(product_param.PostProductExtRequestParam{
		PostProductRequestParam: param,
		ImageIdList:             imageIdList,
	})
}

/*
	DeleteProductServiceById 根据id删除spu的信息
*/
func (s ProductService) DeleteProductServiceById(id int) error {
	return s.dao.DeleteProductDaoById(id)
}

/*
	GetProductStockListServiceByParam  获取库存
*/
func (s ProductService) GetProductStockListServiceByParam(param product_param.GetProductStockRequestParam) (product_param.ProductStockListResponse, error) {
	total, stockInfos, err := s.dao.GetProductStockListDaoByParam(param)
	if err != nil {
		return product_param.ProductStockListResponse{}, err
	}
	if total == 0 {
		return product_param.ProductStockListResponse{}, errors.New(gin_util.NotFindTip)
	}
	productIdList := make([]int, 0)
	productMainIdList := make([]int, 0)
	for _, stock := range stockInfos {
		productIdList = append(productIdList, stock.ProductId)
		productMainIdList = append(productMainIdList, stock.ProductMainId)
	}
	productInfoMap := make(map[int]string, 0)
	productMainInfoMap := make(map[int]string, 0)
	productInfos, _ := s.dao.QueryProductListDaoByIds(productIdList)
	productMainInfos, _ := s.dao.QueryProductMainListDaoByIds(productMainIdList)
	for _, productInfo := range productInfos {
		productInfoMap[productInfo.Id] = productInfo.Title
	}
	for _, productMainInfo := range productMainInfos {
		productMainInfoMap[productMainInfo.Id] = productMainInfo.Title
	}
	list := make([]product_param.ProductStockResponse, 0)
	for _, stock := range stockInfos {
		list = append(list, product_param.ProductStockResponse{
			Id:              stock.Id,
			ProductMainId:   stock.ProductMainId,
			ProductId:       stock.ProductId,
			ProductName:     productInfoMap[stock.ProductId],
			ProductMainName: productMainInfoMap[stock.ProductMainId],
			StockNumber:     stock.StockNumber,
		})
	}
	return product_param.ProductStockListResponse{
		Total: total,
		List:  list,
	}, nil
}

/*
	PostProductStockServiceByParam  新建库存
*/
func (s ProductService) PostProductStockServiceByParam(param product_param.PostProductStockRequestParam) error {
	return s.dao.PostProductStockDaoByParam(param)
}

/*
	PatchProductStockServiceByParam  更新库存
*/
func (s ProductService) PatchProductStockServiceByParam(param product_param.PostProductStockRequestParam) error {
	return s.dao.PatchProductStockDaoByParam(param)
}

/*
	DeleteProductStockServiceByParam  删除库存
*/
func (s ProductService) DeleteProductStockServiceById(id int) error {
	return s.dao.DeleteProductStockDaoById(id)
}

/*
	DeleteProductStockByIdServiceById  批量删除库存
*/
func (s ProductService) DeleteProductStockByIdServiceById(param product_param.PostProductStockByIdsRequestParam) error {
	if param.Ids != "" {
		idStringList := strings.Split(param.Ids, ",")
		for _, idString := range idStringList {
			id, err := strconv.Atoi(idString)
			if err != nil {
				zap.L().Error("DeleteProductStockByIdServiceById Ids strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
				continue
			}
			param.IdList = append(param.IdList, id)
		}
	}
	if param.ProductIds != "" {
		idStringList := strings.Split(param.ProductIds, ",")
		for _, idString := range idStringList {
			id, err := strconv.Atoi(idString)
			if err != nil {
				zap.L().Error("DeleteProductStockByIdServiceById ProductIds strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
				continue
			}
			param.ProductIdList = append(param.ProductIdList, id)
		}
	}
	if param.ProductMainIds != "" {
		idStringList := strings.Split(param.ProductMainIds, ",")
		for _, idString := range idStringList {
			id, err := strconv.Atoi(idString)
			if err != nil {
				zap.L().Error("DeleteProductStockByIdServiceById ProductMainIds strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
				continue
			}
			param.ProductMainIdList = append(param.ProductMainIdList, id)
		}
	}
	if len(param.IdList) == 0 && len(param.ProductIdList) == 0 && len(param.ProductMainIdList) == 0 {
		return errors.New(product_const.ParamEmptyTip)
	}
	return s.dao.DeleteProductStockByIdDaoByParam(param)
}

/*
	GetProductStockByIdServiceById  根据 id product_id  product_main_id 批量查询库存
*/
func (s ProductService) GetProductStockByIdServiceById(param product_param.PostProductStockByIdsRequestParam) ([]product_param.ProductStockResponse, error) {
	if param.Ids != "" {
		idStringList := strings.Split(param.Ids, ",")
		for _, idString := range idStringList {
			id, err := strconv.Atoi(idString)
			if err != nil {
				zap.L().Error("GetProductStockByIdServiceById Ids strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
				continue
			}
			param.IdList = append(param.IdList, id)
		}
	}
	if param.ProductIds != "" {
		idStringList := strings.Split(param.ProductIds, ",")
		for _, idString := range idStringList {
			id, err := strconv.Atoi(idString)
			if err != nil {
				zap.L().Error("GetProductStockByIdServiceById ProductIds strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
				continue
			}
			param.ProductIdList = append(param.ProductIdList, id)
		}
	}
	if param.ProductMainIds != "" {
		idStringList := strings.Split(param.ProductMainIds, ",")
		for _, idString := range idStringList {
			id, err := strconv.Atoi(idString)
			if err != nil {
				zap.L().Error("GetProductStockByIdServiceById ProductMainIds strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
				continue
			}
			param.ProductMainIdList = append(param.ProductMainIdList, id)
		}
	}
	if len(param.IdList) == 0 && len(param.ProductIdList) == 0 && len(param.ProductMainIdList) == 0 {
		return nil, errors.New(product_const.ParamEmptyTip)
	}
	stockInfos, err := s.dao.GetProductStockByIdDaoByParam(param)
	if err != nil {
		return nil, err
	}
	productIdList := make([]int, 0)
	productMainIdList := make([]int, 0)
	for _, stock := range stockInfos {
		productIdList = append(productIdList, stock.ProductId)
		productMainIdList = append(productMainIdList, stock.ProductMainId)
	}
	productInfoMap := make(map[int]string, 0)
	productMainInfoMap := make(map[int]string, 0)
	productInfos, _ := s.dao.QueryProductListDaoByIds(productIdList)
	productMainInfos, _ := s.dao.QueryProductMainListDaoByIds(productMainIdList)
	for _, productInfo := range productInfos {
		productInfoMap[productInfo.Id] = productInfo.Title
	}
	for _, productMainInfo := range productMainInfos {
		productMainInfoMap[productMainInfo.Id] = productMainInfo.Title
	}
	list := make([]product_param.ProductStockResponse, 0)
	for _, stock := range stockInfos {
		list = append(list, product_param.ProductStockResponse{
			Id:              stock.Id,
			ProductMainId:   stock.ProductMainId,
			ProductId:       stock.ProductId,
			ProductName:     productInfoMap[stock.ProductId],
			ProductMainName: productMainInfoMap[stock.ProductMainId],
			StockNumber:     stock.StockNumber,
		})
	}
	return list, nil
}

/*
	GetProductByProductMainIdsServiceByParam  根据ids product_main_ids 查询product
*/
func (s ProductService) GetProductByProductMainIdsServiceByParam(ctx context.Context, param product_param.PostProductIdsRequestParam) ([]product_param.ProductExtResponse, error) {
	if param.Ids != "" {
		idStringList := strings.Split(param.Ids, ",")
		for _, idString := range idStringList {
			id, err := strconv.Atoi(idString)
			if err != nil {
				zap.L().Error("GetProductByProductMainIdsServiceByParam Ids strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
				continue
			}
			param.IdList = append(param.IdList, id)
		}
	}
	if param.ProductMainIds != "" {
		idStringList := strings.Split(param.ProductMainIds, ",")
		for _, idString := range idStringList {
			id, err := strconv.Atoi(idString)
			if err != nil {
				zap.L().Error("GetProductByProductMainIdsServiceByParam ProductMainIds strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
				continue
			}
			param.ProductMainIdList = append(param.ProductMainIdList, id)
		}
	}
	if len(param.IdList) == 0 && len(param.ProductMainIdList) == 0 {
		return nil, errors.New(product_const.ParamEmptyTip)
	}
	data, err := s.dao.QueryProductListDaoByParam(param)
	if err != nil {
		return nil, err
	}
	productIdList := make([]int, 0)
	for _, item := range data {
		productIdList = append(productIdList, item.Id)
	}
	stockMap := make(map[int]int64, 0)
	stocks, _ := s.dao.QueryProductStockByProductIds(productIdList)
	for _, stock := range stocks {
		stockMap[stock.ProductId] = stock.StockTotal
	}
	imageMap := make(map[int]product_param.ImageParam, 0)
	images, _ := s.dao.QueryProductImageByProductIds(productIdList)
	imageMapList := make(map[int][]int, 0)
	if len(images) != 0 {
		imageIdAllList := make([]int, 0)
		for _, image := range images {
			imageIdList := make([]int, 0)
			imageIdStringList := strings.Split(image.ImageIds, ",")
			for _, imageIdString := range imageIdStringList {
				imageId, err := strconv.Atoi(imageIdString)
				if err != nil {
					zap.L().Error("GetProductByProductMainIdsServiceByParam strconv.Atoi error", zap.Any("error", err))
					continue
				}
				imageIdList = append(imageIdList, imageId)
				imageIdAllList = append(imageIdAllList, imageId)
			}
			imageMapList[image.ProductId] = imageIdList
		}

		imageChanMap := make(chan map[int]product_param.ImageParam, 1)
		go remote_rpc.GetImageListByImageIdsChannel(ctx, imageIdAllList, imageChanMap)
		imageMap = <-imageChanMap
	}
	list := make([]product_param.ProductExtResponse, 0)
	for _, item := range data {
		productExtResponse := product_param.ProductExtResponse{
			ProductResponse: item,
		}
		imageIdList := imageMapList[item.Id]
		productExtResponse.Images = imageIdList
		imageMapList := make([]product_param.ImageParam, 0)
		for _, imageId := range imageIdList {
			data, ok := imageMap[imageId]
			if ok {
				imageMapList = append(imageMapList, data)
			}
		}
		productExtResponse.ImageMapList = imageMapList
		productExtResponse.Stock = stockMap[item.Id]
		list = append(list, productExtResponse)
	}
	return list, nil
}

/*
	PostDeleteProductMainAllServiceByParam  根据ids product_main_ids 删除product_main product product_image product_stock
*/
func (s ProductService) PostDeleteProductMainAllServiceByParam(param product_param.PostDeleteProductMainAllRequestParam) error {
	idStringList := strings.Split(param.Ids, ",")
	for _, idString := range idStringList {
		id, err := strconv.Atoi(idString)
		if err != nil {
			zap.L().Error("PostDeleteProductMainAllServiceByParam Ids strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
			continue
		}
		param.IdList = append(param.IdList, id)
	}
	err := s.dao.DeleteProductMainDaoByIds(param.IdList)
	if err != nil {
		return err
	}
	productInfos, _ := s.dao.QueryProductListDaoByProductMainIds(param.IdList)
	if len(productInfos) != 0 {
		productIdList := make([]int, 0)
		for _, product := range productInfos {
			productIdList = append(productIdList, product.Id)
		}
		_ = s.dao.DeleteProductImageByProductIds(productIdList)
	}
	_ = s.dao.DeleteProductImageByProductMainIds(param.IdList)
	_ = s.dao.DeleteProductStockDaoByProductMainIds(param.IdList)
	return nil
}

/*
	PostDeleteProductServiceByParam  根据ids productIds 删除 product product_image product_stock
*/
func (s ProductService) PostDeleteProductServiceByParam(param product_param.DeletePostProductIdsRequestParam) error {
	idStringList := strings.Split(param.Ids, ",")
	for _, idString := range idStringList {
		id, err := strconv.Atoi(idString)
		if err != nil {
			zap.L().Error("PostDeleteProductServiceByParam Ids strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
			continue
		}
		param.IdList = append(param.IdList, id)
	}
	err := s.dao.DeleteProductDaoByIds(param.IdList)
	if err != nil {
		return err
	}
	_ = s.dao.DeleteProductImageByProductIds(param.IdList)
	_ = s.dao.DeleteProductStockDaoByProductIds(param.IdList)
	return nil
}

/*
	PostProductMainToEsServiceByParam  根据product_main_ids 发送到es
*/
func (s ProductService) PostProductMainToMqServiceByParam(param product_param.PostProductIdsToMqRequestParam) (int64, error) {
	idStringList := strings.Split(param.Ids, ",")
	for _, idString := range idStringList {
		id, err := strconv.Atoi(idString)
		if err != nil {
			zap.L().Error("PostProductMainToEsServiceByParam Ids strconv.Atoi error", zap.Any("idString", idString), zap.Any("error", err))
			continue
		}
		param.IdList = append(param.IdList, id)
	}
	return product_mq.SendProductMainToMq(param.IdList)

}
