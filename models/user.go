package models

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User TODO create another struct for InputUser and OutputUser
type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := DB.Where("username = @username", sql.Named("username", username)).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hash, err := MakePassword(u.Password)
		if err != nil {
			return nil
		}
		u.Password = hash
	}
	return
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hash, err := MakePassword(u.Password)
		if err != nil {
			return nil
		}
		u.Password = hash
	}
	return
}

func MakePassword(plainPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	return string(bytes), err
}
