package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
	"github.com/jinzhu/gorm"
)

type Goal struct {
	gorm.Model
	UserID   string `json:"user_id"`
	GenreID  string `json:"genre_id"`
	GoalName string `json:"goal_name"`
}

// PostDailyKpt...
func (h *Handler) SetMyGoal(c *gin.Context) {
	var apiMyGoal Goal
	_ = c.BindJSON(&apiMyGoal)

	r := models.NewGoalRepository()
	r.UserID, _ = strconv.Atoi(apiMyGoal.UserID)
	r.GenreID, _ = strconv.Atoi(apiMyGoal.GenreID)
	r.GoalName = apiMyGoal.GoalName
	r.Add(&r)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Created",
		"id":   r.ID,
	})
}

// EditMyGoal...
func (h *Handler) EditMyGoal(c *gin.Context) {
	r := models.NewGoalRepository()

	id := c.DefaultQuery("id", "0")
	myGoalId, _ := strconv.Atoi(id)
	newMyGoal := r.GetOne(myGoalId)

	if newMyGoal.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  "Not Found",
		})
		return
	}

	var goal Goal
	_ = c.BindJSON(&goal)

	newMyGoal.GoalName = goal.GoalName
	newMyGoal.GenreID, _ = strconv.Atoi(goal.GenreID)

	r.Edit(newMyGoal)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Created",
		"id":   newMyGoal.ID,
	})
}
