package models

import (
	"time"
)

// コメント
type Comment struct {
	CommentID int       `json:"comment_id"`
	ArticleID int       `json:"article_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

// 記事
type Article struct {
	ID          int       `json:"article_id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	Author      string    `json:"author"`
	NiceNum     int       `json:"nice_num"`
	CommentList []Comment `json:"comments"`
	CreatedAt   time.Time `json:"created_at"`
}
