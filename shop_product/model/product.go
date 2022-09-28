package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string  `gorm:"type:string;not null"`
	Price float32 `gorm:"type:decimal(11,2)"`
	Num   int     `gorm:"type:int;DEFAULT:0"`
	Unit  string  `gorm:"type:string;not null"`
	Pic   string  `gorm:"type:string"`
	Desc  string  `gorm:"type:string;not null"`
}

func ProductList(currentPage int, pageSize int) (products []Product, count int64, err error) {
	err = Db.Model(&Product{}).Count(&count).Error
	err = Db.Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&products).Error
	return
}

func ProductAdd(data *Product) error {
	err := Db.Create(&data).Error
	return err
}

func ProductDel(id int) error {
	err := Db.Delete(&Product{}, id).Error
	return err
}

func ProductInfo(id int) (product *Product, err error) {
	err = Db.Where("ID=?", id).First(&product).Error
	return
}

func ProductEdit(id int, data *Product) error {
	err := Db.Model(&Product{}).Where("ID=?", id).Updates(&data).Error
	return err
}
