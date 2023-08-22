package communicate_service

import (
	"context"
	"fmt"
	"github.com/louis296/mesence-communicate/dao"
	"github.com/louis296/mesence-communicate/pkg/kafka"
	"github.com/louis296/mesence-communicate/pkg/log"
	"github.com/louis296/mesence-communicate/pkg/pb"
	"github.com/louis296/mesence-communicate/pkg/redis_client"
	"github.com/louis296/mesence-communicate/pkg/util"
	"github.com/louis296/mesence-communicate/pkg/ws"
)

func HandleMessage(msg *pb.Msg, conn *ws.UserConn) {
	var err error
	var resp *pb.Msg
	switch msg.Type {
	case pb.Type_Word:
		resp, err = handleWord(msg)
	case pb.Type_FriendRequest:
		resp, err = handleFriendRequest(msg)
	case pb.Type_GetMaxSeq:
		resp, err = handleGetMaxSeq(msg)
	case pb.Type_PullMessage:
		resp, err = handlePullMessage(msg)
	case pb.Type_Offer:
		fallthrough
	case pb.Type_Answer:
		fallthrough
	case pb.Type_Candidate:
		resp, err = doTransfer(msg)
	}
	doReply(conn, resp, err)
}

func handleWord(msg *pb.Msg) (*pb.Msg, error) {
	_, _, err := kafka.ProducerClient.SendMessage(context.Background(), GenConversationKey(msg.Data.From, msg.Data.To), msg)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func handleFriendRequest(msg *pb.Msg) (*pb.Msg, error) {
	_, _, err := kafka.ProducerClient.SendMessage(context.Background(), GenConversationKey(msg.Data.From, msg.Data.To), msg)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func handleGetMaxSeq(msg *pb.Msg) (*pb.Msg, error) {
	seq, err := redis_client.GetMaxSeq(GenConversationKey(msg.Data.From, msg.Data.To))
	if err != nil {
		return msg, err
	}
	return &pb.Msg{
		Type: pb.Type_GetMaxSeq,
		Data: msg.Data,
		Seq:  seq,
	}, nil
}

func handlePullMessage(msg *pb.Msg) (*pb.Msg, error) {
	_, _, err := kafka.ProducerClient.SendMessage(context.Background(), GenConversationKey(msg.Data.From, msg.Data.To), msg)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func doTransfer(msg *pb.Msg) (*pb.Msg, error) {
	PushMessage(msg, msg.Data.To)
	return msg, nil
}

func PushMessage(msg *pb.Msg, userId string) {
	if item, ok := ws.UserConnMap.Load(userId); ok {
		receiverConn := item.(*ws.UserConn)
		err := receiverConn.Send(util.Marshal(msg))
		if err != nil {
			log.Error("Push message to user [%v] error", receiverConn.UserPhone)
		}
	} else {
		log.Warn("receiver [%v] is offline, push message abort", msg.Data.To)
	}
}

func OnlineNotify(userPhone string) error {
	friends, err := dao.GetFriendRelationsByUserPhone(userPhone)
	if err != nil {
		return err
	}
	message := &pb.Msg{
		Type: pb.Type_Online,
		Data: &pb.Data{OnlineUsers: []string{userPhone}},
	}
	for _, friend := range friends {
		if item, ok := ws.UserConnMap.Load(friend.UserPhone); ok {
			userConn := item.(*ws.UserConn)
			if err := userConn.Send(util.Marshal(message)); err != nil {
				log.Error("Send online notify to user [%v] error", userPhone)
			}
		}
	}
	return nil
}

func OfflineNotify(userPhone string) error {
	friends, err := dao.GetFriendRelationsByUserPhone(userPhone)
	if err != nil {
		return err
	}
	message := &pb.Msg{
		Type: pb.Type_Offline,
		Data: &pb.Data{OnlineUsers: []string{userPhone}},
	}
	for _, friend := range friends {
		if item, ok := ws.UserConnMap.Load(friend.UserPhone); ok {
			userConn := item.(*ws.UserConn)
			if err := userConn.Send(util.Marshal(message)); err != nil {
				log.Error("Send offline notify to user [%v] error", userPhone)
			}
		}
	}
	return nil
}

func doReply(conn *ws.UserConn, msg *pb.Msg, err error) {
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	resp := &pb.Resp{
		ErrStr: errStr,
		Msg:    msg,
	}
	conn.Send(util.Marshal(resp))
}

func GenConversationKey(from, to string) string {
	return fmt.Sprintf("conversation_%v_%v", from, to)
}
