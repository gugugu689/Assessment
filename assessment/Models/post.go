package Models

import "time"

type Post struct {
	ID            int64     `json:"id"`
	PostID        int64     `json:"post_id"`
	Title         string    `json:"title" binding:"required"`
	Content       string    `json:"content" binding:"required"`
	AuthorID      int64     `json:"author_id"`
	AuthorName    string    `json:"author_name"`
	CommunityID   int64     `json:"community_id" binding:"required"`
	CommunityName string    `json:"community_name"`
	View          int64     `json:"view"`
	CreateTime    time.Time `json:"create_time"`
}
