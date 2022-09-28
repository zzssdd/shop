package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name string `gorm:"type:string;not null"`
	Pic  string `gorm:"type:string"`
}

type Payment struct {
	gorm.Model
	Email    string  `gorm:"type:string;not null"`
	Pid      int     `gorm:"type:int;not null"`
	BuyPrice float32 `gorm:"type:decimal(11,2)"`
	Pro      Product `gorm:"ForeignKey:Pid"`
}

func PaymentList(email string, pageSize int, currentPage int) (payments []Payment, count int64, err error) {
	err = Db.Model(Payment{}).Count(&count).Error
	err = Db.Where("Email=?", email).Offset((currentPage - 1) * pageSize).Limit(pageSize).Preload("Pro").Find(&payments).Error
	return
}

func PaymentAdd(payment *Payment) error {
	err := Db.Create(&payment).Error
	return err
}

func PaymentDel(id int) error {
	err := Db.Delete(&Payment{}, id).Error
	return err
}
