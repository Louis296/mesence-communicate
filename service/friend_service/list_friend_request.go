package friend_service

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/mesence-communicate/dao"
	"github.com/louis296/mesence-communicate/dao/model"
	"github.com/louis296/mesence-communicate/pkg/util"
)

type ListFriendRequestReq struct {
	Limit  int
	Offset int
	Type   int
}

type ListFriendRequestResp struct {
	Total int
	List  []model.FriendRequestResp
}

func (r *ListFriendRequestReq) Handler(c *gin.Context) (interface{}, error) {
	user := util.MustGetCurrentUser(c)

	var err error
	var total int64

	var list []model.FriendRequest
	if r.Type == 0 {
		list, total, err = dao.ListFriendRequestsBySender(user.Phone, r.Limit, r.Offset)
		if err != nil {
			return nil, err
		}
	} else {
		list, total, err = dao.ListFriendRequestsByCandidate(user.Phone, r.Limit, r.Offset)
		if err != nil {
			return nil, err
		}
	}

	var respList []model.FriendRequestResp
	for _, item := range list {
		respList = append(respList, item.GenResp())
	}

	return ListFriendRequestResp{
		Total: int(total),
		List:  respList,
	}, nil
}
