package friend_service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/louis296/mesence-communicate/dao"
	"github.com/louis296/mesence-communicate/dao/model"
	"github.com/louis296/mesence-communicate/pkg/log"
	"github.com/louis296/mesence-communicate/pkg/pb"
	"github.com/louis296/mesence-communicate/pkg/util"
	"github.com/louis296/mesence-communicate/pkg/ws"
)

type FinishFriendRequestReq struct {
	Id   int
	Type int
}

func (r *FinishFriendRequestReq) Handler(c *gin.Context) (interface{}, error) {
	user := util.MustGetCurrentUser(c)
	friendRequest, err := dao.GetFriendRequestById(r.Id)
	if err != nil {
		return nil, err
	}
	if user.Phone != friendRequest.Candidate {
		return nil, errors.New("Friend request must finished by candidate ")
	}
	if friendRequest.RequestStatus != 2 {
		return nil, errors.New("Friend request already finish ")
	}
	if r.Type == 1 {
		friendRequest.RequestStatus = int(pb.RequestStatus_Accepted)
	} else {
		friendRequest.RequestStatus = int(pb.RequestStatus_Refused)
	}

	tx := dao.DB.Begin()
	defer func() {
		if tx.Error != nil {
			tx.Rollback()
		} else if recover() != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = dao.SaveFriendRequest(tx, *friendRequest)
	if err != nil {
		return nil, err
	}

	// accept friend request
	if r.Type == 1 {
		another, err := dao.GetUserByPhone(friendRequest.Sender)
		if err != nil {
			tx.Error = err
			return nil, err
		}
		relationA := model.FriendRelation{
			UserID:         user.Id,
			UserPhone:      user.Phone,
			FriendID:       another.Id,
			FriendPhone:    another.Phone,
			FriendNoteName: "",
		}
		err = dao.CreateFriendRelation(tx, relationA)
		if err != nil {
			return nil, err
		}
		relationB := model.FriendRelation{
			UserID:         another.Id,
			UserPhone:      another.Phone,
			FriendID:       user.Id,
			FriendPhone:    user.Phone,
			FriendNoteName: "",
		}
		err = dao.CreateFriendRelation(tx, relationB)
		if err != nil {
			return nil, err
		}
	}

	// try to notice sender
	if item, ok := ws.UserConnMap.Load(friendRequest.Sender); ok {
		senderConn := item.(*ws.UserConn)
		msg := &pb.Msg{
			Type: pb.Type_FriendRequest,
			Data: &pb.Data{
				From:          friendRequest.Sender,
				Content:       friendRequest.Content,
				SendTime:      util.TimeFormat(friendRequest.StartTime),
				Candidate:     friendRequest.Candidate,
				RequestStatus: pb.RequestStatus(friendRequest.RequestStatus),
			},
		}
		err = senderConn.Send(util.Marshal(msg))
		if err != nil {
			log.Error("Send word message to user [%v] error", friendRequest.Sender)
		}
	} else {
		log.Warn("Friend request message receiver [%v] is offline, send abort", friendRequest.Sender)
	}

	return nil, nil
}
