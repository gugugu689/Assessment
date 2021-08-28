package Controller

import (
	"assessment/Dao/mysql"
	"assessment/Models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//CreateComment 发表评论
func CreateComment(c *gin.Context) {
	//获取参数url里的id
	postID := c.Param("id")
	post_id, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		zap.L().Error("parseint failed", zap.Error(err))
		return
	}
	comment := new(Models.Comment)
	if err := c.ShouldBindJSON(comment); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	comment.PostID = post_id
	comment.AuthorID, _ = GetCtxUserID(c)
	if err := mysql.CreateComment(comment); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

//GetComment 获取评论
func GetComment(c *gin.Context) {
	//获取参数url里的id
	postID := c.Param("id")
	post_id, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		zap.L().Error("parseint failed", zap.Error(err))
		return
	}
	comments := make([]*Models.Comment, 0, 5)
	comments, err = mysql.GetComment(post_id)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, comments)
}
