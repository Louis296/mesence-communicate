package mongodb

import "time"

type Message struct {
	Content  string
	From     string
	To       string
	SendTime time.Time
	Seq      int64
	Uuid     string
}
