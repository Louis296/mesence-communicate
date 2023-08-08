package communicate_service

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/louis296/mesence-communicate/dao"
	"github.com/louis296/mesence-communicate/dao/model"
	"github.com/louis296/mesence-communicate/pkg/log"
	"github.com/louis296/mesence-communicate/pkg/mongodb"
	"github.com/louis296/mesence-communicate/pkg/pb"
	"github.com/louis296/mesence-communicate/pkg/util"
	"github.com/louis296/mesence-communicate/pkg/ws"
	"gorm.io/gorm"
)

func UserConnHandler(conn *ws.UserConn) {
	log.Info("Conn by user [%v] open", conn.UserPhone)

	//todo: check if user already online, and force offline the online user

	// store user conn
	ws.UserConnMap.Store(conn.UserPhone, conn)

	// send online notice to friends
	friendPhones, err := getFriendPhones(conn)
	if err != nil {
		log.Error("Search user friends error, online notify abort")
	}
	onlineNotify(conn, friendPhones)

	// handler message event
	conn.On("message", func(bs []byte) {
		msg := &pb.Msg{}
		if err := proto.Unmarshal(bs, msg); err != nil {
			log.Error("Cannot unmarshal message from user [%v], message: %v", conn.UserPhone, bs)
			return
		}
		switch msg.Type {
		case pb.Type_Word:
			onWord(conn, msg)
		case pb.Type_FriendRequest:
			onFriendRequest(conn, msg)
		case pb.Type_Offer:
			fallthrough
		case pb.Type_Answer:
			fallthrough
		case pb.Type_Candidate:
			onTransfer(conn, msg)
		}
	})

	// handler close event
	conn.On("close", func(code int, text string) {
		// send offline notice to friends
		friendPhones, err := getFriendPhones(conn)
		if err != nil {
			log.Error("Search user friends error, offline notify abort")
			return
		}
		offlineNotify(conn.UserPhone, friendPhones)
		ws.UserConnPool.Put(conn)
	})
}

func onWord(conn *ws.UserConn, msg *pb.Msg) {
	data := msg.Data

	// check if receiver is valid
	_, err := dao.GetUserByPhone(data.To)
	if err != nil {
		log.Error("No user [%v] or mongodb error, send word message abort", data.To)
		return
	}

	// store message
	err = mongodb.SaveMessage(msg)
	if err != nil {
		log.Error("Save message error")
		return
	}

	// try to send message
	if item, ok := ws.UserConnMap.Load(data.To); ok {
		receiverConn := item.(*ws.UserConn)
		sendMessage := &pb.Msg{Type: pb.Type_Word}
		data.From = conn.UserPhone
		sendMessage.Data = data
		err = receiverConn.Send(util.Marshal(sendMessage))
		if err != nil {
			log.Error("Send word message to user [%v] error", receiverConn.UserPhone)
		}
	} else {
		log.Warn("Word message receiver [%v] is offline, send abort", data.To)
	}
}

func onFriendRequest(conn *ws.UserConn, msg *pb.Msg) {
	data := msg.Data

	// check if candidate is valid
	_, err := dao.GetUserByPhone(data.To)
	if err != nil {
		log.Error("No user [%v] or mongodb error, send word message abort", data.Candidate)
		return
	}

	// check if friend relation already exist
	_, err = dao.GetFriendRelationByUserAndFriend(conn.UserPhone, data.To)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Friend relation already exist")
		return
	}

	// check if friend request already exist and not finish
	_, err = dao.GetFriendRequestBySenderAndCandidateAndStatus(conn.UserPhone, data.To, 2)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Friend request already exist")
		return
	}

	// store friend request
	friendRequest := model.FriendRequest{
		Sender:        conn.UserPhone,
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
		return
	}

	// try to send friend request notice
	if item, ok := ws.UserConnMap.Load(data.To); ok {
		candidateConn := item.(*ws.UserConn)
		notice := &pb.Msg{
			Type: pb.Type_FriendRequest,
			Data: &pb.Data{
				From:          conn.UserPhone,
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
}

func onTransfer(conn *ws.UserConn, msg *pb.Msg) {
	if item, ok := ws.UserConnMap.Load(msg.Data.To); ok {
		receiverConn := item.(*ws.UserConn)
		msg.Data.From = conn.UserPhone
		err := receiverConn.Send(util.Marshal(msg))
		if err != nil {
			log.Error("Send word message to user [%v] error", receiverConn.UserPhone)
		}
	} else {
		log.Warn("Word message receiver [%v] is offline, send abort", msg.Data.To)
	}
}

func onlineNotify(conn *ws.UserConn, userPhones []string) {
	message := &pb.Msg{
		Type: pb.Type_Online,
		Data: &pb.Data{OnlineUsers: []string{conn.UserPhone}},
	}
	for _, userPhone := range userPhones {
		if item, ok := ws.UserConnMap.Load(userPhone); ok {
			userConn := item.(*ws.UserConn)
			if err := userConn.Send(util.Marshal(message)); err != nil {
				log.Error("Send online notify to user [%v] error", userPhone)
			}
		}
	}
}

func offlineNotify(userPhone string, userPhones []string) {
	message := &pb.Msg{
		Type: pb.Type_Offline,
		Data: &pb.Data{OfflineUsers: []string{userPhone}},
	}
	for _, userPhone := range userPhones {
		if item, ok := ws.UserConnMap.Load(userPhone); ok {
			userConn := item.(*ws.UserConn)
			if err := userConn.Send(util.Marshal(message)); err != nil {
				log.Error("Send offline notify to user [%v] error", userPhone)
			}
		}
	}
}

func getFriendPhones(conn *ws.UserConn) ([]string, error) {
	friendRelations, err := dao.GetFriendRelationsByUserPhone(conn.UserPhone)
	if err != nil {
		return nil, err
	}
	var friendPhones []string
	for _, relation := range friendRelations {
		friendPhones = append(friendPhones, relation.FriendPhone)
	}
	return friendPhones, nil
}
