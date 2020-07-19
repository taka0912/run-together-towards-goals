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

// GetAllUsers...
func (h *Handler) GetAllUsers(c *gin.Context) {
	r := models.NewUserRepository()
	users := r.GetAll()

	c.HTML(http.StatusOK, "users.html", gin.H{
		"users": users,
	})
}

// AddUsers...
func (h *Handler) AddUser(c *gin.Context) {
	r := models.NewUserRepository()

	nickname, _ := c.GetPostForm("nickname")
	password, _ := c.GetPostForm("password")
	age, _ := c.GetPostForm("age")
	ageFmt, _ := strconv.Atoi(age)
	role, _ := c.GetPostForm("role")
	roleFmt, _ := strconv.Atoi(role)

	// TODO
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

// GetUsers...
func (h *Handler) GetUser(c *gin.Context) {
	r := models.NewUserRepository()
	rg := models.NewGenreRepository()

	id, _ := strconv.Atoi(c.Param("id"))
	user := r.GetAllInfo(id)

	loginUser := r.GetLoginUser(sessions.Default(c).Get("UserId"))
	adminFlag := false
	// 「ログインユーザが管理者、もしくは編集ページのユーザー本人」の場合
	if loginUser.Role == models.AdminUser || loginUser.ID == user.ID {
		adminFlag = true
	}

	c.HTML(http.StatusOK, "user_view.html", gin.H{
		"user":      user,
		"genres":    rg.GetAll(),
		"adminFlag": adminFlag,
	})
}

// EditUsers...
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

// DeleteUsers...
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

	rg := models.NewGoalRepository()
	//rg.UserID = int(r.ID)
	genreID, _ := c.GetPostForm("genre_id")
	rg.GenreID, _ = strconv.Atoi(genreID)
	rg.GoalName, _ = c.GetPostForm("goal")
	rg.Add(&rg)

	rt := models.NewTodoListRepository()
	rt.GoalID = int(rg.ID)
	rt.RequiredElements, _ = c.GetPostForm("required_elements")
	rt.SpecificGoal, _ = c.GetPostForm("specific_goal")
	limitDate, _ := c.GetPostForm("limit_date")
	rt.LimitDate, _ = time.Parse("2006-01-02", limitDate)
	rt.LimitDate, _ = time.Parse("2006-01-02", limitDate)
	rt.Add(&rt)

	c.HTML(http.StatusOK, "login.html", gin.H{
		"msg": "Welcome! Let's Login.",
	})
}

// GetUsers...
func (h *Handler) GetMyPage(c *gin.Context) {
	r := models.NewUserRepository()

	loginUser := r.GetLoginUser(sessions.Default(c).Get("UserId"))
	user := r.GetAllInfo(int(loginUser.ID))

	rg := models.NewGenreRepository()
	genres := rg.GetAll()

	c.HTML(http.StatusOK, "my_page.html", gin.H{
		"user":   &user,
		"genres": genres,
	})
}

// EditMyPage...
func (h *Handler) EditMyPage(c *gin.Context) {
	r := models.NewUserRepository()
	loginUser := r.GetLoginUser(sessions.Default(c).Get("UserId"))
	user := r.GetAllInfo(int(loginUser.ID))

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
		c.HTML(http.StatusOK, "my_page.html", gin.H{
			"user": user,
			"err":  err,
		})
		return
	}

	h.GetMyPage(c)
}

// EditGoal...
func (h *Handler) EditGoal(c *gin.Context) {
	r := models.NewGoalRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	goal := r.GetOne(id)

	tail := "_" + c.Param("id")

	genreId, _ := c.GetPostForm("genre_id" + tail)
	goal.GenreID, _ = strconv.Atoi(genreId)
	goal.GoalName, _ = c.GetPostForm("goal_name" + tail)

	r.Edit(goal)
	h.GetMyPage(c)
}

// DeleteGoal...
func (h *Handler) DeleteGoal(c *gin.Context) {
	r := models.NewGoalRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	r.Delete(id)
	h.GetMyPage(c)
}

// EditTodo...
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

// DeleteTodo...
func (h *Handler) DeleteTodo(c *gin.Context) {
	r := models.NewTodoListRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	r.Delete(id)
	h.GetMyPage(c)
}

// AddGoal...
func (h *Handler) AddGoal(c *gin.Context) {
	r := models.NewGoalRepository()

	ur := models.NewUserRepository()
	loginUser := ur.GetLoginUser(sessions.Default(c).Get("UserId"))
	r.UserID = int(loginUser.ID)
	// TODO
	//loginUserId, err := GetLoginUserId(c)
	//if err != nil {
	//	c.Redirect(http.StatusMovedPermanently, "/logout")
	//}

	genreId, _ := c.GetPostForm("genre_id")
	r.GenreID, _ = strconv.Atoi(genreId)
	r.GoalName, _ = c.GetPostForm("goal_name")

	r.Add(&r)
	h.GetMyPage(c)
}

// AddTodo...
func (h *Handler) AddTodo(c *gin.Context) {
	r := models.NewTodoListRepository()

	goalId, _ := c.GetPostForm("goal_id")
	r.GoalID, _ = strconv.Atoi(goalId)
	r.RequiredElements, _ = c.GetPostForm("required_elements")
	r.Todo, _ = c.GetPostForm("todo")
	r.SpecificGoal, _ = c.GetPostForm("specific_goal")
	limitDate, _ := c.GetPostForm("limit_date")
	r.LimitDate, _ = time.Parse("2006-01-02", limitDate)

	r.Add(&r)
	h.GetMyPage(c)
}
