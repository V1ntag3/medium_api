package models

import "time"

type User struct {
	Id           string    `json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	Surname      string    `gorm:"not null" json:"surname"`
	About        string    `json:"about"`
	Email        string    `gorm:"unique" json:"email"`
	Password     []byte    `gorm:"not null" json:"-"`
	DateMember   time.Time `json:"timeMember"`
	ImageProfile string    `json:"imageProfile"`
	Followers    []*User   `gorm:"many2many:user_followers;joinForeignKey:follower_id;" json:"followers"`
	Following    []*User   `gorm:"many2many:user_followers;joinForeignKey:user_id;" json:"following"`
	Articles     []Article `gorm:"foreignKey:UserId;references:Id"`
}

type UserFollower struct {
	UserID     string `gorm:"primaryKey"`
	FollowerID string `gorm:"primaryKey"`
}
