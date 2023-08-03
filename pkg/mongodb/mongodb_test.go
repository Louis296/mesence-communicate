package mongodb

import (
	"github.com/louis296/mesence-communicate/pkg/pb"
	"testing"
	"time"
)

func TestSaveMessage(t *testing.T) {
	err := InitClient("127.0.0.1:27017")
	if err != nil {
		t.FailNow()
	}
	err = SaveMessage(&pb.Msg{
		Type: pb.Type_Word,
		Data: &pb.Data{
			To:       "test",
			From:     "test",
			Content:  "test",
			SendTime: time.Now().Format(time.RFC3339),
		},
	})
	if err != nil {
		t.FailNow()
	}
}
