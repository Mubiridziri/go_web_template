package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&User{})

	adminUser := User{Name: "Admin", Username: "admin", Password: "admin"}
	err = db.Create(&adminUser).Error
	if err != nil {
		return
	}

	DB = db
}
