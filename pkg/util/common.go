package util

import (
	"encoding/json"
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
