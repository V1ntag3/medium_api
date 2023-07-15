package models

import "time"

type Article struct {
	Id          uint      `json:"id"`
	Abstract    string    `gorm:"not null" json:"abstract"`
	Title       string    `gorm:"not null" json:"title"`
	Subtile     string    `gorm:"not null" json:"subtitle"`
	Text        string    `gorm:"not null" json:"text"`
	BannerImage string    `gorm:"not null" json:"photoBanner"`
	CreateTime  time.Time `gorm:"not null" json:"createTime"`
	UserId      uint
	User        User `gorm:"foreignKey:UserId"`
}
