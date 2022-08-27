package model

import (
	"github.com/louis296/mesence-communicate/pkg/enum"
	"time"
)

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

type FriendRequestResp struct {
	Id            int
	Sender        string
	Candidate     string
	RequestStatus string
	StartTime     time.Time
	Content       string
}

func (m *FriendRequest) GenResp() FriendRequestResp {
	return FriendRequestResp{
		Id:            m.Id,
		Sender:        m.Sender,
		Candidate:     m.Candidate,
		RequestStatus: enum.FriendRequestStatusMap[m.RequestStatus],
		StartTime:     m.StartTime,
		Content:       m.Content,
	}
}
