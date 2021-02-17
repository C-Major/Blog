package model

import (
	"time"

	"github.com/c-major/blog/constdef"
)

// BlogUser .
type BlogUser struct {
	ID         uint64    `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Name       string    `json:"name"`
	Status     int8      `json:"status"`
	CreateTime time.Time `gorm:"default:null" json:"create_time"`
	UpdateTime time.Time `gorm:"default:null" json:"update_time"`
}

// TableName .
func (BlogUser) TableName() string {
	return constdef.BlogUserTable
}
