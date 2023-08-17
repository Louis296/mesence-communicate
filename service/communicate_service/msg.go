package communicate_service

import (
	"errors"
	"github.com/IBM/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/louis296/mesence-communicate/dao"
	"github.com/louis296/mesence-communicate/dao/model"
	"github.com/louis296/mesence-communicate/pkg/log"
	"github.com/louis296/mesence-communicate/pkg/mongodb"
	"github.com/louis296/mesence-communicate/pkg/pb"
	"github.com/louis296/mesence-communicate/pkg/util"
	"gorm.io/gorm"
)

var MsgFromMQConsumer MsgConsumer

type MsgConsumer struct {
}

func (m MsgConsumer) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (m MsgConsumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (m MsgConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if len(msg.Value) != 0 {
			handleMsgFromMQ(msg)
		}
	}
	return nil
}

func handleMsgFromMQ(msgFromMQ *sarama.ConsumerMessage) {
	msg := &pb.Msg{}
	err := proto.Unmarshal(msgFromMQ.Value, msg)
	if err != nil {
		log.Error(err.Error())
		return
	}
	switch msg.Type {
	case pb.Type_Word:
		handleWordMsgFromMQ(msg)
	case pb.Type_FriendRequest:
		handleFriendRequestMsgFromMQ(msg)

	}
}

func handleWordMsgFromMQ(msg *pb.Msg) {
	data := msg.Data

	// check if receiver is valid
	_, err := dao.GetUserByPhone(data.To)
	if err != nil {
		log.Error("No user [%v] or mongodb error, send word message abort", data.To)
		return
	}

	// store message
	err = mongodb.SaveMessage(msg)
	if err != nil {
		log.Error("Save message error")
		return
	}

	// push message
	PushMessage(msg)
}

func handleFriendRequestMsgFromMQ(msg *pb.Msg) {
	data := msg.Data

	// check if candidate is valid
	_, err := dao.GetUserByPhone(data.To)
	if err != nil {
		log.Error("No user [%v] or db error, send word message abort", data.Candidate)
		return
	}

	// check if friend relation already exist
	_, err = dao.GetFriendRelationByUserAndFriend(data.From, data.To)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Friend relation already exist")
		return
	}

	// check if friend request already exist and not finish
	_, err = dao.GetFriendRequestBySenderAndCandidateAndStatus(data.From, data.To, 2)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Friend request already exist")
		return
	}

	// store friend request
	friendRequest := model.FriendRequest{
		Sender:        data.From,
		Candidate:     data.To,
		RequestStatus: int(pb.RequestStatus_Waiting),
		StartTime:     util.TimeParse(data.SendTime),
		Content:       data.Content,
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
	if err = dao.CreateFriendRequest(tx, friendRequest); err != nil {
		log.Error("Store request to mongodb error")
		return
	}

	msg.Data.RequestStatus = pb.RequestStatus_Waiting
	PushMessage(msg)

	return
}
