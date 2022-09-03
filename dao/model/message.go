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

type MessageResp struct {
	Id       int
	Content  string
	From     string
	To       string
	SendTime time.Time
}

func (m *Message) GenResp() MessageResp {
	return MessageResp{
		Id:       m.Id,
		Content:  m.Content,
		From:     m.From,
		To:       m.To,
		SendTime: m.SendTime,
	}
}
