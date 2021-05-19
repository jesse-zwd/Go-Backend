package model

import (
	"golang.org/x/crypto/bcrypt"
	"backend/global"
)

// User model
type User struct {
	global.GORM_MODEL
	UserName       string
	PasswordDigest string
	Nickname       string
	Status         string
	Avatar         string `gorm:"size:1000"`
	Email 		   string
	Bio            string `gorm:"size:2000"`
	Website        string
}

const (
	// PassWordCost 
	PassWordCost = 12
	// Active user
	Active string = "active"
	// Inactive user
	Inactive string = "inactive"
	// Suspended user
	Suspend string = "suspend"
)

// GetUser get user by id
func GetUser(ID interface{}) (User, error) {
	var user User
	result := global.GORM_DB.First(&user, ID)
	return user, result.Error
}

// SetPassword set password
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword check password
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
