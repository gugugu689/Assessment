package Controller

import (
	"assessment/Dao/mysql"
	"assessment/Logic"
	"assessment/Models"
	"assessment/pkg/jwt"
	"errors"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

//SignUp 注册
func SignUp(c *gin.Context) {
	// 校验请求参数
	userform := new(Models.UserForm)
	if err := c.ShouldBindJSON(userform); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	if len(userform.UserName) == 0 || len(userform.Password) == 0 || len(userform.RePassword) == 0 {
		ResponseErrorWithMsg(c, CodeInvalidParam, "username & password can not be null")
		return
	}
	if userform.Password != userform.RePassword {
		ResponseErrorWithMsg(c, CodeInvalidParam, "password & repassword must be same")
		return
	}
	//逻辑处理
	err := Logic.SignUp(userform)
	if errors.Is(err, mysql.ErrorUserExit) {
		ResponseError(c, CodeUserExist)
		return
	}
	if err != nil {
		zap.L().Error("Logic.SignUp failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}

//Login 登陆
func Login(c *gin.Context) {
	// 校验请求参数
	user := new(Models.User)
	if err := c.ShouldBindJSON(user); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	if len(user.UserName) == 0 || len(user.Password) == 0 {
		ResponseErrorWithMsg(c, CodeInvalidParam, "username & password can not be null")
		return
	}
	// 逻辑处理
	err := Logic.Login(user)
	if errors.Is(err, mysql.ErrorPasswordWrong) {
		ResponseError(c, CodePasswordWrong)
		return
	}
	if errors.Is(err, mysql.ErrorUserExit) {
		ResponseError(c, CodeUserNotExist)
		return
	}
	if err != nil {
		zap.L().Error("Logic.Login failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	//    生成Token
	aToken, rToken, _ := jwt.GenToken(user.UserID)
	ResponseSuccess(c, gin.H{
		"accessToken":  aToken,
		"refreshToken": rToken,
		"userID":       user.UserID,
		"username":     user.UserName,
	})
}

//RefreshToken 刷新aToken rToken
func RefreshToken(c *gin.Context) {
	//假设rToken放在url
	rt := c.Query("refresh_token")

	authHeader := c.Request.Header.Get("Authorization")

	parts := strings.SplitN(authHeader, " ", 2)

	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	if err != nil {
		ResponseError(c, CodeNotLogin)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}
