package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/mesence-communicate/service/message_service"
)

func (h Handler) ListMessageRecord20220101(c *gin.Context) {
	handler(c, &message_service.ListMessageRecordReq{})
}
