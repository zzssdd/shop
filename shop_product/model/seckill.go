package model

import "time"

type Seckill struct {
	Id        int     `gorm:"type:int;primaryKey;Auto_increment"`
	SecPrice  float32 `gorm:"type:decimal(11,2)"`
	SecNum    int     `gorm:"type:int"`
	Pid       int     `gorm:"type:int;not null"`
	SecDesc   string  `gorm:"type:string;"`
	StartTime time.Time
	EndTime   time.Time
	Pro       Product `gorm:"foreignkey:Pid"`
}

func SeckillList(currentPage int, pageSize int) (seckills []*Seckill, count int64, err error) {
	err = Db.Model(&Seckill{}).Count(&count).Error
	err = Db.Preload("Pro").Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&seckills).Error
	return
}

func SeckillAdd(seckill *Seckill) error {
	err := Db.Create(&seckill).Error
	return err
}

func SeckillDel(id int) error {
	err := Db.Delete(&Seckill{}, id).Error
	return err
}

func SeckillInfo(id int) (seckill *Seckill, err error) {
	err = Db.Where("Id=?", id).First(&seckill).Error
	return
}

func SeckillEdit(data *Seckill) error {
	err := Db.Updates(&data).Error
	return err
}
