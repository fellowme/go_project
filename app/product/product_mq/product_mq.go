package product_mq

import (
	"encoding/json"
	"github.com/apache/pulsar-client-go/pulsar"
	gin_pulsar "github.com/fellowme/gin_common_library/mq"
	"go_project/app/product/product_const"
)

func SendProductMainToMq(ids []int) (int64, error) {
	idsByte, _ := json.Marshal(ids)
	messageId, err := gin_pulsar.SendPulsarMqMessage(pulsar.ProducerOptions{
		Topic: product_const.ProductMainTopic,
	}, pulsar.ProducerMessage{
		Payload: idsByte,
	})
	return messageId.LedgerID(), err
}
