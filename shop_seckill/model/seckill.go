package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Seckill struct {
	Id        int     `gorm:"type:int;primaryKey;Auto_increment"`
	SecPrice  float32 `gorm:"type:decimal(11,2)"`
	SecNum    int     `gorm:"type:int"`
	Pid       int     `gorm:"type:int;not null"`
	SecDesc   string  `gorm:"type:string;"`
	StartTime time.Time
	EndTime   time.Time
}

func SeckillInfo(id int) (seckill *Seckill, err error) {
	err = Db.Where("Id=?", id).First(&seckill).Error
	if err != nil {
		fmt.Println(err)
	}
	return
}

func SeckillSuccess(email string, id int) error {
	tx := Db.Begin()
	err := tx.Model(Seckill{}).Where("Id=?", id).Update("Sec_Num", gorm.Expr("Sec_Num-1")).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Create(&Order{
		Email: email,
		Sid:   id,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
