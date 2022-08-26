package model

import "time"

type FriendRequest struct {
	Base
	Sender        string    `gorm:"column:sender;NOT NULL"`
	Candidate     string    `gorm:"column:candidate;NOT NULL"`
	RequestStatus int       `gorm:"column:request_status;NOT NULL"`
	StartTime     time.Time `gorm:"column:start_time;NOT NULL"`
	Content       string    `gorm:"column:content"`
}

func (m *FriendRequest) TableName() string {
	return "friend_request"
}
