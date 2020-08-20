package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
)

// GetAllMonthlyPlans ...
func (h *Handler) GetAllMonthlyPlans(c *gin.Context) {
	r := models.NewMonthlyPlanRepository()
	monthlyPlans := r.GetAll()

	loginUserID := GetLoginUserID(c)

	rg := models.NewGoalRepository()
	goals := rg.GetByUserID(loginUserID)

	c.HTML(http.StatusOK, "monthly_plans.html", gin.H{
		"monthlyPlans": monthlyPlans,
		"goals":        goals,
	})
}

// AddMonthlyPlans ...
func (h *Handler) AddMonthlyPlans(c *gin.Context) {
	r := models.NewMonthlyPlanRepository()
	loginUserID := GetLoginUserID(c)

	r.UserID = loginUserID
	r.GoalID = 0
	month, _ := c.GetPostForm("Month")
	r.Month, _ = time.Parse("2006-01", month)
	r.KeepInLastMonth, _ = c.GetPostForm("KeepInLastMonth")
	r.ProblemInLastMonth, _ = c.GetPostForm("ProblemInLastMonth")
	r.GoalAfterHalfYear, _ = c.GetPostForm("GoalAfterHalfYear")
	r.GoalInThisMonth, _ = c.GetPostForm("GoalInThisMonth")
	r.CurrentState, _ = c.GetPostForm("CurrentState")
	r.DailyTodo, _ = c.GetPostForm("DailyTodo")

	r.Add(&r)

	c.Redirect(http.StatusMovedPermanently, "/_monthly_plans")
}

//// GetGenres...
//func (h *Handler) GetGenre(c *gin.Context) {
//	r := models.NewGenreRepository()
//	id, _ := strconv.Atoi(c.Param("id"))
//	genre := r.GetOne(id)
//	c.HTML(http.StatusOK, "genre_edit.html", gin.H{
//		"genre": genre,
//	})
//}
//
//// EditGenres...
//func (h *Handler) EditGenre(c *gin.Context) {
//	r := models.NewGenreRepository()
//	id, _ := strconv.Atoi(c.Param("id"))
//	genre := r.GetOne(id)
//	genre.GenreName, _ = c.GetPostForm("genreName")
//	r.Edit(genre)
//	c.Redirect(http.StatusMovedPermanently, "/_genres")
//}
//
//// DeleteGenres...
//func (h *Handler) DeleteGenre(c *gin.Context) {
//	r := models.NewGenreRepository()
//	id, _ := strconv.Atoi(c.Param("id"))
//	r.Delete(id)
//	c.Redirect(http.StatusMovedPermanently, "/_genres")
//}
