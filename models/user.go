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
	Nickname       string `validate:"required,gt=1"`
	Password       string `validate:"required,gt=4"`
	Age            string
	AgeDisplayFlag int
	Address        string
	BirthPlace     string
	Hobby          string
	Occupation     string
	StrongPoint    string
	Skill          string
	Role           int    `validate:"numeric,oneof=0 1"`
	IgnoreMe       string `gorm:"-"`
	Goals          []Goal
}

// NewUser...
func NewUserRepository() User {
	return User{}
}

// DB追加
// TODO：Addの引数は要らない。r.Add()だけで良い
func (o *User) Add(user *User) []string {
	db := Open()

	err := validateUser(user)
	if err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)

	db.Create(user)
	defer db.Close()
	return nil
}

// DB更新
func (o *User) Edit(user User) []string {
	db := Open()

	err := validateUser(&user)
	if err != nil {
		return err
	}
	user.UpdatedAt = time.Now()

	db.Save(user)
	db.Close()
	return nil
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

// DB一つ取得
func (o *User) GetAllInfo(id int) User {
	db := Open()
	var user User
	db.Preload("Goals").Preload("Goals.TodoLists").Find(&user, id)
	db.Close()
	return user
}

// GetByName...
func (o *User) GetByName(nickname string) (User, string) {
	db := Open()
	var user User
	db.Where("nickname = ?", nickname).First(&user)
	db.Close()
	if user.ID == 0 {
		return User{}, "Not found"
	}
	return user, ""
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

func validateUser(user *User) []string {
	var errorMessages []string

	validate := v.New()
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(v.ValidationErrors) {
			var errorMessage string
			fieldName := err.Field() //バリデーションでNGになった変数名を取得

			switch fieldName {
			case "Nickname":
				errorMessage = "Nickname is required"
			case "Password":
				errorMessage = "Password must be 4 or more alphanumeric characters"
			case "Role":
				errorMessage = "Role is 0 or 1 only"
			}
			errorMessages = append(errorMessages, errorMessage)
		}
	}

	// 既にそのニックネームが使われている場合(unique判定)
	//TODO：idではなくてstringのunique判定のやり方
	r := NewUserRepository()
	if tmpUser, err := r.GetByName(user.Nickname); tmpUser.ID != user.ID && err == "" {
		errorMessages = append(errorMessages, "The nickname is already registered or used by another user")
	}
	return errorMessages
}
