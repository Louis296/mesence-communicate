package util

import "encoding/json"

func Marshal(data interface{}) string {
	if res, err := json.Marshal(data); err != nil {
		return ""
	} else {
		return string(res)
	}
}
