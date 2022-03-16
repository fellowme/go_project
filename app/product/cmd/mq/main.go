package main

import (
	"github.com/apache/pulsar-client-go/pulsar"
	gin_app "github.com/fellowme/gin_common_library/app"
	gin_pulsar "github.com/fellowme/gin_common_library/mq"
	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"
	"go_project/app/product/product_const"
	"go_project/app/product/product_es"
	"time"
)

func registerCustomerMq() {
	service := product_es.GetProductEsService()
	stopChan := make(chan error)
	// 初始化 协程池
	pool, err := ants.NewPool(product_const.GoroutinePoolSize, ants.WithPreAlloc(true),
		ants.WithMaxBlockingTasks(product_const.MaxBlockingTasks), ants.WithNonblocking(true))
	if err != nil {
		zap.L().Error("registerCustomerMq NewPool error", zap.Any("error", err))
		return
	}
	// 监听消息处理消息
	go gin_pulsar.ReceivePulsarMqMessage(pulsar.ConsumerOptions{
		Topic:            product_const.ProductMainTopic,
		SubscriptionName: product_const.ProductMainConsumerName,
		Type:             1,
	}, service.SendProductMainToEs, stopChan, pool)

	go gin_pulsar.ReceivePulsarMqMessage(pulsar.ConsumerOptions{
		Topic:            product_const.ProductTopic,
		SubscriptionName: product_const.ProductConsumerName,
		Type:             1,
	}, service.SendProductToEs, stopChan, pool)
	// 阻塞
	select {
	case pulsarError := <-stopChan:
		zap.L().Error("registerCustomerMq error", zap.Any("error", pulsarError))
	}
	//  协程池 执行的func 为0的状态 执行后续操作
	for {
		if pool.Running() == 0 {
			break
		}
		time.Sleep(10 * time.Second)
	}
	//  释放 pool
	pool.Release()
	close(stopChan)

}

/*
	主程序
*/
func main() {
	gin_app.CreateAppMqServer("/app/product/product_config/", "go_product_mq", registerCustomerMq)
}
