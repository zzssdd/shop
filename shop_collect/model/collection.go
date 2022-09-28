package model

import (
	"fmt"
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
type Collection struct {
	gorm.Model
	Email string  `gorm:"type:string;not null"`
	Pid   int     `gorm:"type:int;not null"`
	Pro   Product `gorm:"ForeignKey:Pid"`
}

func CollectList(email string, currentPage, pageSize int) (collections []*Collection, count int64, err error) {
	err = Db.Model(&Collection{}).Where("Email=?", email).Count(&count).Error
	err = Db.Where("Email=?", email).Preload("Pro").Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&collections).Error
	return
}

func CollectAdd(data *Collection) error {
	err := Db.Create(&data).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func CollectDel(id int) error {
	err := Db.Delete(&Collection{}, id).Error
	return err
}
