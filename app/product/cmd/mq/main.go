package main

import (
	"github.com/apache/pulsar-client-go/pulsar"
	gin_app "github.com/fellowme/gin_common_library/app"
	gin_pulsar "github.com/fellowme/gin_common_library/mq"
	"go.uber.org/zap"
	"go_project/app/product/product_const"
	"go_project/app/product/product_es"
)

func registerCustomerMq() {
	service := product_es.GetProductEsService()
	stopChan := make(chan error)
	go gin_pulsar.ReceivePulsarMqMessage(pulsar.ConsumerOptions{
		Topic:            product_const.ProductMainTopic,
		SubscriptionName: product_const.ProductMainConsumerName,
		Type:             1,
	}, service.SendProductMainToEs, stopChan)
	select {
	case err := <-stopChan:
		zap.L().Error("registerCustomerMq error", zap.Any("error", err))
	}
}

/*
	主程序
*/
func main() {
	gin_app.CreateAppMqServer("/app/product/product_config/", "go_product_mq", registerCustomerMq)
}
