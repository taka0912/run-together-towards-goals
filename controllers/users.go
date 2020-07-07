package controllers

import (
	"fmt"
	"github.com/daisuzuki829/run-together-towards-goals/models"
	"github.com/davecgh/go-spew/spew"
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
	age, _ := c.GetPostForm("age")
	ageFmt, _ := strconv.Atoi(age)
	role, _ := c.GetPostForm("role")
	roleFmt, _ := strconv.Atoi(role)

	err := r.Add(&models.User{Nickname: nickname, Password: password, Age: ageFmt, Role: roleFmt})
	users := r.GetAll()
	if err != "" {
		c.HTML(http.StatusOK, "users.html", gin.H{
			"err":   err,
			"users": users,
		})
	}
	c.Redirect(http.StatusMovedPermanently, "/_users")
}

// GetUsers ...
func (h *Handler) GetUser(c *gin.Context) {
	r := models.NewUserRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	user := r.GetAllInfo(id)

	//rg := models.NewMyGoalRepository()
	//myGoals := rg.GetByUserId(user.ID)
	//user.MyGoals = rg.GetByUserId(user.ID)

	fmt.Printf("r.GetAllInfo(id) : ")
	spew.Dump(user)
	fmt.Printf("\n")

	//user.IgnoreMe = user.LimitDate.Format("2006-01-02")
	//
	//rg := models.NewGenreRepository()
	//genres := rg.GetAll()
	//
	//loginUser := r.GetLoginUser(c)
	//adminFlag := false
	//if loginUser.Role == models.AdminUser {
	//	adminFlag = true
	//}

	c.HTML(http.StatusOK, "user_view.html", gin.H{
		"user":   user,
		//"genres": genres,
		//"adminFlag": adminFlag,
	})
}

// EditUsers ...
func (h *Handler) EditUser(c *gin.Context) {
	r := models.NewUserRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	user := r.GetOne(id)

	loginUser := r.GetLoginUser(c)
	// 「ログインユーザが管理者、もしくは編集ページのユーザー本人」でない場合
	if !(loginUser.Role == models.AdminUser || loginUser.ID == user.ID) {
		c.HTML(http.StatusOK, "user_view.html", gin.H{
			"err": "Unauthorized",
		})
		return
	}

	user.Nickname, _ = c.GetPostForm("nickname")
	password, _ := c.GetPostForm("password")
	if password != "" {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user.Password = string(hashPassword)
	}
	age, _ := c.GetPostForm("age")
	user.Age, _ = strconv.Atoi(age)
	role, _ := c.GetPostForm("role")
	user.Role, _ = strconv.Atoi(role)

	err := r.Edit(user)
	if err != "" {
		c.HTML(http.StatusOK, "user_edit.html", gin.H{
			"err": err,
		})
		return
	}

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

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, "Not registered or Password is incorrect"
	}

	return user, ""
}

// NewRegistration...
func NewRegistration(c *gin.Context) {
	r := models.NewUserRepository()
	r.Nickname, _ = c.GetPostForm("nickname")
	r.Password, _ = c.GetPostForm("password")
	age, _ := c.GetPostForm("age")
	r.Age, _ = strconv.Atoi(age)
	role, _ := c.GetPostForm("role")
	r.Role, _ = strconv.Atoi(role)
	err := r.Add(&r)
	if err != "" {
		c.HTML(http.StatusMovedPermanently, "registration.html", gin.H{
			"err": err,
		})
		return
	}

	rg := models.NewMyGoalRepository()
	rg.UserID = int(r.ID)
	genreID, _ := c.GetPostForm("genre_id")
	rg.GenreID, _ = strconv.Atoi(genreID)
	rg.Goal, _ = c.GetPostForm("goal")
	rg.Add(&rg)

	rt := models.NewTodoRepository()
	rt.GoalId = int(rg.ID)
	rt.RequiredElements, _ = c.GetPostForm("required_elements")
	rt.SpecificGoal, _ = c.GetPostForm("specific_goal")
	limitDate, _ := c.GetPostForm("limit_date")
	rt.LimitDate, _ = time.Parse("2006-01-02", limitDate)
	rt.Add(&rt)

	c.HTML(http.StatusOK, "login.html", gin.H{
		"msg": "Welcome! Let's Login.",
	})
}

// GetUsers ...
func (h *Handler) GetMyPage(c *gin.Context) {
	r := models.NewUserRepository()

	loginUser := r.GetLoginUser(c)

	fmt.Printf("loginUser : ")
	spew.Dump(loginUser)
	fmt.Printf("\n")

	user := r.GetAllInfo(int(loginUser.ID))
	//user.IgnoreMe = user.LimitDate.Format("2006-01-02")

	rg := models.NewGenreRepository()
	genres := rg.GetAll()

	c.HTML(http.StatusOK, "my_page.html", gin.H{
		"user":   user,
		"genres": genres,
	})
}