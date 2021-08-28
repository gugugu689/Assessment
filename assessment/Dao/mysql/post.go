package mysql

import (
	"assessment/Models"
	"database/sql"
)

// GetPostByID 根据post_id获取post详情
func GetPostByID(postid int64) (post *Models.Post, err error) {
	post = new(Models.Post)
	post.PostID = postid
	post.View++
	err = UpdateView(post)
	if err = db.Where("post_id=?", post.PostID).Find(post).Error; err == sql.ErrNoRows {
		err = ErrorInvalidID
		return
	}
	if err != nil {
		err = ErrorQueryFailed
		return
	}
	return
}

//CreatePost 创建帖子
func CreatePost(post1 *Models.Post) (err error) {
	post := new(Models.Post)
	user, err := GetUserByID(post1.AuthorID)
	if err != nil {
		return
	}
	com, err := GetCommunityByID(post1.CommunityID)
	if err != nil {
		return
	}
	post = post1
	post.AuthorName = user.UserName
	post.CommunityName = com.CommunityName
	db.Create(post)
	return
}

//DeletePost 删除帖子
func DeletePost(post_id int64) (err error) {
	err = db.Delete(Models.Post{}, "post_id=?", post_id).Error
	err = db.Delete(Models.Comment{}, "post_id=?", post_id).Error
	return
}

//UpdatePost 更新帖子
func UpdatePost(postform *Models.Post) (post *Models.Post, err error) {
	//获取（更新的）用户名
	user, err := GetUserByID(postform.AuthorID)
	if err != nil {
		return
	}
	//获取（更新的）社区名
	com, err := GetCommunityByID(postform.CommunityID)
	if err != nil {
		return
	}
	post1 := new(Models.Post)
	post1 = postform
	post1.AuthorName = user.UserName
	post1.CommunityName = com.CommunityName

	db.Save(post1)

	post = new(Models.Post)

	err = db.Where("post_id=?", postform.PostID).Find(post).Error
	return
}

//UpdateView 更新浏览量
func UpdateView(post *Models.Post) error {
	return db.Model(post).Updates(map[string]interface{}{
		"view": post.View,
	}).Error
}

//GetPostsBy 获取帖子列表
func GetPostsBy(p *Models.ParamCondition) (posts []*Models.Post, err error) {
	posts = make([]*Models.Post, 0, p.Size)
	sqlstr := p.Order + " " + "desc"
	if err = db.Limit(p.Size).Offset((p.Page - 1) * p.Size).Order(sqlstr).Find(posts).Error; err != nil {
		return
	}
	return
}
