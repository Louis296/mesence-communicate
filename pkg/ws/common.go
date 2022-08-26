package ws

type Message struct {
	Type string
	Data interface{}
}

type OnlineMessageData struct {
	Users []string
}
