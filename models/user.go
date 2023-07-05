package models

import "time"

type User struct {
	Id         uint
	Name       string `gorm:"not null"`
	Surname    string `gorm:"not null"`
	About      string
	Email      string `gorm:"unique"`
	Password   []byte `gorm:"not null"`
	DateMember time.Time
}
