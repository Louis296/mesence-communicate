package kafka

import (
	"context"
	"fmt"
	"github.com/louis296/mesence-communicate/pkg/pb"
	"testing"
)

func TestNewProducer(t *testing.T) {
	p, err := NewProducer("127.0.0.1:9092", "test")
	if err != nil {
		fmt.Print(err)
		t.FailNow()
	}
	_, _, err = p.SendMessage(context.Background(), "test", &pb.Msg{
		Type: pb.Type_Word,
		Data: &pb.Data{Content: "test_kafka_2"},
	})
	if err != nil {
		fmt.Print(err)
		t.FailNow()
	}
}
