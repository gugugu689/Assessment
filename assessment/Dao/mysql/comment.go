package mysql

import "assessment/Models"

//CreateComment 发表评论
func CreateComment(comment *Models.Comment) (err error) {
	err = db.Create(comment).Error
	return
}

//GetComment 获取评论
func GetComment(postid int64) (comments []*Models.Comment, err error) {
	comments = make([]*Models.Comment, 0, 5)
	if err = db.Where("post_id=?", postid).Find(comments).Error; err != nil {
		return nil, err
	}
	return
}
