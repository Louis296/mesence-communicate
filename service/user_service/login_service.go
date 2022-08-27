package user_service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/louis296/mesence-communicate/dao"
	"github.com/louis296/mesence-communicate/pkg/jwt"
)

type UserLoginReq struct {
	Phone    string
	Password string
}

type UserLoginResp struct {
	Token string
}

func (r *UserLoginReq) Handler(c *gin.Context) (interface{}, error) {
	user, err := dao.GetUserByPhone(r.Phone)
	if err != nil {
		return nil, err
	}
	if user.Password != r.Password {
		return nil, errors.New("Wrong password ")
	}
	token, err := jwt.GenerateToken(user.Phone, user.Name)
	if err != nil {
		return nil, err
	}
	return UserLoginResp{Token: token}, nil
}
