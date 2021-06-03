package model

import (
	"backend/global"
)

// User model
type User struct {
	global.GORM_MODEL
	UserName string
	Password string `json:"-"`
	Nickname string
	Status   string
	Avatar   string `gorm:"size:1000"`
	Email    string
}
