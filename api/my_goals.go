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
	UserID    string `json:"user_id"`
	Goal      string `json:"goal"`
	GenreID   string `json:"genre_id"`
	LimitDate string `json:"limit_date"`
}

// PostDailyKpt...
func (h *Handler) SetMyGoal(c *gin.Context) {
	var myGoal Goal
	_ = c.BindJSON(&myGoal)

	userID, _ := strconv.Atoi(myGoal.UserID)
	genreID, _ := strconv.Atoi(myGoal.GenreID)

	r := models.NewGoalRepository()
	// TODO
	r.Add(&models.Goal{
		UserID:   userID,
		GoalName: myGoal.Goal,
		GenreID:  genreID,
	})

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Created",
		// TODO
		"id": r.Count(),
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

	newMyGoal.GoalName = goal.Goal
	newMyGoal.GenreID, _ = strconv.Atoi(goal.GenreID)

	r.Edit(newMyGoal)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Created",
		"id":   newMyGoal.ID,
	})
}
