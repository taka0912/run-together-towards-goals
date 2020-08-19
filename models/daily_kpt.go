package models

import (
	"time"

	"github.com/hariNEzuMI928/run-together-towards-goals/redis"
	"github.com/jinzhu/gorm"
	v "gopkg.in/go-playground/validator.v9"
)

// DailyKpt is
type DailyKpt struct {
	gorm.Model
	UserID   int    `gorm:"not null" validate:"required,numeric"`
	Keep     string `gorm:"not null"`
	Problem  string `gorm:"not null"`
	Try      string `gorm:"not null"`
	Good     int    `validate:"numeric"`
	Fight    int    `validate:"numeric"`
	IgnoreMe string `gorm:"-"`
}

// Results ...
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

// NewDailyKptRepository ...
func NewDailyKptRepository() DailyKpt {
	return DailyKpt{}
}

// Add ...
func (o *DailyKpt) Add(dailyKpt *DailyKpt) string {
	validate := v.New()
	err := validate.Struct(dailyKpt)
	if err != nil {
		return err.Error()
	}

	db := Open()
	defer db.Close()
	db.Create(dailyKpt)

	if c, err := redis.Connection(); err == nil {
		var dailyKptList = []string{dailyKpt.Keep, dailyKpt.Problem, dailyKpt.Try}
		redis.SetList(dailyKpt.ID, dailyKptList, c)
	}
	return ""
}

// Edit ...
func (o *DailyKpt) Edit(dailyKpt DailyKpt) {
	db := Open()
	dailyKpt.UpdatedAt = time.Now()
	db.Save(dailyKpt)
	db.Close()
}

// GetAll ...
func (o *DailyKpt) GetAll() []Results {
	db := Open()
	var results []Results
	db.Table("daily_kpts").
		Select("daily_kpts.*, users.nickname").
		Joins("inner JOIN users ON daily_kpts.user_id = users.id").
		Where("daily_kpts.deleted_at is null").
		Order("daily_kpts.id").
		Find(&results)

	db.Close()
	return results
}

// GetOne ...
func (o *DailyKpt) GetOne(id int) DailyKpt {
	db := Open()
	var dailyKpt DailyKpt
	db.First(&dailyKpt, id)
	db.Close()
	return dailyKpt
}

// Delete ...
func (o *DailyKpt) Delete(id int) {
	db := Open()
	var dailyKpt DailyKpt
	db.First(&dailyKpt, id)
	db.Delete(&dailyKpt)
	db.Close()
}

// Count ...
func (o *DailyKpt) Count() int {
	db := Open()
	var count = 0
	db.Table("daily_kpts").Count(&count)
	db.Close()
	return count
}
