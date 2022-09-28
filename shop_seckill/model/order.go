package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Email string `gorm:"type:string"`
	Sid   int    `gorm:"type:int"`
}

type Payment struct {
	gorm.Model
	Email string  `gorm:"type:string;not null"`
	Pid   int     `gorm:"type:int;not null"`
	Price float32 `gorm:"type:decimal(11,2)"`
}

func OrderCheck(email string, id int) bool {
	var count int64
	err := Db.Model(&Order{}).Select("Email=? AND Sid=?", email, id).Count(&count).Error
	return err == nil && count == 0
}
