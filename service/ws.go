package service

import (
	"encoding/json"
	"github.com/louis296/mesence-communicate/dao"
	"github.com/louis296/mesence-communicate/pkg/enum"
	"github.com/louis296/mesence-communicate/pkg/log"
	"github.com/louis296/mesence-communicate/pkg/util"
	"github.com/louis296/mesence-communicate/pkg/ws"
)

func UserConnHandler(conn *ws.UserConn) {
	log.Info("Conn by user [%v] open", conn.UserPhone)

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
			log.Error("Cannot marshal message from user [%v], message: %v", conn.UserPhone, message)
			return
		}
		switch msg.Type {
		case enum.WordPackageMessageType:
			onWord(conn, msg)
		case enum.FriendRequestMessageType:
			onFriendRequest(conn, msg)
		}
	})

	// handler close event
	conn.On("close", func(code int, text string) {
		log.Warn("[UserConnClose]--%d--%s", code, text)
		// send offline notice to friends
		friendPhones, err := getFriendPhones(conn)
		if err != nil {
			log.Error("Search user friends error, offline notify abort")
		}
		offlineNotify(conn, friendPhones)
		ws.UserConnMap.Delete(conn.UserPhone)
	})
}

func onWord(conn *ws.UserConn, message ws.Message) {

}

func onFriendRequest(conn *ws.UserConn, message ws.Message) {

}

func onlineNotify(conn *ws.UserConn, userPhones []string) {
	message := ws.Message{
		Type: enum.OnlineMessageType,
		Data: ws.OnlineMessageData{Users: []string{conn.UserPhone}},
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
		Data: ws.OnlineMessageData{Users: []string{conn.UserPhone}},
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
