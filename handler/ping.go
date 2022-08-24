package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/mesence-communicate/service"
)

func (h Handler) Ping20220101(c *gin.Context) {
	handler(c, &service.PingReq{})
}
