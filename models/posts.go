package models

type Post struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ImageUrls string `json:"image"`
	CreatedAt string `json:"created_at"`
	Views     int    `json:"views"`
	UpdatedAt string `json:"updated_at"`
}
