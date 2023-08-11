package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/louis296/mesence-communicate/pkg/jwt"
	"github.com/louis296/mesence-communicate/pkg/log"
	"github.com/louis296/mesence-communicate/pkg/ws"
	"github.com/louis296/mesence-communicate/service/communicate_service"
	"net/http"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}}

func WebSocketHandler(c *gin.Context) {
	token := c.Query("Token")
	claims, err := jwt.ParseToken(token)
	if err != nil {
		c.JSON(200, gin.H{"Message": "do not have token or token invalid"})
		c.Abort()
		return
	}

	// check if user already online, and force offline the online user
	if v, ok := ws.UserConnMap.Load(claims.Phone); ok {
		conn := v.(*ws.UserConn)
		conn.Close()
		ws.UserConnMap.Delete(claims.Phone)
	}

	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Panic("%v", err)
	}
	userConn := ws.NewUserConn(socket)
	userConn.UserPhone = claims.Phone
	log.Info("Conn by user [%v] open", userConn.UserPhone)

	// online notify
	err = communicate_service.OnlineNotify(claims.Phone)
	if err != nil {
		log.Error(err.Error())
	}

	userConn.MsgHandler = communicate_service.HandleMessage
	userConn.OnClose = communicate_service.OfflineNotify
	ws.UserConnMap.Store(claims.Phone, userConn)

	userConn.StartReadMessage()
}
