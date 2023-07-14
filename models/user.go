package models

import "time"

type User struct {
	Id           uint      `json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	Surname      string    `gorm:"not null" json:"surname"`
	About        string    `json:"about"`
	Email        string    `gorm:"unique" json:"email"`
	Password     []byte    `gorm:"not null" json:"-"`
	DateMember   time.Time `json:"timeMember"`
	ImageProfile string    `json:"imageProfile"`
	// Follows      []User     `json:"follows"`
	// Following    []User     `json:"followings"`
	Articles []Article `gorm:"foreignKey:UserId;references:Id"`
}
