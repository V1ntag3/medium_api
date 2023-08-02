package models

import "time"

// Articles table
type Article struct {
	Id          string    `json:"id"`
	Abstract    string    `gorm:"not null" json:"abstract"`
	Title       string    `gorm:"not null" json:"title"`
	Subtile     string    `gorm:"not null" json:"subtitle"`
	Text        string    `gorm:"not null" json:"text"`
	BannerImage string    `gorm:"not null" json:"photoBanner"`
	CreateTime  time.Time `gorm:"not null" json:"createTime"`
	UserId      string
	User        User `gorm:"foreignKey:UserId"`
}
