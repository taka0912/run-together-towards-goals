package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
)

// GetAllMonthlyPlans ...
func (h *Handler) GetAllMonthlyPlans(c *gin.Context) {
	loginUserID := GetLoginUserID(c)

	r := models.NewMonthlyPlanRepository()
	monthlyPlans := r.GetAll(loginUserID)

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
	goalID, _ := c.GetPostForm("GoalID")
	r.GoalID, _ = strconv.Atoi(goalID)
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

// GetGenres...
func (h *Handler) GetMonthlyPlan(c *gin.Context) {
	r := models.NewMonthlyPlanRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	monthlyPlan := r.GetOne(id)

	rg := models.NewGoalRepository()
	goals := rg.GetByUserID(GetLoginUserID(c))

	c.HTML(http.StatusOK, "monthly_plan_edit.html", gin.H{
		"monthlyPlan": monthlyPlan,
		"goals":       goals,
	})
}

// EditGenres...
func (h *Handler) EditMonthlyPlan(c *gin.Context) {
	r := models.NewMonthlyPlanRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	monthlyPlan := r.GetOne(id)

	r.GoalID, _ = strconv.Atoi(c.Param("GoalID"))
	month, _ := c.GetPostForm("Month")
	monthlyPlan.Month, _ = time.Parse("2006-01", month)
	monthlyPlan.KeepInLastMonth, _ = c.GetPostForm("KeepInLastMonth")
	monthlyPlan.ProblemInLastMonth, _ = c.GetPostForm("ProblemInLastMonth")
	monthlyPlan.GoalAfterHalfYear, _ = c.GetPostForm("GoalAfterHalfYear")
	monthlyPlan.GoalInThisMonth, _ = c.GetPostForm("GoalInThisMonth")
	monthlyPlan.CurrentState, _ = c.GetPostForm("CurrentState")
	monthlyPlan.DailyTodo, _ = c.GetPostForm("DailyTodo")

	r.Edit(monthlyPlan)
	c.Redirect(http.StatusMovedPermanently, "/_monthly_plans")
}

// DeleteGenres...
func (h *Handler) DeleteMonthlyPlan(c *gin.Context) {
	r := models.NewMonthlyPlanRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	r.Delete(id)
	c.Redirect(http.StatusMovedPermanently, "/_monthly_plans")
}
