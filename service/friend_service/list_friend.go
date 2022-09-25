package friend_service

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/mesence-communicate/dao"
	"github.com/louis296/mesence-communicate/dao/model"
	"github.com/louis296/mesence-communicate/pkg/util"
)

type ListFriendReq struct {
}

type ListFriendResp struct {
	List []model.FriendRelationResp
}

func (r *ListFriendReq) Handler(c *gin.Context) (interface{}, error) {
	user := util.MustGetCurrentUser(c)
	relations, err := dao.GetFriendRelationsByUserPhone(user.Phone)
	if err != nil {
		return nil, err
	}
	var friendIds []int
	for _, relation := range relations {
		friendIds = append(friendIds, relation.FriendID)
	}
	friendLookup, err := dao.GetUserIdLookupByIds(friendIds)
	if err != nil {
		return nil, err
	}
	var respList []model.FriendRelationResp
	for _, relation := range relations {
		resp := relation.GenResp()
		resp.Friend = friendLookup[relation.FriendID].GenResp()
		respList = append(respList, resp)
	}
	return ListFriendResp{List: respList}, nil
}
