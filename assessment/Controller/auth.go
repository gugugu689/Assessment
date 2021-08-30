package Controller

import (
	"assessment/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ContextUserIDKey = "userID"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 假设Token放在Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" { //请求头缺少Auth Token
			ResponseErrorWithMsg(c, CodeInvalidToken, "请登陆")
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") { //Token格式不对
			ResponseErrorWithMsg(c, CodeInvalidToken, "请登录")
			c.Abort()
			return
		}
		// parts[1]获取到的tokenString
		mc, err := jwt.ParseToken(parts[1])
		//aToken错误
		if err != nil {
			c.Redirect(301, "127.0.0.1:8080/blog//refresh_token")
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set(ContextUserIDKey, mc.UserID)
		c.Next()
	}
}
