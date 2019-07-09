package kafka

import (
	"github.com/astaxie/beego/logs"
	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer
)

func InitKafka(addr string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logs.Error("init kafka failed, err:", err)
		return
	}
	return
}

func SendMsgToKafka(data, topic string) (err error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed, err:%v, topic:%v, data:%v", err, topic, data)
		return
	}
	logs.Debug("send msg success, pid:%v, offset:%v, topic:%v", pid, offset, topic)
	return
}