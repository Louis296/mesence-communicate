package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/mesence-communicate/service/friend_service"
)

func (h Handler) ListFriendRequest20220101(c *gin.Context) {
	handler(c, &friend_service.ListFriendRequestReq{})
}

func (h Handler) FinishFriendRequest20220101(c *gin.Context) {
	handler(c, &friend_service.FinishFriendRequestReq{})
}
