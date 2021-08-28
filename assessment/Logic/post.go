package Logic

import (
	"assessment/Dao/mysql"
	"assessment/Models"
	"assessment/pkg/snowflake"

	"go.uber.org/zap"
)

//GetPostDetail 查看帖子详情
func GetPostDetail(postID int64) (post *Models.Post, err error) {
	post = new(Models.Post)
	post, err = mysql.GetPostByID(postID)
	if err == mysql.ErrorInvalidID {
		return
	}
	if err != nil {
		zap.L().Error("mysql.GetPostByID failed", zap.Error(err))
		return nil, err
	}
	return
}

//CreatePost 创建帖子
func CreatePost(post *Models.Post) (err error) {
	post.PostID = snowflake.GenID()
	if err = mysql.CreatePost(post); err != nil {
		zap.L().Error("mysql.CreatePost failed", zap.Error(err))
		return
	}
	return
}

//DeletePost 删除帖子
func DeletePost(postid int64) (err error) {
	err = mysql.DeletePost(postid)
	return err
}

//UpdatePost 更新帖子
func UpdatePost(postform *Models.Post) (post *Models.Post, err error) {
	post = new(Models.Post)
	post, err = mysql.UpdatePost(postform)
	if err != nil {
		zap.L().Error("mysql.updatepost failed", zap.Error(err))
		return
	}
	return
}

//GetPostsBy 获取帖子列表
func GetPostsBy(p *Models.ParamCondition) (posts []*Models.Post, err error) {
	posts = make([]*Models.Post, 0, p.Size)
	posts, err = mysql.GetPostsBy(p)
	if err != nil {
		zap.L().Error("mysql.GetPosts failed", zap.Error(err))
		return
	}
	return
}
