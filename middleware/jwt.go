package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/mesence-communicate/dao"
	"github.com/louis296/mesence-communicate/pkg/enum"
	"github.com/louis296/mesence-communicate/pkg/jwt"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check token
		valid := false
		token := c.GetHeader("Authorization")
		if token != "" {
			claims, err := jwt.ParseToken(token)
			if err == nil {
				user, _ := dao.GetUserByPhone(claims.Phone)
				if user != nil {
					valid = true
					c.Set(enum.CurrentUser, user)
				}
			}
		}
		if valid {
			c.Next()
		} else {
			c.JSON(200, gin.H{"Message": "do not have token or token invalid"})
			c.Abort()
		}
	}
}
