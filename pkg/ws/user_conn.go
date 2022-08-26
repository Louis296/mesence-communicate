package ws

import (
	"github.com/chuckpreslar/emission"
	"github.com/gorilla/websocket"
	"github.com/louis296/mesence-communicate/pkg/enum"
	"github.com/louis296/mesence-communicate/pkg/log"
	"github.com/louis296/mesence-communicate/pkg/util"
	"net"
	"sync"
	"time"
)

// heart package send time
const pingPeriod = 10 * time.Second

// UserConnMap is a sync.Map to store use conn
var UserConnMap sync.Map

type UserConn struct {
	emission.Emitter
	socket    *websocket.Conn
	mutex     *sync.Mutex
	closed    bool
	UserPhone string
}

func NewUserConn(socket *websocket.Conn) *UserConn {
	conn := UserConn{
		Emitter: *emission.NewEmitter(),
		socket:  socket,
		mutex:   new(sync.Mutex),
		closed:  false,
	}
	// closed by client
	conn.socket.SetCloseHandler(func(code int, text string) error {
		conn.Emit("close", code, text)
		conn.closed = true
		return nil
	})
	return &conn
}

func (conn *UserConn) StartReadMessage() {
	in := make(chan []byte)
	pingTicker := time.NewTicker(pingPeriod)

	c := conn.socket
	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Warn("Get error: %v, user conn may closed", err.Error())
				if e, ok := err.(*websocket.CloseError); ok {
					conn.Emit("close", e.Code, e.Text)
				} else if e, ok := err.(*net.OpError); ok {
					conn.Emit("close", 1008, e.Error())
				}
				break
			}
			in <- message
		}
	}()

	for {
		select {
		case <-pingTicker.C:
			log.Info("Send heart package")
			heartPackage := &Message{Type: enum.HeartPackageMessageType}
			if err := conn.Send(util.Marshal(heartPackage)); err != nil {
				log.Error("Send heart package error on user [%v] conn", conn.UserPhone)
				pingTicker.Stop()
				conn.Emit("close", 1000, "cannot send heart package")
				//conn.Close()
				return
			}
		case message := <-in:
			log.Info("Receive data: %v from user [%v]", message, conn.UserPhone)
			conn.Emit("message", message)
		}
	}
}

func (conn *UserConn) Send(message string) error {
	log.Info("Send data: %s", message)
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	return conn.socket.WriteMessage(websocket.TextMessage, []byte(message))
}

func (conn *UserConn) Close() {
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	if !conn.closed {
		log.Info("Close ws connection: %v", conn)
		conn.socket.Close()
		conn.closed = true
	} else {
		log.Info("Connection already closed: %v", conn)
	}
}
