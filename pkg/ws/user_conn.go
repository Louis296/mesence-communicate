package ws

import (
	"github.com/chuckpreslar/emission"
	"github.com/gorilla/websocket"
	"github.com/louis296/mesence-communicate/pkg/log"
	"github.com/louis296/mesence-communicate/pkg/pb"
	"github.com/louis296/mesence-communicate/pkg/util"
	"net"
	"sync"
	"time"
)

// heart package send time
const pingPeriod = 10 * time.Second

// UserConnMap is a sync.Map to store use conn
var UserConnMap sync.Map
var UserConnPool = sync.Pool{New: func() interface{} {
	return new(UserConn)
}}

type UserConn struct {
	emission.Emitter
	socket    *websocket.Conn
	mutex     *sync.Mutex
	closed    bool
	UserPhone string
}

func NewUserConn(socket *websocket.Conn) *UserConn {
	conn := UserConnPool.Get().(*UserConn)
	conn.ResetUserConn(
		*emission.NewEmitter(),
		socket,
		new(sync.Mutex),
		false,
		"")
	// closed by client
	conn.socket.SetCloseHandler(func(code int, text string) error {
		conn.mutex.Lock()
		conn.closed = true
		conn.mutex.Unlock()

		log.Warn("[UserConnClose]--%d--%s", code, text)
		conn.Emit("close", code, text)
		UserConnMap.Delete(conn.UserPhone)
		return nil
	})
	return conn
}

func (conn *UserConn) ResetUserConn(
	emitter emission.Emitter,
	socket *websocket.Conn,
	mutex *sync.Mutex,
	closed bool,
	userPhone string) {
	conn.Emitter = emitter
	conn.socket = socket
	conn.mutex = mutex
	conn.closed = closed
	conn.UserPhone = userPhone
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
				if conn.closed {
					break
				}
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
			heartPackage := &pb.Msg{Type: pb.Type_HeartPackage}
			if err := conn.Send(util.Marshal(heartPackage)); err != nil {
				log.Error("Send heart package error on user [%v] conn", conn.UserPhone)
				pingTicker.Stop()
				return
			}
		case message := <-in:
			log.Info("Receive data: %v from user [%v]", string(message), conn.UserPhone)
			conn.Emit("message", message)
		}
	}
}

func (conn *UserConn) Send(bs []byte) error {
	log.Info("Send data: %v", bs)
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	return conn.socket.WriteMessage(websocket.BinaryMessage, bs)
}

// Close 主动关闭用户连接
func (conn *UserConn) Close() {
	defer UserConnPool.Put(conn)
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
