package controllers

import (
	"github.com/daisuzuki829/run_together_towards_goals/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

// GetUsers ...
func (h *Handler) GetAllUsers(c *gin.Context) {
	r := models.NewUserRepository()
	users := r.GetAll()

	c.HTML(http.StatusOK, "users.html", gin.H{
		"users": users,
	})
}

// AddUsers ...
func (h *Handler) AddUser(c *gin.Context) {
	r := models.NewUserRepository()

	nickname, _ := c.GetPostForm("nickname")
	password, _ := c.GetPostForm("password")
	age, _      := c.GetPostForm("age")
	ageFmt, _   := strconv.Atoi(age)
	role, _     := c.GetPostForm("role")
	roleFmt, _  := strconv.Atoi(role)

	r.Add(&models.User{Nickname: nickname, Password: password, Age: ageFmt, Role: roleFmt})
	c.Redirect(http.StatusMovedPermanently, "/_users")
}

// GetUsers ...
func (h *Handler) GetUser(c *gin.Context) {
	r := models.NewUserRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	user := r.GetOne(id)
	c.HTML(http.StatusOK, "user_edit.html", gin.H{
		"user": user,
	})
}

// EditUsers ...
func (h *Handler) EditUser(c *gin.Context) {
	r     := models.NewUserRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	user  := r.GetOne(id)

	user.Nickname, _ = c.GetPostForm("nickname")
	password, _      := c.GetPostForm("password")
	if password != "" {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user.Password = string(hashPassword)
	}
	age,      _ := c.GetPostForm("age")
	user.Age, _ = strconv.Atoi(age)
	role,     _ := c.GetPostForm("role")
	user.Role, _ = strconv.Atoi(role)
	user.UpdatedAt = time.Now()
	r.Edit(user)

	c.Redirect(http.StatusMovedPermanently, "/_users")
}

// DeleteUsers ...
func (h *Handler) DeleteUser(c *gin.Context) {
	r := models.NewUserRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	r.Delete(id)
	c.Redirect(http.StatusMovedPermanently, "/_users")
}

// LoginUser...
func LoginUser(c *gin.Context) (models.User, string) {
	nickname, _ := c.GetPostForm("nickname")
	password, _ := c.GetPostForm("password")

	r := models.NewUserRepository()
	user := r.GetByName(nickname)

	if user.Role == models.PublicUser {
		return user, "You do not have authority"
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, "Not registered or Password is incorrect"
	}

	return user, ""
}
