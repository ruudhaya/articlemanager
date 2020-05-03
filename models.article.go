package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{
		ID:      1,
		Title:   "Article 1",
		Content: "Content of Article 1",
	},
	article{
		ID:      2,
		Title:   "Article 2",
		Content: "Content of Article 2",
	},
}

func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, article := range articleList {
		if article.ID == id {
			return &article, nil
		}
	}
	return nil, errors.New("Article Not Found")
}
