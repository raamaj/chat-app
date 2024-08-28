package entity

import "time"

type Article struct {
	ID           int64     `gorm:"column:id;primaryKey"`
	Title        string    `gorm:"column:title"`
	ImageUrl     string    `gorm:"column:image_url"`
	Content      string    `gorm:"column:content"`
	Author       string    `gorm:"column:author"`
	LikeCount    int       `gorm:"column:like_count"`
	CommentCount int       `gorm:"column:comment_count"`
	CreatedAt    time.Time `gorm:"column:created_at"`
}

func (u *Article) TableName() string {
	return "articles"
}
