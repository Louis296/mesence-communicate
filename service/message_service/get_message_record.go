package message_service

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/mesence-communicate/pkg/mongodb"
	"github.com/louis296/mesence-communicate/pkg/util"
)

type ListMessageRecordReq struct {
	Limit       int64
	Offset      int64
	AnotherUser string
	StartTime   string
	EndTime     string
}

type ListMessageRecordResp struct {
	List []mongodb.Message
}

func (r *ListMessageRecordReq) Handler(c *gin.Context) (interface{}, error) {
	user := util.MustGetCurrentUser(c)
	//records, err := dao.ListMessageByTwoUserPhone(user.Phone, r.AnotherUser, r.Offset, r.Limit, r.StartTime, r.EndTime)
	records, err := mongodb.ListMessage(user.Phone, r.AnotherUser, r.Offset, r.Limit, util.TimeParse(r.StartTime), util.TimeParse(r.EndTime))
	if err != nil {
		return nil, err
	}
	return ListMessageRecordResp{List: records}, nil
}
