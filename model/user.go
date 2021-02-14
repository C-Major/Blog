package model

import "time"

// User .
type User struct {
	ID         uint64    `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Name       string    `json:"name"`
	Status     int8      `json:"status"`
	CreateTime time.Time `json:"create_time"`
}
