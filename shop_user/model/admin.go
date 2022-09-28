package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string `gorm:"type:string;not null;unique"`
	Password string `gorm:"type:string"`
}

func AdminCheckLogin(name string, password string) bool {
	var count int64
	err := Db.Model(&Admin{}).Select("Name=? AND Password=?", name, password).Count(&count).Error
	return err == nil && count > 0
}

func UserList(pageNum int, pageSize int) (users []User, count int64, err error) {
	err = Db.Offset((pageNum - 1) * pageSize).Limit(pageNum).Find(&users).Count(&count).Error
	return
}

func UserDel(id int) error {
	err := Db.Delete(&User{}, id).Error
	return err
}
