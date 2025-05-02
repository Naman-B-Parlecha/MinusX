package models

import "time"

type Comment struct {
	ID        string    `json:"id"`
	PostID    string    `json:"postid"`
	UserID    string    `json:"userid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `json:"content"`
}
