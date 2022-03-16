package product_es

import (
	"context"
	"encoding/json"
	"github.com/apache/pulsar-client-go/pulsar"
	gin_es "github.com/fellowme/gin_common_library/elastic"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"go_project/app/product/product_const"
	"go_project/app/product/product_param"
	"go_project/app/product/product_remote_service/remote_rpc"
	"strconv"
	"strings"
)

func (p ProductEsService) SendProductToEs(message pulsar.Message) {
	zap.L().Info("SendProductToEs message", zap.Any("message", message))
	if len(message.Payload()) != 0 {
		var ids []int
		if err := json.Unmarshal(message.Payload(), &ids); err != nil {
			zap.L().Error("SendProductToEs json.Unmarsha error", zap.Any("error", err), zap.Any("message", message))
			return
		}
		productInfoList, err := p.dao.QueryProductListDaoByIds(ids)
		if err != nil {
			zap.L().Error("SendProductToEs QueryProductListDaoByIds error", zap.Any("error", err), zap.Any("message", message))
			return
		}
		if len(productInfoList) == 0 {
			zap.L().Error("SendProductToEs QueryProductListDaoByIds empty", zap.Any("message", message))
			return
		}
		productIdList := make([]int, 0)
		for _, item := range productInfoList {
			productIdList = append(productIdList, item.Id)
		}
		stockMap := make(map[int]int64, 0)
		stocks, _ := p.dao.QueryProductStockByProductIds(productIdList)
		for _, stock := range stocks {
			stockMap[stock.ProductId] = stock.StockTotal
		}
		ctx := context.Background()
		imageMap := make(map[int]product_param.ImageParam, 0)
		images, _ := p.dao.QueryProductImageByProductIds(productIdList)
		imageMapList := make(map[int][]int, 0)
		if len(images) != 0 {
			imageIdAllList := make([]int, 0)
			for _, image := range images {
				imageIdList := make([]int, 0)
				imageIdStringList := strings.Split(image.ImageIds, ",")
				for _, imageIdString := range imageIdStringList {
					imageId, err := strconv.Atoi(imageIdString)
					if err != nil {
						zap.L().Error("SendProductToEs strconv.Atoi error", zap.Any("error", err))
						continue
					}
					imageIdList = append(imageIdList, imageId)
					imageIdAllList = append(imageIdAllList, imageId)
				}
				imageMapList[image.ProductId] = imageIdList
			}

			imageChanMap := make(chan map[int]product_param.ImageParam, 1)
			go remote_rpc.GetImageListByImageIdsChannel(nil, imageIdAllList, imageChanMap)
			imageMap = <-imageChanMap
		}

		list := make([]product_param.ProductExtResponse, 0)
		for _, item := range productInfoList {
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
		es := gin_es.GetElasticClient()
		req := es.Bulk().Index(product_const.ProductIndex)
		for _, product := range list {
			doc := elastic.NewBulkIndexRequest().Id(strconv.Itoa(product.Id)).Doc(product)
			req.Add(doc)
		}
		if req.NumberOfActions() < 0 {
			zap.L().Info("SendProductToEs NewBulkIndexRequest empty")
			return
		}
		if _, err := req.Do(ctx); err != nil {
			zap.L().Error("SendProductToEs req NewBulkIndexRequest error ", zap.Any("error", err))
		}
	}
}
