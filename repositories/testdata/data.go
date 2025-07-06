package testdata

import "github.com/kodaiozekijp/go-blog-api-practice/models"

var SelectArticleTestData = []models.Article{
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

var InsertArticleTestData = models.Article{
	Title:    "insertTest",
	Contents: "insert test",
	Author:   "kodai",
	NiceNum:  0,
}
