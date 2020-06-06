package models

import (
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

// NewUser ...
func NewUserRepository() User {
	return User{}
}

// DB追加
func (o *User) Add(user *User) {
	db := Open()
	db.Create(user)
	defer db.Close()
}

// DB更新
func (o *User) Edit(user User) {
	db := Open()
	db.Save(user)
	db.Close()
}

// DB全取得
func (o *User) GetAll() []User {
	db := Open()
	var users []User
	db.Find(&users)
	db.Close()
	return users
}

// DB一つ取得
func (o *User) GetOne(id int) User {
	db := Open()
	var user User
	db.First(&user, id)
	db.Close()
	return user
}

// DB削除
func (o *User) Delete(id int) {
	db := Open()
	var user User
	db.First(&user, id)
	db.Delete(&user)
	db.Close()
}

