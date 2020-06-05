package models

import (
	"github.com/daisuzuki829/run_together_towards_goals/db"
	"github.com/jinzhu/gorm"
	"time"
)

type MyGoal struct {
	gorm.Model
	UserID      int        `gorm:"not null"`
	Goal        string     `gorm:"not null"`
	GenreID     int        `gorm:"not null"`
	LimitDate   time.Time  `sql:"not null;type:date"`
	IgnoreMe    string     `gorm:"-"`
}

// MyGoalRepository is
type MyGoalRepository struct {
}

// NewMyGoalRepository ...
func NewMyGoalRepository() MyGoalRepository {
	return MyGoalRepository{}
}

//DB追加
func (r *MyGoalRepository) Add(myGoal *MyGoal) {
	db := db.Open()
	db.Create(myGoal)
	defer db.Close()
}

//DB更新
func (r *MyGoalRepository) Edit(myGoal MyGoal) {
	db := db.Open()
	db.Save(myGoal)
	db.Close()
}

//DB全取得
func (r *MyGoalRepository) GetAll() []MyGoal {
	db := db.Open()
	var myGoals []MyGoal
	db.Find(&myGoals)
	db.Close()
	return myGoals
}

//DB一つ取得
func (r *MyGoalRepository) GetOne(id int) MyGoal {
	db := db.Open()
	var myGoal MyGoal
	db.First(&myGoal, id)
	db.Close()
	return myGoal
}

//DB削除
func (r *MyGoalRepository) Delete(id int) {
	db := db.Open()
	var myGoal MyGoal
	db.First(&myGoal, id)
	db.Delete(&myGoal)
	db.Close()
}

