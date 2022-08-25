package service

import (
	"github.com/louis296/mesence-communicate/pkg/log"
	"github.com/louis296/mesence-communicate/pkg/ws"
)

func UserConnHandler(conn *ws.UserConn) {
	log.Info("Conn by user [%v] open", conn.UserPhone)
	conn.On("message", func(message []byte) {
		conn.Send("hello" + string(message))
	})
}
