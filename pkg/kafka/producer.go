package kafka

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/louis296/mesence-communicate/pkg/log"
)

var ProducerClient *Producer

func InitProducer(url, topic string) (err error) {
	ProducerClient, err = NewProducer(url, topic)
	return
}

type Producer struct {
	topic    string
	config   *sarama.Config
	producer sarama.SyncProducer
}

func NewProducer(url, topic string) (*Producer, error) {
	p := &Producer{topic: topic, config: sarama.NewConfig()}
	p.config.Producer.RequiredAcks = sarama.WaitForAll
	p.config.Producer.Partitioner = sarama.NewHashPartitioner
	p.config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{url}, p.config)
	if err != nil {
		return nil, err
	}
	p.producer = client
	return p, nil
}

func (p *Producer) SendMessage(c context.Context, key string, msg proto.Message) (int32, int64, error) {
	log.Info("send message to kafka msg = [%v] ", msg)
	kMsg := &sarama.ProducerMessage{}
	kMsg.Topic = p.topic
	kMsg.Key = sarama.StringEncoder(key)
	bs, err := proto.Marshal(msg)
	if err != nil {
		return 0, 0, err
	}
	kMsg.Value = sarama.ByteEncoder(bs)

	partition, offset, err := p.producer.SendMessage(kMsg)
	if err != nil {
		return 0, 0, err
	}
	return partition, offset, nil
}
