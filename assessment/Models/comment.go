package Models

type Comment struct {
	ID       int64  `json:"id"`
	PostID   int64  `json:"post_id"`
	AuthorID int64  `json:"author_id"`
	Content  string `json:"content" binding:"required"`
}
