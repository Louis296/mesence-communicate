package model

import "time"

type Message struct {
	Base
	Content  string    `gorm:"column:content"`
	From     string    `gorm:"column:from"`
	To       string    `gorm:"column:to"`
	SendTime time.Time `gorm:"column:send_time"`
}

func (m *Message) TableName() string {
	return "message"
}
