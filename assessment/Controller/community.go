package Controller

import (
	"assessment/Dao/mysql"
	"assessment/Logic"
	"assessment/Models"

	"github.com/gin-gonic/gin"
)

// CreateCommunity 创建社区
func CreateCommunity(c *gin.Context) {
	//获取参数
	com := new(Models.Community)
	if err := c.ShouldBindJSON(com); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//逻辑处理
	if err := Logic.CreateCommunity(com); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}

//GetCommunities 查看所有社区
func GetCommunities(c *gin.Context) {
	coms := make([]*Models.Community, 0, 5)
	var err error
	coms, err = mysql.GetCommunities()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, coms)
}
