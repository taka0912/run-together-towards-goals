package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	v "gopkg.in/go-playground/validator.v9"
	"time"
)

const (
	PublicUser = iota
	AdminUser
)

type User struct {
	gorm.Model
	Nickname     string  `validate:"required,gt=1"`
	Password     string  `validate:"required,gt=4"`
	Age          int     `validate:"numeric"`
	Role         int     `validate:"numeric,oneof=0 1"`
	IgnoreMe     string  `gorm:"-"`
}

// NewUser ...
func NewUserRepository() User {
	return User{}
}

// DB追加
func (o *User) Add(user *User) string {
	db := Open()

	validate := v.New()
	err := validate.Struct(user)
	if err != nil {
		return err.Error()
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)

	db.Create(user)
	defer db.Close()
	return ""
}

// DB更新
func (o *User) Edit(user User) string {
	db := Open()

	validate := v.New()
	err := validate.Struct(user)
	if err != nil {
		return err.Error()
	}
	user.UpdatedAt = time.Now()

	db.Save(user)
	db.Close()
	return ""
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

// GetByName...
func (o *User) GetByName(nickname string) User {
	db := Open()
	var user User
	db.Where("nickname = ?", nickname).First(&user)
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

// Count...
func (o *User) Count() int {
	db := Open()
	var count = 0
	db.Table("users").Count(&count)
	db.Close()
	return count
}

// GetLoginUser
func (o *User) GetLoginUser(id interface{}) User {
	db := Open()
	var user User
	db.First(&user, id)
	db.Close()
	return user
}

