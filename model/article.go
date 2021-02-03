package model

import (
	"fmt"
)

// Article .
type Article struct {
	ID      int    `json:"Id"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
}

var articleList = []Article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// GetAllArticles returns a list of all the articles
func GetAllArticles() []Article {
	return articleList
}

// GetArticleByID returns the article by specified id
func GetArticleByID(id int) (*Article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}

	return nil, fmt.Errorf("Article %d not found", id)
}
