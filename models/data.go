package models

import (
	"time"
)

// コメント用モックデータの宣言
var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "first commnet",
		CreatedAt: time.Now(),
	}

	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 2,
		Message:   "second comment",
		CreatedAt: time.Now(),
	}
)

// 記事用モックデータの宣言
var (
	Article1 = Article{
		ID:          1,
		Title:       "first article",
		Contents:    "This is the test content1",
		Author:      "kodai",
		NiceNum:     1,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt:   time.Now(),
	}

	Article2 = Article{
		ID:          2,
		Title:       "sencond article",
		Contents:    "This is the test content2",
		Author:      "kodai",
		NiceNum:     2,
		CommentList: []Comment{Comment1},
		CreatedAt:   time.Now(),
	}
)
