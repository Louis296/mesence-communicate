package communicate_service

import (
	"errors"
	"fmt"
	"github.com/louis296/mesence-communicate/dao"
	"github.com/louis296/mesence-communicate/dao/model"
	"github.com/louis296/mesence-communicate/pkg/log"
	"github.com/louis296/mesence-communicate/pkg/mongodb"
	"github.com/louis296/mesence-communicate/pkg/pb"
	"github.com/louis296/mesence-communicate/pkg/redis_client"
	"github.com/louis296/mesence-communicate/pkg/util"
	"github.com/louis296/mesence-communicate/pkg/ws"
	"gorm.io/gorm"
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
	data := msg.Data

	// check if receiver is valid
	_, err := dao.GetUserByPhone(data.To)
	if err != nil {
		log.Error("No user [%v] or mongodb error, send word message abort", data.To)
		return msg, err
	}

	// store message
	err = mongodb.SaveMessage(msg)
	if err != nil {
		log.Error("Save message error")
		return msg, err
	}

	// try to send message
	if item, ok := ws.UserConnMap.Load(data.To); ok {
		receiverConn := item.(*ws.UserConn)
		sendMessage := &pb.Msg{Type: pb.Type_Word}
		sendMessage.Data = data
		err = receiverConn.Send(util.Marshal(sendMessage))
		if err != nil {
			log.Error("Send word message to user [%v] error", receiverConn.UserPhone)
		}
	} else {
		log.Warn("Word message receiver [%v] is offline, send abort", data.To)
	}
	return msg, nil
}

func handleFriendRequest(msg *pb.Msg) (*pb.Msg, error) {
	data := msg.Data

	// check if candidate is valid
	_, err := dao.GetUserByPhone(data.To)
	if err != nil {
		log.Error("No user [%v] or db error, send word message abort", data.Candidate)
		return msg, err
	}

	// check if friend relation already exist
	_, err = dao.GetFriendRelationByUserAndFriend(data.From, data.To)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Friend relation already exist")
		return msg, err
	}

	// check if friend request already exist and not finish
	_, err = dao.GetFriendRequestBySenderAndCandidateAndStatus(data.From, data.To, 2)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Friend request already exist")
		return msg, err
	}

	// store friend request
	friendRequest := model.FriendRequest{
		Sender:        data.From,
		Candidate:     data.To,
		RequestStatus: int(pb.RequestStatus_Waiting),
		StartTime:     util.TimeParse(data.SendTime),
		Content:       data.Content,
	}
	tx := dao.DB.Begin()
	defer func() {
		if tx.Error != nil {
			tx.Rollback()
		} else if recover() != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if err = dao.CreateFriendRequest(tx, friendRequest); err != nil {
		log.Error("Store request to mongodb error")
		return msg, err
	}

	// try to send friend request notice
	if item, ok := ws.UserConnMap.Load(data.To); ok {
		candidateConn := item.(*ws.UserConn)
		notice := &pb.Msg{
			Type: pb.Type_FriendRequest,
			Data: &pb.Data{
				From:          data.From,
				Content:       data.Content,
				SendTime:      data.SendTime,
				To:            data.To,
				RequestStatus: pb.RequestStatus_Waiting,
			},
		}
		err = candidateConn.Send(util.Marshal(notice))
		if err != nil {
			log.Error("Send word message to user [%v] error", candidateConn.UserPhone)
		}
	} else {
		log.Warn("Friend request receiver [%v] is offline, send abort", data.Candidate)
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

func doTransfer(msg *pb.Msg) (*pb.Msg, error) {
	if item, ok := ws.UserConnMap.Load(msg.Data.To); ok {
		receiverConn := item.(*ws.UserConn)
		err := receiverConn.Send(util.Marshal(msg))
		if err != nil {
			log.Error("Send word message to user [%v] error", receiverConn.UserPhone)
			return msg, err
		}
	} else {
		log.Warn("Word message receiver [%v] is offline, send abort", msg.Data.To)
	}
	return msg, nil
}

func PushMessage(msg *pb.Msg) {
	if item, ok := ws.UserConnMap.Load(msg.Data.To); ok {
		receiverConn := item.(*ws.UserConn)
		err := receiverConn.Send(util.Marshal(msg))
		if err != nil {
			log.Error("Send word message to user [%v] error", receiverConn.UserPhone)
		}
	} else {
		log.Warn("Word message receiver [%v] is offline, send abort", msg.Data.To)
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
