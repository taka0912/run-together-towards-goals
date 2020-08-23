package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	v "gopkg.in/go-playground/validator.v9"
)

const (
	// PublicUser ...
	PublicUser = iota
	// AdminUser ...
	AdminUser
)

// User ...
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

// NewUserRepository ...
func NewUserRepository() User {
	return User{}
}

// Add ...
// TODO：Addの引数は要らない。r.Add()だけで良い
func (o *User) Add(user *User) []string {
	db := Open()
	defer db.Close()

	err := validateUser(user)
	if err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)

	db.Create(user)
	return nil
}

// Edit ...
func (o *User) Edit(user User) []string {
	db := Open()
	defer db.Close()

	err := validateUser(&user)
	if err != nil {
		return err
	}
	user.UpdatedAt = time.Now()

	db.Save(user)
	return nil
}

// GetAll ...
func (o *User) GetAll() []User {
	var users []User
	db := Open()
	defer db.Close()
	db.Find(&users)
	return users
}

// GetOne ...
func (o *User) GetOne(id int) User {
	var user User
	db := Open()
	defer db.Close()
	db.First(&user, id)
	return user
}

// GetAllInfo ...
func (o *User) GetAllInfo(id int) User {
	var user User
	db := Open()
	defer db.Close()
	db.Preload("Goals").Preload("Goals.TodoLists").Find(&user, id)
	return user
}

// GetByName ...
func (o *User) GetByName(nickname string) (User, string) {
	var user User
	db := Open()
	defer db.Close()
	db.Where("nickname = ?", nickname).First(&user)
	if user.ID == 0 {
		return User{}, "Not found"
	}
	return user, ""
}

// Delete ...
func (o *User) Delete(id int) {
	var user User
	db := Open()
	defer db.Close()
	db.First(&user, id)
	db.Delete(&user)
}

// Count ...
func (o *User) Count() int {
	db := Open()
	defer db.Close()
	var count = 0
	db.Table("users").Count(&count)
	return count
}

// GetLoginUser ...
func (o *User) GetUserByInterfaceID(id interface{}) User {
	var user User
	db := Open()
	defer db.Close()
	db.First(&user, id)
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
