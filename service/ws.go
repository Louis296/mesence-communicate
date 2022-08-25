package service

import (
	"encoding/json"
	"github.com/louis296/mesence-communicate/pkg/enum"
	"github.com/louis296/mesence-communicate/pkg/log"
	"github.com/louis296/mesence-communicate/pkg/ws"
)

func UserConnHandler(conn *ws.UserConn) {
	log.Info("Conn by user [%v] open", conn.UserPhone)

	// store user conn
	ws.UserConnMap[conn.UserPhone] = conn

	// notify all friends that user is online

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
	conn.On("close", func(code int, text string) {
		// on close func
	})
}

func onWord(conn *ws.UserConn, message ws.Message) {

}

func onFriendRequest(conn *ws.UserConn, message ws.Message) {

}

func onlineNotify() {

}
