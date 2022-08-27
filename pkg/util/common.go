package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/louis296/mesence-communicate/dao/model"
	"github.com/louis296/mesence-communicate/pkg/enum"
	"github.com/louis296/mesence-communicate/pkg/log"
	"time"
)

func Marshal(data interface{}) string {
	if res, err := json.Marshal(data); err != nil {
		return ""
	} else {
		return string(res)
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
