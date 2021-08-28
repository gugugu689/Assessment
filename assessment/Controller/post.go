package Controller

import (
	"assessment/Dao/mysql"
	"assessment/Logic"
	"assessment/Models"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

//GetPostDetail 查看帖子详情
func GetPostDetail(c *gin.Context) {
	//获取参数url里的id
	postID := c.Param("id")
	post_id, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		zap.L().Error("parseint failed", zap.Error(err))
		return
	}
	//逻辑处理
	post, err := Logic.GetPostDetail(post_id)
	if err == mysql.ErrorInvalidID {
		ResponseErrorWithMsg(c, CodeServerBusy, "post_id错误")
		return
	}
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, "查找错误")
		zap.L().Error("logic.GetPostDetail failed", zap.Error(err))
		return
	}
	//返回响应
	ResponseSuccess(c, post)
}

//CreatePost 创建帖子
func CreatePost(c *gin.Context) {
	//获取参数
	post := new(Models.Post)
	if err := c.ShouldBindJSON(post); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 获取用户id
	userID, err := GetCtxUserID(c)
	if err != nil {
		ResponseError(c, CodeNotLogin)
		return
	}
	post.AuthorID = userID
	//业务处理
	if err := Logic.CreatePost(post); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

//UpdatePost 更新帖子
func UpdatePost(c *gin.Context) {
	//获取参数url里的id
	postID := c.Param("id")
	post_id, _ := strconv.ParseInt(postID, 10, 64)
	postform := new(Models.Post)
	postform.PostID = post_id
	if err := c.ShouldBindJSON(postform); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//逻辑处理
	userID, err := GetCtxUserID(c)
	if err != nil {
		ResponseError(c, CodeNotLogin)
		return
	}
	postform.AuthorID = userID
	post, err := Logic.UpdatePost(postform)
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, "Logic.UpdatePost failed")
		zap.L().Error("Update post failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, post)

}

//DeletePost 删除帖子
func DeletePost(c *gin.Context) {
	//获取参数url里的id
	postID := c.Param("id")
	post_id, _ := strconv.ParseInt(postID, 10, 64)
	//逻辑处理
	err := Logic.DeletePost(post_id)
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, "Logic.DeletePost failed")
		zap.L().Error("delete post failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, nil)
}

//GetPostBy 根据 浏览数/时间 获取帖子列表
func GetPostsBy(c *gin.Context) {
	// 从url里获取 query参数 /posts_by?page=1&size=10&order=create_time
	p := &Models.ParamCondition{
		Page:  1,
		Size:  10,
		Order: "create_time",
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("posts ShouldBindQuery failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	posts, err := Logic.GetPostsBy(p)
	if err != nil {
		zap.L().Error("logic.GetPostListBy failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, posts)
}
