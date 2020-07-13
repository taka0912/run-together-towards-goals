package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// DailyKpt is
type DailyKpt struct {
	gorm.Model
	UserID    int     `gorm:"not null"`
	Keep      string  `gorm:"not null"`
	Problem   string  `gorm:"not null"`
	Try       string  `gorm:"not null"`
	Good      int
	Fight     int
	IgnoreMe  string  `gorm:"-"`
}

type Results struct {
	ID        int
	UserID    int
	CreatedAt time.Time
	Nickname  string
	Keep      string
	Problem   string
	Try       string
	Good      int
	Fight     int
}

// NewUser...
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
	dailyKpt.UpdatedAt = time.Now()
	db.Save(dailyKpt)
	db.Close()
}

// DB全取得
func (o *DailyKpt) GetAll() []Results {
	db := Open()
	var results []Results
	db.Table("daily_kpts").
		Select("daily_kpts.*, users.nickname").
		Joins("inner JOIN users ON daily_kpts.user_id = users.id").
		Where("daily_kpts.deleted_at is null").
		Find(&results)

	db.Close()
	return results
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

// Count...
func (o *DailyKpt) Count() int {
	db := Open()
	var count = 0
	db.Table("daily_kpts").Count(&count)
	db.Close()
	return count
}