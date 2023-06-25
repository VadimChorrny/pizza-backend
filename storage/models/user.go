package models

import "time"

type User struct {
	ID        string `gorm:"primary_key"`
	Email     string `gorm:"unique_index"`
	Password  string
	FirstName string
	LastName  string
	Role      string
	CreatedAt time.Time
}

func (u User) TableName() string {
	return "users"
}
