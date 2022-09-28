package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:string;not null;unique;"`
	Password string `gorm:"type:string;not null;"`
	Status   int    `gorm:"type:int;DEFAULT:0"`
}

func UserLogin(email string, password string) bool {
	var count int64
	err := Db.Model(&User{}).Where("Email=? AND Password=?", email, password).Count(&count).Error
	return err == nil && count > 0
}

func UserRegistry(email string, password string) error {
	user := User{
		Email:    email,
		Password: password,
	}
	err := Db.Create(&user).Error
	return err
}

func UserExistEmail(email string) bool {
	var count int64
	err := Db.Select("Email=?", email).Find(&User{}).Count(&count).Error
	return err == nil && count == 0
}

func UserByName(email string) (user *User, err error) {
	err = Db.Where("Email=?", email).First(&user).Error
	return
}
