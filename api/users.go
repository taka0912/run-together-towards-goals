package api

import (
	"github.com/daisuzuki829/run_together_towards_goals/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type User struct {
	Id        string  `json:"id"`
	Nickname  string  `json:"nickname"`
	Password  string  `json:"password"`
	Age       string  `json:"age"`
	Role      string  `json:"role"`
}

// GetUser...
func (h *Handler) GetUser(c *gin.Context) {
	userId := c.DefaultQuery("id", "1")

	r := models.NewUserRepository()
	userIdFmt, _ := strconv.Atoi(userId)
	user := r.GetOne(userIdFmt)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// AddUser...
func (h *Handler) AddUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	passwordFmt := string(password)
	age, _      := strconv.Atoi(user.Age)
	role, _     := strconv.Atoi(user.Role)

	r := models.NewUserRepository()
	r.Add(&models.User{
		Nickname: user.Nickname,
		Password: passwordFmt,
		Age:      age,
		Role:     role,
	})

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg" : "Created",
		"id"  : r.Count(),
	})
}

