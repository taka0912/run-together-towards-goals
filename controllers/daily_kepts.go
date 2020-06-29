package controllers

import (
	"fmt"
	"github.com/daisuzuki829/run-together-towards-goals/models"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUsers ...
func (h *Handler) GetAllDailyKpts(c *gin.Context) {
	r := models.NewDailyKptRepository()
	dailyKpts := r.GetAll()

	c.HTML(http.StatusOK, "daily_kpts.html", gin.H{
		"dailyKpts": dailyKpts,
	})
}

// AddGenres ...
func (h *Handler) AddDailyKpt(c *gin.Context) {

	session := sessions.Default(c)
	userId := session.Get("UserId").(int)
	//userId := LoginInfo.UserId.(int)
	//userId := int(LoginInfo.UserId.(int))

	fmt.Printf("userId : ")
	spew.Dump(userId)
	fmt.Printf("\n")

	keep, _    := c.GetPostForm("keep")
	problem, _ := c.GetPostForm("problem")
	try, _     := c.GetPostForm("try")

	r := models.NewDailyKptRepository()
	r.Add(&models.DailyKpt{UserID:userId, Keep:keep, Problem:problem, Try:try})

	c.Redirect(http.StatusMovedPermanently, "/_daily_kpts")
}

// DeleteUsers ...
func (h *Handler) DeleteDailyKpt(c *gin.Context) {
	r := models.NewDailyKptRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	r.Delete(id)
	c.Redirect(http.StatusMovedPermanently, "/_daily_kpts")
}

