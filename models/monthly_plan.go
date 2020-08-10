package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type MonthlyPlan struct {
	gorm.Model
	UserID             int       `gorm:"not null" validate:"required,numeric"`
	GoalID             int       `gorm:"not null" validate:"required,numeric"`
	Month              time.Time `gorm:"not null" sql:"not null;type:date"`
	KeepInLastMonth    string
	ProblemInLastMonth string
	GoalAfterHalfYear  string
	GoalInThisMonth    string
	CurrentState       string
	DailyTodo          string
}

// NewMonthlyPlanRepository...
func NewMonthlyPlanRepository() MonthlyPlan {
	return MonthlyPlan{}
}

// DB追加
func (o *MonthlyPlan) Add(monthlyPlan *MonthlyPlan) {
	db := Open()
	db.Create(monthlyPlan)
	defer db.Close()
}

// DB更新
func (o *MonthlyPlan) Edit(monthlyPlan MonthlyPlan) {
	db := Open()
	db.Save(monthlyPlan)
	db.Close()
}

// DB全取得
func (o *MonthlyPlan) GetAll() []MonthlyPlan {
	db := Open()
	var monthlyPlans []MonthlyPlan
	db.Find(&monthlyPlans)
	db.Close()
	return monthlyPlans
}

// DB一つ取得
func (o *MonthlyPlan) GetOne(id int) MonthlyPlan {
	db := Open()
	var monthlyPlan MonthlyPlan
	db.Find(&monthlyPlan, id)
	db.Close()
	return monthlyPlan
}

// DB削除
func (o *MonthlyPlan) Delete(id int) {
	db := Open()
	var monthlyPlan MonthlyPlan
	db.First(&monthlyPlan, id)
	db.Delete(&monthlyPlan)
	db.Close()
}

// Count...
func (o *MonthlyPlan) Count() int {
	db := Open()
	var count = 0
	db.Table("monthly_plans").Count(&count)
	db.Close()
	return count
}
