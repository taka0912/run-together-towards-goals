package models

import (
	"github.com/daisuzuki829/run_together_towards_goals/db"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Nickname     string  `gorm:"not null"`
	Password     string  `gorm:"not null"`
	Age          int
	Role         string  `gorm:"size:255"`
	IgnoreMe     string  `gorm:"-"`
}

// UserRepository is
type UserRepository struct {
}

// NewUserRepository ...
func NewUserRepository() UserRepository {
	return UserRepository{}
}

//DB追加
func (r *UserRepository) Add(user *User) {
	db := db.Open()
	db.Create(user)
	defer db.Close()
}

//DB更新
func (r *UserRepository) Edit(user User) {
	db := db.Open()
	db.Save(user)
	db.Close()
}

//DB全取得
func (r *UserRepository) GetAll() []User {
	db := db.Open()
	var users []User
	db.Find(&users)
	db.Close()
	return users
}

//DB一つ取得
func (r *UserRepository) GetOne(id int) User {
	db := db.Open()
	var user User
	db.First(&user, id)
	db.Close()
	return user
}

//DB削除
func (r *UserRepository) Delete(id int) {
	db := db.Open()
	var user User
	db.First(&user, id)
	db.Delete(&user)
	db.Close()
}

