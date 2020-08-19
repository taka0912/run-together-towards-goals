package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
	"golang.org/x/crypto/bcrypt"
)

// GetAllUsers ...
func (h *Handler) GetAllUsers(c *gin.Context) {
	r := models.NewUserRepository()
	users := r.GetAll()

	c.HTML(http.StatusOK, "users.html", gin.H{
		"users": users,
	})
}

// AddUser ...
func (h *Handler) AddUser(c *gin.Context) {
	r := models.NewUserRepository()

	r.Nickname, _ = c.GetPostForm("nickname")
	r.Password, _ = c.GetPostForm("password")
	r.Role = models.PublicUser

	err := r.Add(&r)
	users := r.GetAll()
	if err != nil {
		c.HTML(http.StatusOK, "users.html", gin.H{
			"errs":  err,
			"users": users,
		})
	}
	c.Redirect(http.StatusMovedPermanently, "/_users")
}

// GetUser ...
func (h *Handler) GetUser(c *gin.Context) {
	r := models.NewUserRepository()
	rg := models.NewGenreRepository()

	id, _ := strconv.Atoi(c.Param("id"))
	user := r.GetAllInfo(id)
	loginUser := r.GetLoginUser(sessions.Default(c).Get("UserId"))

	c.HTML(http.StatusOK, "user_view.html", gin.H{
		"user":      user,
		"genres":    rg.GetAll(),
		"adminFlag": loginUser.Role,
		"loginFlag": loginUser.ID == user.ID,
	})
}

// EditUser ...
func (h *Handler) EditUser(c *gin.Context) {
	r := models.NewUserRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	user := r.GetOne(id)

	user.Nickname, _ = c.GetPostForm("nickname")
	password, _ := c.GetPostForm("password")
	if password != "" {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user.Password = string(hashPassword)
	}
	role, _ := c.GetPostForm("role")
	user.Role, _ = strconv.Atoi(role)

	err := r.Edit(user)
	if err != nil {
		c.HTML(http.StatusOK, "user_edit.html", gin.H{
			"err": err,
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/_users")
}

// DeleteUser ...
func (h *Handler) DeleteUser(c *gin.Context) {
	r := models.NewUserRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	r.Delete(id)
	c.Redirect(http.StatusMovedPermanently, "/_users")
}

// LoginUser ...
func LoginUser(c *gin.Context) (models.User, string) {
	nickname, _ := c.GetPostForm("nickname")
	password, _ := c.GetPostForm("password")

	r := models.NewUserRepository()
	user, _ := r.GetByName(nickname)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, "Not registered or Password is incorrect"
	}

	return user, ""
}

// NewRegistration ...
func NewRegistration(c *gin.Context) {
	r := models.NewUserRepository()
	r.Nickname, _ = c.GetPostForm("nickname")
	r.Password, _ = c.GetPostForm("password")
	role, _ := c.GetPostForm("role")
	r.Role, _ = strconv.Atoi(role)
	err := r.Add(&r)

	if err != nil {
		c.HTML(http.StatusMovedPermanently, "registration.html", gin.H{
			"err": err,
		})
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{
		"msg": "Welcome! Let's Login.",
	})
}

// GetMyPage ...
func (h *Handler) GetMyPage(c *gin.Context) {
	r := models.NewUserRepository()

	loginUserID, _ := GetloginUserID(c)
	user := r.GetAllInfo(loginUserID)

	rg := models.NewGenreRepository()
	genres := rg.GetAll()

	c.HTML(http.StatusOK, "my_page.html", gin.H{
		"user":   user,
		"genres": genres,
	})
}

// EditMyPage ...
func (h *Handler) EditMyPage(c *gin.Context) {
	r := models.NewUserRepository()
	user := r.GetLoginUser(sessions.Default(c).Get("UserId"))

	user.Nickname, _ = c.GetPostForm("nickname")
	password, _ := c.GetPostForm("password")
	if password != "" {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user.Password = string(hashPassword)
	}

	user.Age, _ = c.GetPostForm("age")
	ageDisplayFlag, _ := c.GetPostForm("age_display_flag")
	user.AgeDisplayFlag, _ = strconv.Atoi(ageDisplayFlag)
	user.Address, _ = c.GetPostForm("address")
	user.BirthPlace, _ = c.GetPostForm("birth_place")
	user.Hobby, _ = c.GetPostForm("hobby")
	user.Occupation, _ = c.GetPostForm("occupation")
	user.StrongPoint, _ = c.GetPostForm("strong_point")
	user.Skill, _ = c.GetPostForm("skill")

	err := r.Edit(user)

	if err != nil {
		c.HTML(http.StatusOK, "my_page.html", gin.H{
			"user": user,
			"err":  err,
		})
		return
	}

	h.GetMyPage(c)
}

// EditGoal ...
func (h *Handler) EditGoal(c *gin.Context) {
	r := models.NewGoalRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	goal := r.GetOne(id)

	tail := "_" + c.Param("id")

	genreID, _ := c.GetPostForm("genre_id" + tail)
	goal.GenreID, _ = strconv.Atoi(genreID)
	goal.GoalName, _ = c.GetPostForm("goal_name" + tail)
	displayFlag, _ := c.GetPostForm("display_flag" + tail)
	goal.DisplayFlag, _ = strconv.Atoi(displayFlag)

	r.Edit(goal)
	h.GetMyPage(c)
}

// DeleteGoal ...
func (h *Handler) DeleteGoal(c *gin.Context) {
	r := models.NewGoalRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	r.Delete(id)
	h.GetMyPage(c)
}

// EditTodo ...
func (h *Handler) EditTodo(c *gin.Context) {
	r := models.NewTodoListRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	goal := r.GetOne(id)

	tail := "_" + c.Param("id")

	goal.RequiredElements, _ = c.GetPostForm("required_elements" + tail)
	goal.SpecificGoal, _ = c.GetPostForm("specific_goal" + tail)
	limitDate, _ := c.GetPostForm("limit_date" + tail)
	goal.LimitDate, _ = time.Parse("2006-01-02", limitDate)

	r.Edit(goal)
	h.GetMyPage(c)
}

// DeleteTodo ...
func (h *Handler) DeleteTodo(c *gin.Context) {
	r := models.NewTodoListRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	r.Delete(id)
	h.GetMyPage(c)
}

// AddGoal ...
func (h *Handler) AddGoal(c *gin.Context) {
	r := models.NewGoalRepository()

	r.UserID, _ = GetloginUserID(c)
	genreID, _ := c.GetPostForm("genre_id")
	r.GenreID, _ = strconv.Atoi(genreID)
	r.GoalName, _ = c.GetPostForm("goal_name")

	r.Add(&r)
	h.GetMyPage(c)
}

// AddTodo ...
func (h *Handler) AddTodo(c *gin.Context) {
	r := models.NewTodoListRepository()

	goalID, _ := c.GetPostForm("goal_id")
	r.GoalID, _ = strconv.Atoi(goalID)
	r.RequiredElements, _ = c.GetPostForm("required_elements")
	r.Todo, _ = c.GetPostForm("todo")
	r.SpecificGoal, _ = c.GetPostForm("specific_goal")
	limitDate, _ := c.GetPostForm("limit_date")
	r.LimitDate, _ = time.Parse("2006-01-02", limitDate)

	r.Add(&r)
	h.GetMyPage(c)
}
