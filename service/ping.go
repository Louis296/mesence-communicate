package service

import "github.com/gin-gonic/gin"

type PingReq struct {
}

func (r *PingReq) Handler(c *gin.Context) (interface{}, error) {
	return "pong", nil
}
