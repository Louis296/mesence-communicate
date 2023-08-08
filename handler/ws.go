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
	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Panic("%v", err)
	}
	userConn := ws.NewUserConn(socket)
	//user := util.MustGetCurrentUser(c)
	userConn.UserPhone = claims.Phone
	communicate_service.UserConnHandler(userConn)
	userConn.StartReadMessage()
}
