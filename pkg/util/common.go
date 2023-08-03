package util

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/louis296/mesence-communicate/dao/model"
	"github.com/louis296/mesence-communicate/pkg/enum"
	"github.com/louis296/mesence-communicate/pkg/log"
	"github.com/louis296/mesence-communicate/pkg/pb"
	"time"
)

func Marshal(data *pb.Msg) []byte {
	if res, err := proto.Marshal(data); err != nil {
		return nil
	} else {
		return res
	}
}

func TimeParse(s string) time.Time {
	ans, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Error("Time parse error [%v]", err.Error())
	}
	return ans
}

func TimeFormat(t time.Time) string {
	return t.Format(time.RFC3339)
}

func MustGetCurrentUser(c *gin.Context) *model.User {
	v, _ := c.Get(enum.CurrentUser)
	return v.(*model.User)
}
