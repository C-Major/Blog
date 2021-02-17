package model

import (
	"time"

	"github.com/c-major/blog/constdef"
)

// BlogArticle .
type BlogArticle struct {
	ID         uint64    `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	AuthorID   uint64    `json:"author_id"`
	Status     int8      `json:"status"`
	CreateTime time.Time `gorm:"default:null" json:"create_time"`
	UpdateTime time.Time `gorm:"default:null" json:"update_time"`
}

// TableName .
func (BlogArticle) TableName() string {
	return constdef.BlogArticleTable
}
