package models

import (
	"github.com/jinzhu/gorm"
)

// DailyKpt is
type DailyKpt struct {
	gorm.Model
	Keep      string  `gorm:"not null"`
	Problem   string  `gorm:"not null"`
	Try       string  `gorm:"not null"`
	Good      int
	Fight     int
	IgnoreMe  string  `gorm:"-"`
}

// NewUser ...
func NewDailyKptRepository() DailyKpt {
	return DailyKpt{}
}

// DB追加
func (o *DailyKpt) Add(dailyKpt *DailyKpt) {
	db := Open()
	db.Create(dailyKpt)
	defer db.Close()
}

// DB更新
func (o *DailyKpt) Edit(dailyKpt DailyKpt) {
	db := Open()
	db.Save(dailyKpt)
	db.Close()
}

// DB全取得
func (o *DailyKpt) GetAll() []DailyKpt {
	db := Open()
	var dailyKpts []DailyKpt
	db.Find(&dailyKpts)
	db.Close()
	return dailyKpts
}

// DB一つ取得
func (o *DailyKpt) GetOne(id int) DailyKpt {
	db := Open()
	var dailyKpt DailyKpt
	db.First(&dailyKpt, id)
	db.Close()
	return dailyKpt
}

// DB削除
func (o *DailyKpt) Delete(id int) {
	db := Open()
	var dailyKpt DailyKpt
	db.First(&dailyKpt, id)
	db.Delete(&dailyKpt)
	db.Close()
}
