package model

import "time"

type ArticleResponse struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	ImageUrl     string    `json:"image_url"`
	Content      string    `json:"content"`
	Author       string    `json:"author"`
	LikeCount    int       `json:"like_count"`
	CommentCount int       `json:"comment_count"`
	CreatedAt    time.Time `json:"posted"`
}
