package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Email	  string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	IsMember  bool      `json:"isMember"`
}

type UserResponse struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Email	  string    `json:"email" gorm:"unique"`
}