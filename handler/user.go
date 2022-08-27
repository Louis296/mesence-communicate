package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/mesence-communicate/service/user_service"
)

func UserLogin(c *gin.Context) {
	handler(c, &user_service.UserLoginReq{})
}
