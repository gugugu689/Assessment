package Controller

import (
	"github.com/gin-gonic/gin"

	"errors"
)

var ErrorUserNotLogin = errors.New("请登录")

//GetCtxUserID 获取当前登陆用户id
func GetCtxUserID(c *gin.Context) (userID int64, err error) {
	id, ok := c.Get("userID")
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID = id.(int64)
	return
}
