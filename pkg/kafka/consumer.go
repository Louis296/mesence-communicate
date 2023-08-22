package kafka

import (
	"context"
	"github.com/IBM/sarama"
)

func InitConsumer(handler sarama.ConsumerGroupHandler, url, topic string) error {
	config := sarama.NewConfig()
	group, err := sarama.NewConsumerGroup([]string{url}, "msg", config)
	if err != nil {
		return err
	}
	go func() {
		for {
			group.Consume(context.Background(), []string{topic}, handler)
		}
	}()
	return nil
}
