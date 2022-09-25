package user_service

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/mesence-communicate/dao/model"
	"github.com/louis296/mesence-communicate/pkg/util"
)

type GetUserInfoReq struct {
}

type GetUserInfoResp struct {
	Info model.UserResp
}

func (r *GetUserInfoReq) Handler(c *gin.Context) (interface{}, error) {
	currentUser := util.MustGetCurrentUser(c)
	return GetUserInfoResp{Info: currentUser.GenResp()}, nil
}
