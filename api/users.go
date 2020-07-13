package api

import (
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
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
	r := models.NewUserRepository()
	userId, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	user := r.GetOne(userId)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// AddUser...
func (h *Handler) AddUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	age, _      := strconv.Atoi(user.Age)
	role, _     := strconv.Atoi(user.Role)

	r := models.NewUserRepository()
	err := r.Add(&models.User{
		Nickname: user.Nickname,
		Password: user.Password,
		Age:      age,
		Role:     role,
	})
	if err != "" {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg" : err,
			"id"  : r.Count(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg" : "Created",
		"id"  : r.Count(),
	})
}

// EditUser...
func (h *Handler) EditUser(c *gin.Context) {
	var beforeUser User
	c.BindJSON(&beforeUser)

	id, _ := strconv.Atoi(beforeUser.Id)
	r := models.NewUserRepository()
	user  := r.GetOne(id)

	user.Nickname = beforeUser.Nickname
	if beforeUser.Password == "" {
		password, _ := bcrypt.GenerateFromPassword([]byte(beforeUser.Password), bcrypt.DefaultCost)
		user.Password = string(password)
	}
	user.Age, _ = strconv.Atoi(beforeUser.Age)

	err := r.Edit(user)
	if err != "" {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg" : err,
			"id"  : user.ID,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg" : "Created",
		"id"  : user.ID,
	})
}
