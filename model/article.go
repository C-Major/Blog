package model

import (
	"fmt"
	"time"
)

// Article .
type Article struct {
	ID         uint64    `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	AuthorID   uint64    `json:"author_id"`
	Status     int8      `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
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
func GetArticleByID(id uint64) (*Article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}

	return nil, fmt.Errorf("Article %d not found", id)
}
