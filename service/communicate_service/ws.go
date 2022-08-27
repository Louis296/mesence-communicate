package communicate_service

import (
	"encoding/json"
	"github.com/louis296/mesence-communicate/dao"
	"github.com/louis296/mesence-communicate/dao/model"
	"github.com/louis296/mesence-communicate/pkg/enum"
	"github.com/louis296/mesence-communicate/pkg/log"
	"github.com/louis296/mesence-communicate/pkg/util"
	"github.com/louis296/mesence-communicate/pkg/ws"
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
	conn.On("message", func(message []byte) {
		var msg ws.Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Error("Cannot unmarshal message from user [%v], message: %v", conn.UserPhone, string(message))
			return
		}
		switch msg.Type {
		case enum.WordPackageMessageType:
			onWord(conn, message)
		case enum.FriendRequestMessageType:
			onFriendRequest(conn, message)
		}
	})

	// handler close event
	conn.On("close", func(code int, text string) {
		log.Warn("[UserConnClose]--%d--%s", code, text)
		// send offline notice to friends
		friendPhones, err := getFriendPhones(conn)
		if err != nil {
			log.Error("Search user friends error, offline notify abort")
			return
		}
		offlineNotify(conn, friendPhones)
		ws.UserConnMap.Delete(conn.UserPhone)
	})
}

func onWord(conn *ws.UserConn, message []byte) {
	var msg WordMessage
	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Error("Cannot unmarshal data from user [%v], message: %v", conn.UserPhone, string(message))
	}
	data := msg.Data

	// check if receiver is valid
	_, err = dao.GetUserByPhone(data.To)
	if err != nil {
		log.Error("No user [%v] or db error, send word message abort", data.To)
		return
	}

	// store message
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
	recordMsg := model.Message{
		Content:  data.Content,
		From:     conn.UserPhone,
		To:       data.To,
		SendTime: util.TimeParse(data.SendTime),
	}
	err = dao.CreateMessage(recordMsg, tx)
	if err != nil {
		log.Error("Store message to db error, send word message continue without store")
	}

	// try to send message
	if item, ok := ws.UserConnMap.Load(data.To); ok {
		receiverConn := item.(*ws.UserConn)
		sendMessage := ws.Message{Type: enum.WordPackageMessageType}
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

func onFriendRequest(conn *ws.UserConn, message []byte) {
	var msg FriendRequestMessage
	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Error("Cannot unmarshal data from user [%v], message: %v", conn.UserPhone, string(message))
	}
	data := msg.Data

	// check if candidate is valid
	_, err = dao.GetUserByPhone(data.Candidate)
	if err != nil {
		log.Error("No user [%v] or db error, send word message abort", data.Candidate)
		return
	}

	// store friend request
	friendRequest := model.FriendRequest{
		Sender:        conn.UserPhone,
		Candidate:     data.Candidate,
		RequestStatus: 0,
		StartTime:     util.TimeParse(data.StartTime),
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
		log.Error("Store request to db error")
		return
	}

	// try to send friend request notice
	if item, ok := ws.UserConnMap.Load(data.Candidate); ok {
		candidateConn := item.(*ws.UserConn)
		notice := ws.Message{Type: enum.FriendRequestMessageType}
		notice.Data = FriendRequestData{
			Sender:        conn.UserPhone,
			Candidate:     data.Candidate,
			Content:       data.Content,
			StartTime:     data.StartTime,
			RequestStatus: enum.FriendRequestStatusMap[0],
		}
		err = candidateConn.Send(util.Marshal(notice))
		if err != nil {
			log.Error("Send word message to user [%v] error", candidateConn.UserPhone)
		}
	} else {
		log.Warn("Friend request receiver [%v] is offline, send abort", data.Candidate)
	}
}

func onlineNotify(conn *ws.UserConn, userPhones []string) {
	message := ws.Message{
		Type: enum.OnlineMessageType,
		Data: OnlineMessageData{Users: []string{conn.UserPhone}},
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

func offlineNotify(conn *ws.UserConn, userPhones []string) {
	message := ws.Message{
		Type: enum.OfflineMessageType,
		Data: OfflineMessageData{Users: []string{conn.UserPhone}},
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
