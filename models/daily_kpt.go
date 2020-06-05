package models

import (
	"github.com/daisuzuki829/run_together_towards_goals/db"
	"github.com/jinzhu/gorm"
)

// DailyKpt is
type DailyKpt struct {
	gorm.Model
	Keep      string  `gorm:"not null"`
	Problem   string  `gorm:"not null"`
	Try       string  `gorm:"not null"`
	Good      int
	Figh      int
	IgnoreMe  string  `gorm:"-"`
}

// NewDailyKpt ...
func NewDailyKpt() DailyKpt {
	return DailyKpt{}
}

//DB追加
func (r *DailyKpt) Add(dailyKpt *DailyKpt) {
	db := db.Open()
	db.Create(dailyKpt)
	defer db.Close()
}

//DB更新
func (r *DailyKpt) Edit(dailyKpt DailyKpt) {
	db := db.Open()
	db.Save(dailyKpt)
	db.Close()
}

//DB全取得
func (r *DailyKpt) GetAll() []DailyKpt {
	db := db.Open()
	var dailyKpts []DailyKpt
	db.Find(&dailyKpts)
	db.Close()
	return dailyKpts
}

//DB一つ取得
func (r *DailyKpt) GetOne(id int) DailyKpt {
	db := db.Open()
	var dailyKpt DailyKpt
	db.First(&dailyKpt, id)
	db.Close()
	return dailyKpt
}

//DB削除
func (r *DailyKpt) Delete(id int) {
	db := db.Open()
	var dailyKpt DailyKpt
	db.First(&dailyKpt, id)
	db.Delete(&dailyKpt)
	db.Close()
}

