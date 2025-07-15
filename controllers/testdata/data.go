package testdata

import "github.com/kodaiozekijp/go-blog-api-practice/models"

var articleTestData = []models.Article{
	models.Article{
		ID:       1,
		Title:    "1st Post",
		Contents: "This is my first blog",
		Author:   "kodai",
		NiceNum:  1,
	},
	models.Article{
		ID:       2,
		Title:    "2nd Post",
		Contents: "Second blog post",
		Author:   "kodai",
		NiceNum:  2,
	},
}

var commentTestData = []models.Comment{
	models.Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "1st comment",
	},
	models.Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "2nd comment",
	},
}
