package models

import "time"

type Articule struct {
	Id          uint      `json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Subtile     string    `gorm:"not null" json:"subtitle"`
	Text        string    `gorm:"not null" json:"text"`
	PhotoBanner string    `gorm:"not null" json:"photoBanner"`
	CreateTime  time.Time `gorm:"not null" json:"createTime"`
	UserId      uint
}
