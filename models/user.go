package models

import "time"

// User
type User struct {
	Id           string    `json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	Surname      string    `gorm:"not null" json:"surname"`
	About        string    `json:"about"`
	Email        string    `gorm:"unique" json:"email"`
	Password     []byte    `gorm:"not null" json:"-"`
	DateMember   time.Time `json:"timeMember"`
	ImageProfile string    `json:"imageProfile"`
	Followers    []*User   `gorm:"many2many:user_followers;joinForeignKey:follower_id;joinReferences:user_id" json:"followers"`
	Followings   []*User   `gorm:"many2many:user_followers;joinForeignKey:user_id;joinReferences:follower_id" json:"following"`
	Articles     []Article `gorm:"foreignKey:UserId;references:Id"`
}

// Followers table
type UserFollower struct {
	UserID     string `gorm:"primaryKey"`
	FollowerID string `gorm:"primaryKey"`
}

// Save valids tokens
type ValidToken struct {
	Token string `db:"token"`
}
