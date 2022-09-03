package message_service

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/mesence-communicate/dao"
	"github.com/louis296/mesence-communicate/dao/model"
	"github.com/louis296/mesence-communicate/pkg/util"
)

type ListMessageRecordReq struct {
	Limit       int
	Offset      int
	AnotherUser string
	StartTime   string
	EndTime     string
}

type ListMessageRecordResp struct {
	List []model.MessageResp
}

func (r *ListMessageRecordReq) Handler(c *gin.Context) (interface{}, error) {
	user := util.MustGetCurrentUser(c)
	records, err := dao.ListMessageByTwoUserPhone(user.Phone, r.AnotherUser, r.Offset, r.Limit, r.StartTime, r.EndTime)
	if err != nil {
		return nil, err
	}
	var respList []model.MessageResp
	for _, record := range records {
		respList = append(respList, record.GenResp())
	}
	return ListMessageRecordResp{List: respList}, nil
}
