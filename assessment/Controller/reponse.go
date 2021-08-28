package Controller

import "github.com/gin-gonic/gin"

//响应码

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodePasswordWrong
	CodeServerBusy

	CodeNotLogin
	CodeInvalidToken
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:       "success",
	CodeInvalidParam:  "请求参数错误",
	CodeUserExist:     "用户名已存在",
	CodeUserNotExist:  "用户名不存在",
	CodePasswordWrong: "用户名或密码错误",
	CodeServerBusy:    "服务繁忙",

	CodeNotLogin:     "需要登录",
	CodeInvalidToken: "无效的token",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}

type ResponseData struct {
	Code    ResCode     `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	})
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(200, &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(200, &ResponseData{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}
