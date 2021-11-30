package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Names    string `json:"names"`
	Username string `gorm:"unique_index;not null" json:"username"`
	Email    string `gorm:"unique_index;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

type Token struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	Username  string `gorm:"unique_index;not null" json:"username"`
	Token     string `json:"token"`
}
