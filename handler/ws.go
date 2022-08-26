package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/louis296/mesence-communicate/pkg/log"
	"github.com/louis296/mesence-communicate/pkg/ws"
	"github.com/louis296/mesence-communicate/service/communicate_service"
	"net/http"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}}

func WebSocketHandler(c *gin.Context) {
	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Panic("%v", err)
	}
	userConn := ws.NewUserConn(socket)
	//c.Get(enum.CurrentUser)
	phone, _ := c.GetQuery("phone")
	userConn.UserPhone = phone
	communicate_service.UserConnHandler(userConn)
	userConn.StartReadMessage()
}
